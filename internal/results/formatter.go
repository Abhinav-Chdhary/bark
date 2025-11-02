package results

// Formatter is an interface for formatting scan results
type Formatter interface {
	Format(result *ScanResult) (string, error)
}
