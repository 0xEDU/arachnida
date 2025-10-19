package crawler

import (
	"etachott/spider/pkg/options"
	"fmt"
	"io"
	"net/http"
)

type stackUrl struct {
	address string
	depth   int
}

func fetchData(url string) (string, error) {
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
	// var maxDepth int
	// if opts.UseRecursion {
	// 	maxDepth = opts.RecursionDepth
	// } else {
	// 	maxDepth = 1
	// }

	urlQueue := make([]stackUrl, 0)
	urlQueue = append(urlQueue, stackUrl{address: opts.Arguments[0], depth: 1})

	for len(urlQueue) > 0 {
		var currentUrl stackUrl
		currentUrl, urlQueue = urlQueue[0], urlQueue[1:]

		fmt.Println(fetchData(currentUrl.address))
	}
}
