package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ewilan-riviere/rss-downloader/pkg/rss"
	"github.com/ewilan-riviere/rss-downloader/pkg/slugify"
	"github.com/ewilan-riviere/rss-downloader/pkg/utils"
)

func safeFileName(s string) string {
	s = strings.TrimSpace(s)
	s = slugify.Slugify(s, false)
	return s
}

func DownloadEpisode(i int, item rss.Item, outputDir string) error {
	fmt.Printf("\n")
	if item.Enclosure.URL == "" {
		return fmt.Errorf("no enclosure link found for item %d: %s", i+1, item.Title)
	}

	// Create the downloads/ directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	// Construct the name
	date := utils.ParsePubDate(item.PubDate)
	name := fmt.Sprintf("%02d_%s_%s_%s.mp3",
		i+1,
		safeFileName(item.Title),
		safeFileName(date),
		safeFileName(item.GUID),
	)
	path := filepath.Join(outputDir, name)

	// Skip if already present
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("Already downloaded: %s\n", path)
		return nil
	}

	fmt.Printf("Downloading %s\n", item.Enclosure.URL)
	resp, err := http.Get(item.Enclosure.URL)
	if err != nil {
		return fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	fmt.Printf("✅ Saved: %s\n", path)
	return nil
}
