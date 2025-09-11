package rss

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ssd-81/RSS-feed-/internal/database"
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
	req.Header.Set("User-Agent", "gator")
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("client could not make the request")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("the body could not be parsed")
	}
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

func ScrapeFeeds(ctx context.Context, s *types.State) error {
	fmt.Println("-- scrape feeds called --")
	lastUpdatedFeed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("error encountered 1", err)
		return err
	}
	// fmt.Println(lastUpdatedFeed)
	err = s.Db.MarkFeedFetched(context.Background(), lastUpdatedFeed.ID)
	if err != nil {
		return err
	}
	// double check the below line; highly likely some error would be introduced here
	feedData, err := FetchFeed(context.Background(), lastUpdatedFeed.Url.String)
	if err != nil {
		return err
	}
	// fmt.Println(feedData)
	for _, val := range feedData.Channel.Item {
		// converting the RFC22 -> RFC1123Z (time.Time)
		t, err := time.Parse(time.RFC1123Z, val.PubDate)
		if err != nil {
			return err
		}

		// make sure the following params are properly stored in the db
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
			Title:       sql.NullString{String: val.Title, Valid: true},
			Url:         sql.NullString{String: val.Link, Valid: true},
			Description: sql.NullString{String: val.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: t, Valid: true},
			FeedID:      uuid.NullUUID{UUID: lastUpdatedFeed.ID, Valid: true},


		}
		post, err := s.Db.CreatePost(context.Background(), params)
		if err != nil {
			fmt.Println(err)
			fmt.Println(post)
		} else {
			fmt.Println("post saved to db successfully")
		}
		fmt.Println(val.Title)
		fmt.Println()
	}

	return nil
}
