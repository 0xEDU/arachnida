package main

import (
	"etachott/spider/pkg/crawler"
	"etachott/spider/pkg/options"
	"fmt"
)

func main() {
	opts := options.NewOptions()

	// bfs
	images := crawler.Crawl(opts)

	// download images
	fmt.Println(images)
}
