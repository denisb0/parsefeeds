package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

func readCsvFile(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	urls := make([]string, len(records))

	for i := range records {
		urls[i] = records[i][0]
	}

	return urls
}

func readFeed(fp *gofeed.Parser, url string, maxItems int, maxContentLen int) (*gofeed.Feed, error) {
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	if len(feed.Items) > maxItems {
		feed.Items = feed.Items[:maxItems]
	}

	for i := range feed.Items {
		if len(feed.Items[i].Content) > maxContentLen {
			feed.Items[i].Content = feed.Items[i].Content[:maxContentLen] + "..."
		}
	}

	return feed, nil
}

func main() {

	args := os.Args[1:]
	if len(args) == 0 || args[0] == "" {
		log.Fatal("invalid args")
	}

	urls := readCsvFile(args[0])
	// fmt.Println(urls)

	fp := gofeed.NewParser()
	feeds := make([]*gofeed.Feed, 0, len(urls))
	feedErrors := make([]string, 0)

	for _, url := range urls {
		f, err := readFeed(fp, url, 3, 64)

		if err != nil {
			feedErrors = append(feedErrors, url)
			fmt.Println(fmt.Sprintf("%s feed error: %s", url, err.Error()))
			continue
		}

		feeds = append(feeds, f)

		fmt.Println(fmt.Sprintf("%s feed processed", url))
	}

	j, err := json.MarshalIndent(feeds, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(j))

	err = os.WriteFile("out.json", j, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Errors: %d", len(feedErrors)))
	fmt.Println("Done")
}
