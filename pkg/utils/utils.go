package utils

import "time"

func ParsePubDate(s string) string {
	// Try a few common RSS date formats to pretty print
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
		"Mon, 02 Jan 2006 15:04:05 -0700", // explicit
	}
	for _, f := range formats {
		if t, err := time.Parse(f, s); err == nil {
			return t.Format("2006-01-02 15:04")
		}
	}
	// fallback to raw string
	return s
}
