package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	var requestCount int32 = 0

	portNumber := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Served home request")
		fmt.Fprintf(w, "Server OK")
	})

	http.HandleFunc("/volatile", func(w http.ResponseWriter, r *http.Request) {
		// For every 10 requests, delay the response by 1 second, up to 5 seconds.
		currentCount := atomic.LoadInt32(&requestCount)
		atomic.AddInt32(&requestCount, 1)
		delay := currentCount / 10
		if delay > 5 {
			atomic.StoreInt32(&requestCount, 0)
			delay = 5
		}

		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Fprintf(w, "Produced response after %d second/s", delay)
		fmt.Printf("Produced response after %d second/s \n", delay)
	})

	http.HandleFunc("/buggy", func(w http.ResponseWriter, r *http.Request) {
		// After every 5 requests, throw error
		currentCount := atomic.LoadInt32(&requestCount)
		atomic.AddInt32(&requestCount, 1)
		if currentCount%5 == 0 {
			fmt.Printf("Returning error at %d request \n", currentCount)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		fmt.Fprintf(w, "Ok for request %d", currentCount)
	})

	fmt.Println("Server listening on port ", portNumber)
	err := http.ListenAndServe(":"+portNumber, nil)
	if err != nil {
		fmt.Printf("Failure %s", err.Error())
	}
}
