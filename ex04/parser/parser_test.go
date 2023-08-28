package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	fs, _ := os.ReadDir("../html/")
	htmls := []string{}

	for _, v := range fs {
		p := filepath.Join("..", "html", v.Name())
		f, _ := os.ReadFile(p)
		htmls = append(htmls, string(f))
	}

	tests := []struct {
		name string
		html string
		want []Link
	}{
		{
			name: "Single anchor tag",
			html: htmls[0],
			want: []Link{
				{
					Href: "/other-page",
					Text: "A link to another page",
				},
			},
		},
		{
			name: "Multiple anchor tags with additional tags",
			html: htmls[1],
			want: []Link{
				{
					Href: "https://www.twitter.com/joncalhoun",
					Text: "Check me out on twitter",
				},
				{
					Href: "https://github.com/gophercises",
					Text: "Gophercises is on Github!",
				},
			},
		},
		{
			name: "Multiple anchor tags with additional tags and comments and classes",
			html: htmls[2],
			want: []Link{
				{
					Href: "#",
					Text: "Login",
				},
				{
					Href: "/lost",
					Text: "Lost? Need help?",
				},
				{
					Href: "https://twitter.com/marcusolsson",
					Text: "@marcusolsson",
				},
			},
		},
		{
			name: "Comment inside anchor tag",
			html: htmls[3],
			want: []Link{
				{
					Href: "/dog-cat",
					Text: "dog cat",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.html)
			l, err := Parse(r)
			if err != nil {
				t.Errorf("Couldn't parse html")
			}

			for i := range tt.want {
				if l[i] != tt.want[i] {
					t.Errorf("got %v; want %v", l[i], tt.want[i])
				}
			}
		})
	}
}
