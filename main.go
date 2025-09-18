package main

import (
	"fmt"
	"os"

	"github.com/ewilan-riviere/rss-downloader/pkg/fetch"
	"github.com/ewilan-riviere/rss-downloader/pkg/xml"
	flag "github.com/spf13/pflag"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run main.go <rss_feed_url> [max_items]")
	// 	os.Exit(1)
	// }
	// url := os.Args[1]
	// maxItems := -1
	// if len(os.Args) >= 3 {
	// 	if n, err := strconv.Atoi(os.Args[2]); err == nil && n > 0 {
	// 		maxItems = n
	// 	}
	// }

	// flags
	var reverse = false
	var print = true
	var download = false
	var limit = -1
	var outputDir = "downloads"

	flag.BoolVarP(&reverse, "reverse", "r", false, "Reverse the list of episodes")
	flag.BoolVarP(&print, "print", "p", false, "Print episode list in console")
	flag.BoolVarP(&download, "download", "d", false, "Download episodes")
	flag.IntVarP(&limit, "max", "m", -1, "Max number of episodes to process (default all)")
	flag.StringVarP(&outputDir, "out", "o", "downloads", "Output directory for downloaded episodes")
	flag.Parse()
	// -----

	// handle no args case
	if flag.NArg() < 1 {
		fmt.Println("Usage: go run main.go [options] <rss_feed_url>")
		flag.PrintDefaults()
		os.Exit(1)
	}
	url := flag.Arg(0)

	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Options: reverse=%v, print=%v, download=%v, maxItems=%d\n", reverse, print, download, limit)
	fmt.Printf("-----\n")
	// -----

	// fetch RSS feed
	data, err := fetch.FetchUrl(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur fetch: %v\n", err)
		os.Exit(2)
	}
	// -----

	// parse RSS feed
	xml.ParseRSS(data, limit, reverse, print, download, outputDir)
	// -----
}
