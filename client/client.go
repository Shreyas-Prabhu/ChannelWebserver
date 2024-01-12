package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//arr := []time.Duration{}
	c := make(chan time.Duration)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bodyReader := bytes.NewReader([]byte("hello"))
			func(start time.Time) {
				_, err := http.Post("http://localhost:4000", "application/json", bodyReader)
				if err != nil {
					log.Fatal(err)
				}
				passedTime := time.Since(start)

				c <- passedTime
				//arr = append(arr, passedTime)

			}(time.Now())

		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	for pT := range c {
		fmt.Println("channel", pT)
	}
}
