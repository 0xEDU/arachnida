package options

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	UseRecursion   bool
	RecursionDepth int
	DataPath       string
	Arguments      []string
}

func NewOptions() *Options {
	flagUseRecursion := flag.Bool("r", false, "Use recursion for downloading images")
	flagRecursionDepth := flag.Int("l", 5, "Maximum depth for recursion")
	flagDataPath := flag.String("p", "./data", "Path to store downloaded files")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: spider [-rlp] URL\n\nFlags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	return &Options{
		*flagUseRecursion,
		*flagRecursionDepth,
		*flagDataPath,
		flag.Args(),
	}
}
