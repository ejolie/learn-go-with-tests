package concurrency

import "fmt"

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result) // a channel of result

	for _, url := range urls {
		go func(u string) {
			fmt.Printf("send %s\n", u)
			resultChannel <- result{u, wc(u)} // send
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // receive
		results[r.string] = r.bool
		fmt.Printf("receive %v\n", r)
	}

	return results
}
