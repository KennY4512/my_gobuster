package workers

import (
	"LANG_KENNY_GO25/basics"
	"bufio"
	"net/http"
	"sync"
	"time"
)

func Worker(wordList, targetURL string, workers, expiration int, quiet bool) {
	var wg sync.WaitGroup

	client := &http.Client{
		Timeout: time.Duration(expiration) * time.Second,
	}

	sem := make(chan struct{}, workers)
	scanner := bufio.NewScanner(basics.FileOpener(wordList))

	for scanner.Scan() {
		value := scanner.Text()
		wg.Add(1)
		sem <- struct{}{}

		go func(val string) {
			defer wg.Done()
			defer func() { <-sem }()

			resp, err := client.Get(targetURL + val)

			if err != nil {
				basics.HTTPErr(val, err)
				return
			}

			defer resp.Body.Close()

			if quiet {
				if resp.StatusCode == 200 {
					basics.HTTPDisp(val, resp.StatusCode)
				}
			} else {
				basics.HTTPDisp(val, resp.StatusCode)
			}
		}(value)
	}

	if err := scanner.Err(); err != nil {
		basics.ReadingErr(err)
	}

	wg.Wait()
}
