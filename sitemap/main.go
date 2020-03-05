package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwita/gophercises/link"
)

/*
	1. GET the webpage
 	2. Parse all the links
 	3. process the urls (i.e. add missing domains)
	4. filter out any links w/ a different domain
	5. find all the pages (BFS)
	6. print out XML
*/

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links deep to traverse")

	flag.Parse()

	//fmt.Printf("%s", *urlFlag)

	pages := bfs(*urlFlag, *maxDepth)
	// pages := get(*urlFlag)
	for _, page := range pages {
		// if (!strings.HasPrefix(l.Href, base){
		fmt.Println(page)
		// }
	}

}

// Instead of using struct{}{}:
//		- the first {} for type definition
//		- the second {} is for the instantiation
type empty struct{}

func bfs(urlStr string, maxDepth int) []string {

	// visited nodes
	visited := make(map[string]struct{}) // struct uses less memory?? (from using a boolean) https://dave.cheney.net/2014/03/25/the-empty-struct
	// using struct is a way to implement HashSet instead of using a boolen as the value (which nees space in memory to be allocated)

	// then we need the 'traditional' queue data structure for implementing the bfs algorithm
	var q map[string]empty
	var nq = map[string]empty{
		urlStr: empty{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]empty)
		for url, _ := range q {
			// we don't care about the value here
			if _, ok := visited[url]; ok {
				// if already visited
				continue
			}
			// mark this url as visited
			visited[url] = empty{}
			for _, link := range get(url) {
				nq[link] = empty{}
			}
		}
	}
	ret := make([]string, 0, len(visited))
	for url, _ := range visited {
		ret = append(ret, url)
	}
	return ret
}

// func bfs_default(){

// }

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var ret []string

	// Filter out
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
			//default:
		}
	}
	return ret
}

func get(urlStr string) []string {
	// 1. GET the html page
	resp, err := http.Get(urlStr)
	if err != nil {
		//handle error
		//panic(err)
		return []string{}
	}

	// Run the function when the function I am in ends
	// Benefits:
	// 				1. you put the close just immediately after the use of the http
	//				2. If you put it at the end it is easy to forget about it
	//				3. If your program returns before reaching the last line, it will not close the connection neither

	defer resp.Body.Close()

	/*
		Cases:
			/some-path								Path (usually starts with a slash)
			https://gophercises.com/some-path		Full link
			#fragment
			mailto:name@domain
	*/

	// obtain the current domain this way, because the initial "urlFlag" may redirect
	// to somewhere else
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))

}

func filter(links []string, keepFn func(string) bool) []string {

	// creating a new slice is simpler instead of modifying the existing one
	// on the other side, bi proceeding with the second, we can save some memory space
	// However, the memory overhead here is pretty small

	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
