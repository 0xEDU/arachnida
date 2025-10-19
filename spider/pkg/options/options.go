package options

import (
	"flag"
)

type Options struct {
	UseRecursion   bool
	RecursionDepth int
	DataPath       string
}

func NewOptions() *Options {
	flagUseRecursion := flag.Bool("r", false, "Use recursion for downloading images")
	flagRecursionDepth := flag.Int("l", 5, "Maximum depth for recursion")
	flagDataPath := flag.String("p", "./data", "Path to store downloaded files")

	flag.Parse()

	return &Options {
		*flagUseRecursion,
		*flagRecursionDepth,
		*flagDataPath,
	}
}
