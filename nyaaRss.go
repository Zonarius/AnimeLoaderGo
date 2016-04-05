package AnimeLoaderGo

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

// NyaaTorrent represents a single torrent on nyaa
type NyaaTorrent struct {
	Title       string `xml:"title"`
	Category    string `xml:"category"`
	Link        string `xml:"link"`
	Guid        string `xml:"guid"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type channel struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description string        `xml:"description"`
	Items       []NyaaTorrent `xml:"item"`
}

type rss struct {
	Channel channel `xml:"channel"`
}

var (
	nyaaURL = url.URL{Scheme: "http", Host: "www.nyaa.se", Path: "/"}
)

// Fetch fetches the most recent torrent files from nyaa
func Fetch(searchQuery string) (items []NyaaTorrent, err error) {
	ret := rss{}

	url := nyaaURL
	url.RawQuery = "page=rss&term=" + searchQuery
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := xml.NewDecoder(resp.Body)
	err = dec.Decode(&ret)
	return ret.Channel.Items, err
}
