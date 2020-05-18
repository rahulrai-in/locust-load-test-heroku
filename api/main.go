package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	requestCount := 0

	portNumber := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Served home request")
		fmt.Fprintf(w, "Server OK")
	})

	http.HandleFunc("/volatile", func(w http.ResponseWriter, r *http.Request) {
		// For every 10 requests, delay the response by 1 second, up to 5 seconds.
		requestCount++
		delay := requestCount / 10
		if delay > 5 {
			requestCount = 0
			delay = 5
		}

		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Fprintf(w, "Produced response after %d second/s", delay)
		fmt.Printf("Produced response after %d second/s \n", delay)
	})

	http.HandleFunc("/buggy", func(w http.ResponseWriter, r *http.Request) {
		// After every 5 requests, throw error
		count := requestCount
		delay := count / 10
		if count%5 == 0 {
			fmt.Printf("Returning error at %d request \n", count)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		fmt.Fprintf(w, "Ok for request %d", count)
		fmt.Printf("Produced response after %d second/s \n", delay)
	})

	fmt.Println("Server listening on port ", portNumber)
	err := http.ListenAndServe(":"+portNumber, nil)
	if err != nil {
		fmt.Printf("Failure %s", err.Error())
	}
}
