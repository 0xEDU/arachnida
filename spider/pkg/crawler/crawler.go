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

func newStackUrl(address string, depth int) stackUrl {
	return stackUrl{address: address, depth: depth}
}

func fetchHtml(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	client := &http.Client{}

	resp, err := client.Do(req)
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
	maxDepth := 1
	if opts.UseRecursion {
		maxDepth = opts.RecursionDepth
	}

	urlQueue := make([]stackUrl, 0)
	urlQueue = append(urlQueue, newStackUrl(opts.Arguments[0], 0))

	visitedSet := make(map[string]struct{})

	for len(urlQueue) > 0 {
		var currentUrl stackUrl
		currentUrl, urlQueue = urlQueue[0], urlQueue[1:]

		if currentUrl.depth > maxDepth {
			continue
		}

		if _, visited := visitedSet[currentUrl.address]; visited {
			continue
		}

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

		for _, link := range data.Links {
			urlQueue = append(urlQueue, newStackUrl(link, currentUrl.depth+1))
		}

		visitedSet[currentUrl.address] = struct{}{}
		fmt.Println(data)
		fmt.Printf("current depth = %d \n\n\n\n", currentUrl.depth)
	}
}
