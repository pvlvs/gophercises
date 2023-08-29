package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/pserghei/gophercises/ex04/parser"
	"golang.org/x/exp/slices"
)

type sitemap struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	Urls    []sitemapUrl `xml:"url"`
}

type sitemapUrl struct {
	Loc string `xml:"loc"`
}

type paths []string

type flags struct {
	url   *string
	depth *int
}

var fs flags

func main() {
	fs.url = flag.String("url", "", "The url to create the sitemap for")
	fs.depth = flag.Int("depth", int(math.Inf(1)), "Defines the maximum number of links to follow when building the sitemap")
	flag.Parse()

	if *fs.url == "" {
		log.Fatal("Please pass a valid URL")
        os.Exit(1)
	}
    if *fs.depth < 0 {
		log.Fatal("Please pass a positive number")
        os.Exit(1)
    }

	var ps paths
	ps = append(ps, "/")
	ps.getPaths(*fs.url, 0)

	fmt.Println(ps.makeSitemap())
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLinks(url string) []parser.Link {
	res, err := http.Get(url)
	check(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	check(err)

	html := string(body)
	r := strings.NewReader(html)

	links, err := parser.Parse(r)
	check(err)

	return links
}

func (ps *paths) getPaths(link string, depth int) {
	links := getLinks(link)
	for _, v := range links {
		externalLink := !(strings.HasPrefix(v.Href, link) || strings.HasPrefix(v.Href, "/"))

		if !externalLink && !slices.Contains(*ps, v.Href) {
			*ps = append(*ps, v.Href)

			link := fmt.Sprintf("%v%v", *fs.url, v.Href)

			if depth+1 <= *fs.depth {
				ps.getPaths(link, depth+1)
			}
		}
	}
}

func (ps *paths) makeSitemap() string {
	var sm sitemap
	sm.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	sm.Urls = make([]sitemapUrl, len(*ps))

	for i, v := range *ps {
		var smUrl sitemapUrl

		link := fmt.Sprintf("%v%v", *fs.url, v)
		smUrl.Loc = link
		sm.Urls[i] = smUrl
	}

	xmlData, err := xml.Marshal(sm)
	check(err)

	bs := append([]byte(xml.Header), xmlData...)

	return string(bs)
}
