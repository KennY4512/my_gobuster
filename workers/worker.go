package workers

import (
	"LANG_KENNY_GO25/basics"
	"bufio"
	"net/http"
	"sync"
	"time"
)

// Function that pperforms the work
func worker(wordList, targetURL string, workers, expiration int, quiet, csv bool) {
	var wg sync.WaitGroup
	var c basics.CSV // CSV object to handle CSV creation if needed

	// Defines the timeout for HTTP client
	client := &http.Client{
		Timeout: time.Duration(expiration) * time.Second,
	}

	// Semaphore to limit the number of concurrent workers
	sem := make(chan struct{}, workers)

	// Initalize a scanner with the wordlist
	scanner := bufio.NewScanner(basics.FileOpener(wordList))

	// Loop through each line in the word list
	for scanner.Scan() {
		curentWord := scanner.Text()
		wg.Add(1)
		sem <- struct{}{} // Acquire a semaphore slot

		// Launch a goroutine to handle the HTTP request
		go func(word string) {
			defer wg.Done()
			defer func() { <-sem }() // Release the semaphore slot

			// Send a GET request to the target URL
			resp, err := client.Get(targetURL + word)
			if err != nil {
				basics.HTTPErr(word, err)
				return
			}
			defer resp.Body.Close()

			// Controls how the result is displayed and put the result in a slice (for CSV)
			if quiet && resp.StatusCode == 200 {
				basics.HTTPDisp(word, resp.StatusCode)
				c.Add(word, resp.StatusCode)
			} else if !quiet {
				basics.HTTPDisp(word, resp.StatusCode)
				c.Add(word, resp.StatusCode)
			}

		}(curentWord) // Pass the current word to the goroutine
	}

	// Check for any errors encountered while scanning the file
	if err := scanner.Err(); err != nil {
		basics.FileReadErr(err)
	}

	wg.Wait()

	// If the csv flag is true, write the results to a CSV file
	if csv {
		c.Write()
	}
}
