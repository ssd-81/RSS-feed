package rss

import (
	"context"
	"encoding/xml"
	"errors"
	"html"
	"io"
	"net/http"
	"fmt"

	"github.com/ssd-81/RSS-feed-/internal/types"

)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// creating a new http client (we might need to change headers)
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, errors.New("http request failed to the provided url")
	}
	// keep eye on this part; might cause problems
	req.Header.Set("User-Agent", "gator")
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("client could not make the request")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("the body could not be parsed")
	}
	// I am not very sure with this, this might cause some problems
	var rssData RSSFeed
	err = xml.Unmarshal(body, &rssData)
	if err != nil {
		return nil, errors.New("the rss data could not be read")
	}

	return &rssData, nil
}

func DecodeEscapedChars(r *RSSFeed) {
	r.Channel.Title = html.EscapeString(r.Channel.Title)
	r.Channel.Description = html.EscapeString(r.Channel.Description)
	// just make sure the changes made actually reflect in the original  variable
	for i := range r.Channel.Item {
		r.Channel.Item[i].Title = html.EscapeString(r.Channel.Item[i].Title)
		r.Channel.Item[i].Description = html.EscapeString(r.Channel.Item[i].Description)
	}
}

func ScrapeFeeds(ctx context.Context, s *types.State) (error) {
	
	lastUpdatedFeed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil { 
		return err
	}
	err  = s.Db.MarkFeedFetched(context.Background(), lastUpdatedFeed.ID)
	if err != nil {
		return err
	}
	// double check the below line; highly likely some error would be introduced here
	feedData, err := FetchFeed(context.Background(), lastUpdatedFeed.Url.String)
	if err != nil {
		return err
	}
	for _, val := range feedData.Channel.Item {
		fmt.Println(val.Title)
	}
	return nil 
}
