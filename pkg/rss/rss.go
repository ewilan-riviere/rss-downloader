package rss

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title          string    `xml:"title"`
	Author         string    `xml:"author"`
	ItunesAuthor   string    `xml:"itunes:author"`
	PubDate        string    `xml:"pubDate"`
	GUID           string    `xml:"guid"`
	Description    string    `xml:"description"`
	Enclosure      Enclosure `xml:"enclosure"`
	ItunesDuration string    `xml:"itunes:duration"`
	Explicit       string    `xml:"itunes:explicit"`
}

type Enclosure struct {
	URL    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Length string `xml:"length,attr"`
}
