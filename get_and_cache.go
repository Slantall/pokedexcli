package main

import (
	"fmt"
	"io"
	"net/http"
)

func getAndCache(url string) (data []byte, err error) {
	data, found := cache.Get(url)
	if !found {
		//fmt.Println("Cache miss - making HTTP request...")
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error making request: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return nil, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, data)
		}
		if err != nil {
			return nil, fmt.Errorf("error reading response: %w", err)
		}
		cache.Add(url, data)
	} else {
		//fmt.Println("Cache hit!")
	}
	return data, nil
}
