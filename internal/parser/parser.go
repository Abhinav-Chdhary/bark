package parser

import (
	"fmt"
	"os"
	"strings"

	sitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/debkanchan/bark/internal/results"
)

// Parser handles parsing files for BARK comments
type Parser struct {
	registry *Registry
}

// NewParser creates a new parser with the language registry
func NewParser() *Parser {
	return &Parser{
		registry: NewRegistry(),
	}
}

// ParseFile parses a file and returns findings
func (p *Parser) ParseFile(filePath string) ([]results.Finding, error) {
	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Determine language by filename (checks extension first, then patterns)
	lang, found := p.registry.GetLanguageByFilename(filePath)
	if !found {
		// Not a supported language, skip silently
		return []results.Finding{}, nil
	}

	// Parse the file with tree-sitter
	parser := sitter.NewParser()
	defer parser.Close()

	if lang.Parser == nil {
		return nil, fmt.Errorf("parser is nil for %s", lang.Name)
	}

	parser.SetLanguage(lang.Parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, fmt.Errorf("tree is nil for %s", lang.Name)
	}
	defer tree.Close()

	// Create query for comments
	query, queryErr := sitter.NewQuery(lang.Parser, lang.Query)
	if queryErr != nil && queryErr.Message != "" {
		return nil, fmt.Errorf(
			"failed to create query for %s: %s at offset %d",
			lang.Name,
			queryErr.Message,
			queryErr.Offset,
		)
	}
	if query == nil {
		return nil, fmt.Errorf("query object is nil for %s", lang.Name)
	}
	defer query.Close()

	// Execute query
	cursor := sitter.NewQueryCursor()
	defer cursor.Close()

	// Collect BARK comments
	findings := []results.Finding{}

	// Execute the query and iterate through matches
	matches := cursor.Matches(query, tree.RootNode(), content)
	for match := matches.Next(); match != nil; match = matches.Next() {
		for _, capture := range match.Captures {
			node := capture.Node
			commentText := node.Utf8Text(content)

			// Match the marker "BARK" regardless of what follows (e.g., BARK:, BARK!, BARK - )
			if strings.Contains(commentText, "BARK") {
				startPoint := node.StartPosition()
				finding := results.Finding{
					FilePath: filePath,
					Line:     startPoint.Row + 1, // tree-sitter uses 0-based indexing
					Column:   startPoint.Column + 1,
					Comment:  strings.TrimSpace(commentText),
				}
				findings = append(findings, finding)
			}
		}
	}

	return findings, nil
}

// GetRegistry returns the language registry
func (p *Parser) GetRegistry() *Registry {
	return p.registry
}
