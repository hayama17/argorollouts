package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	client := &http.Client{}

	eg := errgroup.Group{}
	parentCtx := context.Background()

	targetEndpoint := os.Getenv("TARGET_ENDPOINT")
	maxLoop, _ := strconv.Atoi(os.Getenv("REQUEST_MAX_LOOP"))

	errorCnt := 0
	for i := 0; i < maxLoop; i++ {
		log.Printf("Send http request[%d]\n", i)

		i := i
		eg.Go(func() error {
			childCtx, cancel := context.WithTimeout(parentCtx, time.Duration(3*time.Second))
			defer cancel()

			req, _ := http.NewRequestWithContext(childCtx, "GET", targetEndpoint, nil)
			res, err := client.Do(req)
			if err != nil {
				return err
			}

			log.Printf("Response receive[%d]\n", i)

			// http response check
			if code := res.StatusCode; code != http.StatusOK {
				log.Printf("Invalid status code: %d\n", code)
				errorCnt++
			}

			return nil
		})

		sleepDuration, _ := time.ParseDuration(os.Getenv("SLEEP_DURATION"))
		time.Sleep(sleepDuration)
	}

	if err := eg.Wait(); err != nil {
		log.Fatalf("Fatal do http requuest: %v\n", err)
	}

	threshold, _ := strconv.Atoi(os.Getenv("ERROR_THRESHOLD"))
	if errorCnt > threshold {
		log.Fatalf("Fatal error count exceeded threshold: %d > %d\n", errorCnt, threshold)
	}
}
