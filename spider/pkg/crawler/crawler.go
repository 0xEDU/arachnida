package crawler

import (
	"etachott/spider/pkg/options"
	"etachott/spider/pkg/scraper"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type stackUrl struct {
	address string
	depth   int
}

func fetchHtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Crawl(opts *options.Options) {
	// maxDepth := 1
	// if opts.UseRecursion {
	// 	maxDepth = opts.RecursionDepth
	// }

	urlQueue := make([]stackUrl, 0)
	urlQueue = append(urlQueue, stackUrl{address: opts.Arguments[0], depth: 1})

	for len(urlQueue) > 0 {
		var currentUrl stackUrl
		currentUrl, urlQueue = urlQueue[0], urlQueue[1:]

		htmlBytes, err := fetchHtml(currentUrl.address)
		if err != nil {
			continue
		}

		parsedUrl, err := url.Parse(currentUrl.address)
		if err != nil {
			continue
		}

		basePath := parsedUrl.Scheme + "://" + parsedUrl.Host
		data, err := scraper.ExtractData(htmlBytes, basePath)
		if err != nil {
			continue
		}

		fmt.Println(data)
	}
}
