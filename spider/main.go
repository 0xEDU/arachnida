package main

import (
	"etachott/spider/pkg/crawler"
	"etachott/spider/pkg/downloader"
	"etachott/spider/pkg/options"
)

func main() {
	opts := options.NewOptions()
	images := crawler.Crawl(opts)
	downloader.Download(images, opts)
}
