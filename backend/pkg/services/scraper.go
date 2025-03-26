package services

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeGlassdoor scrapes salary data from a provided URL.
// This is a basic example and the actual selectors will depend on the target page structure.
func ScrapeGlassdoor(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(bodyBytes)))
	if err != nil {
		return nil, err
	}

	var salaries []string
	doc.Find(".salaryClass").Each(func(i int, s *goquery.Selection) {
		salary := strings.TrimSpace(s.Text())
		if salary != "" {
			salaries = append(salaries, salary)
		}
	})

	return salaries, nil
}
