package converters

// Provider represents a file conversion provider.
type Provider struct {
	SourceFormat string
	TargetFormat string
	Quality      int
}

// NewProviders returns a list of available providers for conversions.
func NewProviders() []Provider {
	return []Provider{
		{SourceFormat: "Excel", TargetFormat: "Word", Quality: 1},
		{SourceFormat: "Word", TargetFormat: "PDF", Quality: 2},
		{SourceFormat: "Excel", TargetFormat: "PDF", Quality: 3},
		{SourceFormat: "PDF", TargetFormat: "Excel", Quality: 4},
		{SourceFormat: "HTML", TargetFormat: "PDF", Quality: 2},
		{SourceFormat: "HTML", TargetFormat: "Word", Quality: 1},
		{SourceFormat: "PDF", TargetFormat: "HTML", Quality: 3},
		{SourceFormat: "Word", TargetFormat: "HTML", Quality: 2},
	}
}
