package concurrency

type WebsiteChecker func(string) bool

// It returns a map of each URL checked to a boolean value: `true` for a good response; `false` for a bad response.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// Maintain access to the lexical scope in which they are defined - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.

	// The problem here is that the variable url is reused for each iteration of the for loop - it takes a new value from urls each time.
	// But each of our goroutines have a reference to the url variable - they don't have their own independent copy.
	// So they're all writing the value that url has at the end of the iteration - the last url.
	// Which is why the one result we have is the last url.
	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}

	return results
}
