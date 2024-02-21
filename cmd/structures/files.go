package structures

type File struct {
	FileName  string             `json:"fileName"`
	Title     string             `json:"title"`
	Scripts   []string           `json:"scripts"`
	Header    Paragraph          `json:"header"`
	Body      []Chapter          `json:"body"`
	Footer    Paragraph          `json:"footer"`
	SidePanel []SidePanelElement `json:"sidePanel"`
}

func CreateExampleFile() File {
	return File{
		FileName: "example.txt",
		Title:    "Example Title",
		Scripts:  []string{"script1.js", "script2.js"},
		Header:   Paragraph{Title: "Header", Content: "This is the header."},
		Body: []Chapter{
			{
				Title: "Chapter 1",
				Content: []Paragraph{
					{
						Title:   "Intro",
						Content: "Introduction paragraph.",
					},
				},
			},
			{
				Title: "Chapter 2",
				Content: []Paragraph{
					{
						Title:   "Section 1",
						Content: "Section 1 content.",
					},
					{
						Title:   "Section 2",
						Content: "Section 2 content.",
					},
				},
			},
		},
		Footer: Paragraph{Title: "Footer", Content: "This is the footer."},
		SidePanel: []SidePanelElement{
			{
				Link:  "www.google.com",
				Image: "www.google.com",
			},
		},
	}
}
