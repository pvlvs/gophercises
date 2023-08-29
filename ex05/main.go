package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
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

var url *string

func main() {
	url = flag.String("url", "", "The url to create the sitemap for")
	flag.Parse()

	if *url == "" {
		panic("Please pass a valid URL")
	}

	var ps paths
	ps = append(ps, "/")
	ps.getPaths(*url)

	sm := ps.makeSitemap()

	fmt.Println(sm)
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

func (ps *paths) getPaths(link string) {
	links := getLinks(link)
	for _, v := range links {
		externalLink := !(strings.HasPrefix(v.Href, link) || strings.HasPrefix(v.Href, "/"))

		if !externalLink && !slices.Contains(*ps, v.Href) {
			*ps = append(*ps , v.Href)

			link := fmt.Sprintf("%v%v", *url, v.Href)
			ps.getPaths(link)
		}
	}
}

func (ps *paths) makeSitemap() string {
	var sm sitemap
	sm.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	sm.Urls = make([]sitemapUrl, len(*ps))

	for i, v := range *ps {
		var smUrl sitemapUrl

		link := fmt.Sprintf("%v%v", *url, v)
		smUrl.Loc = link
		sm.Urls[i] = smUrl
	}

	xmlData, err := xml.Marshal(sm)
	check(err)

	bs := append([]byte(xml.Header), xmlData...)

	return string(bs)
}
