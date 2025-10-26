package downloader

import (
	"etachott/spider/pkg/options"
	"io"
	"net/http"
	"os"
	"strings"
)

func getFileNameFromUrl(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func Download(imageUrls []string, opts *options.Options) {
	if _, err := os.Stat(opts.DataPath); os.IsNotExist(err) {
		os.MkdirAll(opts.DataPath, os.ModePerm)
	}

	for _, url := range imageUrls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			continue
		}

		file, err := os.Create(opts.DataPath + "/" + getFileNameFromUrl(url))
		if err != nil {
			continue
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			continue
		}
	}
}
