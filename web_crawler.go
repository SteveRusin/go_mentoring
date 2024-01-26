package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeUrlMap struct {
	mu sync.Mutex
	v  map[string]bool
}

func (m *SafeUrlMap) Has(url string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.v[url]

	return ok
}

func (m *SafeUrlMap) Add(url string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.v[url] = true
}

var urlMap = SafeUrlMap{v: make(map[string]bool)}

func Crawl(url string, depth int, fetcher Fetcher) {
	wg.Add(1)
	defer wg.Done()
	if depth <= 0 {
		return
	}

	if urlMap.Has(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
  urlMap.Add(url)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	go Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
	// doesnt print log without wait
	time.Sleep(time.Second)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
