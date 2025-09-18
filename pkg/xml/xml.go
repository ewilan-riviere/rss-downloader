package xml

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/ewilan-riviere/rss-downloader/pkg/download"
	"github.com/ewilan-riviere/rss-downloader/pkg/rss"
	"github.com/ewilan-riviere/rss-downloader/pkg/utils"
)

func ParseRSS(data []byte, limit int, reverse bool, print bool, dl bool, outputDir string) []rss.Item {
	var rss rss.RSS
	if err := xml.Unmarshal(data, &rss); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur parse XML: %v\n", err)
		os.Exit(3)
	}

	fmt.Printf("Podcast: %s\n", rss.Channel.Title)

	items := rss.Channel.Items
	total := len(items)
	if limit > 0 && limit < total {
		total = limit
		items = items[:total]
	}

	if reverse {
		// inverser la liste
		for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
			items[i], items[j] = items[j], items[i]
		}
	}

	if print {
		for i := 0; i < total; i++ {
			printItem(items[i], i)
		}
	}

	fmt.Printf("Total items: %d\n", total)

	if dl {
		for i := 0; i < total; i++ {
			if err := download.DownloadEpisode(i, items[i], outputDir); err != nil {
				fmt.Fprintf(os.Stderr, "Erreur téléchargement: %v\n", err)
			}
		}
	} else {
		fmt.Printf("-----\n")
		fmt.Println("Download skipped (use --download to enable)")
	}

	return items
}

func printItem(it rss.Item, i int) {
	fmt.Printf("\n#%02d: %s\n", i+1, it.Title)
	if it.ItunesAuthor != "" {
		fmt.Printf("   Author: %s\n", it.ItunesAuthor)
	} else if it.Author != "" {
		fmt.Printf("   Author: %s\n", it.Author)
	}
	if it.PubDate != "" {
		fmt.Printf("   Date: %s\n", utils.ParsePubDate(it.PubDate))
	}
	if it.GUID != "" {
		fmt.Printf("   GUID: %s\n", it.GUID)
	}
	if it.Description != "" {
		fmt.Printf("   Description: %s\n", it.Description)
	}
	if it.Enclosure.URL != "" {
		fmt.Printf("   Enclosure: %s\n", it.Enclosure.URL)
	} else {
		fmt.Printf("   Enclosure: (no enclosure found)\n")
	}
	if it.ItunesDuration != "" {
		fmt.Printf("   Duration: %s\n", it.ItunesDuration)
	}
	if it.Explicit == "yes" {
		fmt.Printf("   Explicit: Yes\n")
	} else if it.Explicit == "no" {
		fmt.Printf("   Explicit: No\n")
	} else if it.Explicit != "" {
		fmt.Printf("   Explicite: %s\n", it.Explicit)
	}
}
