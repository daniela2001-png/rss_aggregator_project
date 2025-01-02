package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Link struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Image struct {
	Text   string `xml:",chardata"`
	URL    string `xml:"url"`
	Title  string `xml:"title"`
	Link   string `xml:"link"`
	Width  string `xml:"width"`
	Height string `xml:"height"`
}

type Guid struct {
	Text        string `xml:",chardata"`
	IsPermaLink string `xml:"isPermaLink,attr"`
}

type RSSItem struct {
	Text        string   `xml:",chardata"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Creator     string   `xml:"creator"`
	PubDate     string   `xml:"pubDate"`
	Category    []string `xml:"category"`
	Guid        Guid     `xml:"guid"`
	Description string   `xml:"description"`
}

type Channel struct {
	Text            string    `xml:",chardata"`
	Title           string    `xml:"title"`
	Link            Link      `xml:"link"`
	Description     string    `xml:"description"`
	LastBuildDate   string    `xml:"lastBuildDate"`
	Language        string    `xml:"language"`
	UpdatePeriod    string    `xml:"updatePeriod"`
	UpdateFrequency string    `xml:"updateFrequency"`
	Generator       string    `xml:"generator"`
	Image           Image     `xml:"image"`
	Item            []RSSItem `xml:"item"`
}

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Content string   `xml:"content,attr"`
	Wfw     string   `xml:"wfw,attr"`
	Dc      string   `xml:"dc,attr"`
	Atom    string   `xml:"atom,attr"`
	Sy      string   `xml:"sy,attr"`
	Slash   string   `xml:"slash,attr"`
	Georss  string   `xml:"georss,attr"`
	Geo     string   `xml:"geo,attr"`
	Channel Channel  `xml:"channel"`
}

func urlToFeed(URL string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}
	response, err := httpClient.Get(URL)
	if err != nil {
		fmt.Println(err)
		return RSSFeed{}, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return RSSFeed{}, err
	}
	RSSFeedPointer := RSSFeed{}
	err = xml.Unmarshal(data, &RSSFeedPointer)
	if err != nil {
		fmt.Println(err)
		return RSSFeed{}, err
	}
	return RSSFeedPointer, nil
}
