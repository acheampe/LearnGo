package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)} // send statement
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // receive expression
		results[r.string] = r.bool
	}

	return results
}
