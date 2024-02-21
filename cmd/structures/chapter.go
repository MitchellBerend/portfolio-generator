package structures

type Chapter struct {
	Title   string      `json:"title"`
	Content []Paragraph `json:"content"`
}
