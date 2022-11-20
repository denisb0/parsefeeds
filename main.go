package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
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

func readFeed(fp *gofeed.Parser, url string) (json.RawMessage, error) {
	// feed, err := fp.ParseURL("https://tech.groww.in/feed")
	// if err != nil {
	// 	log.Fatalf("Feed parse error %v", err)
	// }
	// fmt.Println(feed.Title)

	return nil, errors.New("not implemented")
}

func writeJson(fileName string, contents json.RawMessage) error {
	return errors.New("not implemented")
}

func main() {
	urls := readCsvFile("selecturls_2022-11-20T18_27_39.403218Z.csv")
	fmt.Println(urls)
	fp := gofeed.NewParser()

	for i, url := range urls {
		j, err := readFeed(fp, url)

		if err != nil {
			fmt.Println(fmt.Sprintf("%s feed error: %s", url, err.Error()))
			continue
		}

		err = writeJson(fmt.Sprintf("%d.json", i), j)
		if err != nil {
			fmt.Println(fmt.Sprintf("%d file write error: %s", i, err.Error()))
			continue
		}

		fmt.Println(fmt.Sprintf("%s feed processed", url))
	}

	fmt.Println("Done")
}
