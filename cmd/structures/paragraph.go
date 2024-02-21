package structures

var (
	Normal ParagraphType = "normal"
	Link   ParagraphType = "link"
)

type ParagraphType string

type Paragraph struct {
	Type    ParagraphType `json:"type"`
	Title   string        `json:"title"`
	Link    string        `json:"link,omitempty"`
	Content string        `json:"content"`
}
