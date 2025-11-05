package parser

import (
	"github.com/debkanchan/bark/internal/parser/languages"
)

// Registry holds all supported languages
type Registry struct {
	languages       []languages.Language
	extensionLookup map[string]*languages.Language
}

// NewRegistry creates a new language registry with all supported languages
func NewRegistry() *Registry {
	languageList := []languages.Language{
		languages.Go(),
		languages.JavaScript(),
		languages.TypeScript(),
		languages.Python(),
		languages.Java(),
		languages.C(),
		languages.Cpp(),
		languages.JSON(),
		languages.Bash(),
		languages.Lua(),
		languages.HCL(),
		languages.YAML(),
		languages.XML(),
		languages.TOML(),
		languages.Rust(),
		languages.Zig(),
		languages.Kotlin(),
	}

	// Build extension lookup map
	extensionLookup := make(map[string]*languages.Language)
	for i := range languageList {
		for _, ext := range languageList[i].Extensions {
			extensionLookup[ext] = &languageList[i]
		}
	}

	return &Registry{
		languages:       languageList,
		extensionLookup: extensionLookup,
	}
}

// GetLanguageByExtension returns the language for a given file extension
func (r *Registry) GetLanguageByExtension(ext string) (*languages.Language, bool) {
	lang, found := r.extensionLookup[ext]
	return lang, found
}

// GetSupportedExtensions returns all supported file extensions
func (r *Registry) GetSupportedExtensions() []string {
	extensions := make([]string, 0, len(r.extensionLookup))
	for ext := range r.extensionLookup {
		extensions = append(extensions, ext)
	}
	return extensions
}

// GetLanguages returns all supported languages
func (r *Registry) GetLanguages() []languages.Language {
	return r.languages
}
