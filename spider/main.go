package main

import (
	"etachott/spider/pkg/crawler"
	"etachott/spider/pkg/options"
)

func main() {
	opts := options.NewOptions()

	// bfs
	crawler.Crawl(opts)

	// download images
}
