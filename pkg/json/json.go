package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ewilan-riviere/rss-downloader/pkg/rss"
)

func Save(items []rss.Item, jsonPath string) error {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}
	if err := os.WriteFile(jsonPath, data, 0644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	fmt.Printf("✅ JSON saved: %s\n", jsonPath)
	return nil
}
