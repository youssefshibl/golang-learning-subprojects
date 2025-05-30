package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

var baseUrl string
var paths map[string]bool = make(map[string]bool)
var deepLevel int
var lock sync.Mutex
var wg sync.WaitGroup

func main() {

	mainUrl := flag.String("url", "", "enter url of page you want to crawl")
	deepLevelFlag := flag.Int("deep", 10, "enter deep level of crawling")

	flag.Parse()
	if *mainUrl == "" {
		CheckError(errors.New("you should pass url like '-url test.com'"))
	}
	deepLevel = *deepLevelFlag
	baseUrl = *mainUrl
	wg.Add(1)
	fetchPaths([]string{"/"}, 1)
	wg.Wait()

}

func fetchPaths(urls []string, level int) {
	defer wg.Done()
	if level > deepLevel {
		return
	}

	for _, url := range urls {
		lock.Lock()
		if _, exists := paths[url]; exists {
			lock.Unlock()
			continue

		}
		paths[url] = true
		lock.Unlock()
		fullPath := baseUrl + url
		fmt.Printf("Crawling %s at level %d\n", fullPath, level)
		pageContent := getUrlContentPage(fullPath)
		newUrls := ParseLinks(string(pageContent))
		newUrls = filterUrls(newUrls)
		// fmt.Println("Found URLs:", newUrls)
		wg.Add(1)
		go fetchPaths(newUrls, level+1)

	}
}

func getUrlContentPage(url string) string {
	resp, err := http.Get(url)
	CheckError(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	CheckError(err)
	return string(body)
}

func filterUrls(urls []string) []string {
	var newUrls []string
	for _, url := range urls {
		// add path only not other urls
		if !strings.Contains(url, "http") {
			newUrls = append(newUrls, url)
		}
	}

	return newUrls
}
