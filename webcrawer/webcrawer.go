// Code for https://tour.golang.org/concurrency/10

package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mux sync.Mutex
	url map[string]bool
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	//Fetch URLs in parallel.
	//Don't fetch the same URL twice.
	if depth <= 0 {
		return
	}
	ch := make(chan []string)
	go fetchurls(url, ch)
	urls := <-ch
	for _, u := range urls {
		if Cache_hit(&cache, u) {
			continue
		}
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func Cache_load(c *Cache, url string) {

	c.mux.Lock()
	if c.url == nil {
		c.url = make(map[string]bool)
		c.url[url] = false
	}
	if !c.url[url] {
		c.url[url] = true
	}
	c.mux.Unlock()
	return
}

func Cache_hit(c *Cache, url string) bool {

	c.mux.Lock()
	if c.url == nil {
		c.mux.Unlock()
		return false
	}
	if c.url[url] {
		c.mux.Unlock()
		return true
	}
	c.mux.Unlock()
	return false
}

func fetchurls(url string, ch chan []string) {
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch <- nil
		return
	}
	Cache_load(&cache, url)
	fmt.Printf("found: %s %q\n", url, body)
	ch <- urls
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var cache Cache

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
