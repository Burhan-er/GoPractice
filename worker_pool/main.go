package main

import (
	"log"
	"net/http"
)

type Site struct {
	URL string
}
type Result struct {
	Status int
	URL    string
}

func crawl(wID int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("workers ID: %d\n", wID)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Printf("Error Get %s URl Addr, %s", site.URL, err.Error())
		}
		results <- Result{Status: resp.StatusCode, URL: site.URL}
	}

}

func main() {
	log.Print("worker pool started...\n")

	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for i := 1; i <= 3; i++ {
		go crawl(i, jobs, results)
	}

	urls := []string{
		"https://www.google.com",
		"https://www.w3schools.com",
		"https://www.trendyol.com",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}

	for result := range results {
		log.Printf("connection URL: %s, StatusCode: %d", result.URL, result.Status)
	}

}
