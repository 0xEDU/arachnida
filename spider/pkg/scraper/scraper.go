package scraper

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type HtmlData struct {
	Links  []string
	Images []string
}

func formatLink(link string, baseUrl string) string {
	if strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://") {
		return link
	} else if strings.HasPrefix(link, "/") {
		return baseUrl + link
	}
	return ""
}

func formatImageSource(src string, baseUrl string) string {
	if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
		return src
	} else if strings.HasPrefix(src, "//") {
		if strings.HasPrefix(baseUrl, "https://") {
			return "https:" + src
		} else {
			return "http:" + src
		}
	} else if strings.HasPrefix(src, "/") {
		return baseUrl + src
	}
	return ""
}

func ExtractData(htmlBytes string, baseUrl string) (HtmlData, error) {
	document, err := html.Parse(strings.NewReader(htmlBytes))
	if err != nil {
		return HtmlData{}, err
	}

	links := []string{}
	images := []string{}

	for n := range document.Descendants() {
		if n.Type != html.ElementNode {
			continue
		}

		switch n.DataAtom {
		case atom.A:
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}
				links = append(links, formatLink(attr.Val, baseUrl))
			}
		case atom.Img:
			for _, attr := range n.Attr {
				if attr.Key != "src" ||
					(!strings.HasSuffix(attr.Val, ".jpg") &&
						!strings.HasSuffix(attr.Val, ".jpeg") &&
						!strings.HasSuffix(attr.Val, ".png") &&
						!strings.HasSuffix(attr.Val, ".gif") &&
						!strings.HasSuffix(attr.Val, ".bmp")) {
					continue
				}
				images = append(images, formatImageSource(attr.Val, baseUrl))
			}
		}
	}

	return HtmlData{links, images}, nil
}
