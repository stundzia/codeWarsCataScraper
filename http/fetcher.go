package http

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)


func FetchHtmlDoc(url string) (doc *goquery.Document, err error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Fetching `%s` got error: %s", url, err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Fetching `%s` got a non-200 status code: %d", url, resp.StatusCode)
		return nil, err
	}
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("doc from response got error: %s", err)
		return nil, err
	}
	return doc, nil
}