package graph

// Define format constants
const (
	Excel = 1
	Word  = 2
	PDF   = 3
	HTML  = 4
)

// SetupConversions initializes conversion paths and weights.
func SetupConversions(g *Graph) {
	// Add edges with weights to represent quality/cost of conversions
	g.AddEdge(Excel, Word, 1) // High-quality Excel to Word
	g.AddEdge(Word, PDF, 2)   // Word to PDF with medium quality
	g.AddEdge(Excel, PDF, 3)  // Direct Excel to PDF with lower quality
	g.AddEdge(PDF, Excel, 4)  // PDF to Excel with low quality
	g.AddEdge(HTML, PDF, 4)   // HTML to PDF with medium quality
	g.AddEdge(HTML, Word, 1)  // HTML to Word with high quality
	g.AddEdge(PDF, HTML, 3)   // PDF to HTML with lower quality
	g.AddEdge(Word, HTML, 2)  // Word to HTML with medium quality
}
