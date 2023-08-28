package parser

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) (interface{}, error) {
	p, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	links := traverseTree(p)

	return links, nil
}

func traverseTree(n *html.Node) []Link {
	ls := []Link{}

	if n.Type == html.ElementNode && n.Data == "a" {
		l := Link{}
		l.Text = getNodeText(n)

		for _, v := range n.Attr {
			if v.Key == "href" {
				l.Href = v.Val
			}
		}

		ls = append(ls, l)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ls = append(ls, traverseTree(c)...)
	}

	return ls
}

func getNodeText(n *html.Node) string {
	t := ""

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		var buf bytes.Buffer
		html.Render(&buf, c)

		if c.FirstChild != nil {
			t += " " + getNodeText(c)
		}

		if c.Type == html.TextNode {
			t += strings.TrimSpace(buf.String())
		}
	}

	return t
}
