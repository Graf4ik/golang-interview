package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Query struct {
	StatusCode string
	Idx        int
}

func statusCodes(urls []string) []int {
	res := make([]int, len(urls), len(urls))
	resChan := make(chan Query)
	wg := sync.WaitGroup{}

	for i := 0; i < len(urls); i++ {
		idx := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(urls[i])
			if err != nil {
				resChan <- Query{
					StatusCode: -1,
					Idx:        idx,
				}
				return
			}

			resChan <- Query{
				StatusCode: resp.StatusCode,
				Idx:        idx,
			}
		}()
	}

	for ch := range resChan {
		res[ch.Idx] = ch.StatusCode
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	return res
}

func main() {
	urls := []string{
		"https://google.com",
		"https://bad.domain.local",
		"https://httpbin.org/status/404",
	}

	statuses := statusCodes(urls)
	fmt.Println(statuses)
}
