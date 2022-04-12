package reader

import (
	"errors"
	"github.com/mmcdole/gofeed"
	"time"
)

//RssItem will hold parser results
type RssItem struct {
	Title       string
	Source      string
	SourceURL   string
	Link        string
	PublishDate time.Time
	Description string
}

// Parse will retrieve concurrently RSS feed data from given array of URLs
// If Parse encounter critic error it will return it
// If Parse fail to retrieve feed data it will only log the error
func Parse(urls []string) ([]RssItem, error) {

	if len(urls) == 0 {

		err := errors.New("no urls provided")

		return []RssItem{}, err
	}

	urls = getUniqueValuesSlice(urls)

	c := make(chan *gofeed.Feed, len(urls))
	errC := make(chan error, len(urls))

	for _, v := range urls {
		go parseUrl(v, c, errC)
	}

	var (
		result []RssItem
		err    error
	)

	for range urls {

		select {

		case rss := <-c:
			for _, feed := range rss.Items {
				result = append(result, RssItem{
					Title:       feed.Title,
					Source:      rss.Title,
					SourceURL:   rss.Link,
					Link:        feed.Link,
					PublishDate: *feed.PublishedParsed,
					Description: feed.Description,
				})
			}
		case err = <-errC:

		}

	}

	return result, err
}

//getUniqueValuesSlice returns string slice with only unique values inside
func getUniqueValuesSlice(strSlice []string) []string {

	keys := make(map[string]bool, len(strSlice))

	var list []string

	for _, entry := range strSlice {

		if _, found := keys[entry]; !found {

			keys[entry] = true
			list = append(list, entry)

		}

	}

	return list
}

//parseUrl send to given channel gofeed.Feed after parsing given URL
func parseUrl(url string, c chan *gofeed.Feed, errC chan error) {

	feed, err := gofeed.NewParser().ParseURL(url)

	if err != nil {

		errC <- err
		return
	}

	c <- feed
}
