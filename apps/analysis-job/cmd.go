package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// クライアント停止シグナル待ち受けチャネル定義
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		log.Println("Interrupted infinite loop")
		os.Exit(0)
	}()

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	targetEndpoint := os.Getenv("TARGET_ENDPOINT")
	threshold, _ := strconv.Atoi(os.Getenv("ERROR_THRESHOLD"))

	var resArray []int
	for i := 0; i < 60; i++ {
		resArray = append(resArray, 0)
	}

	idx := 0
	for {
		idx++

		// execute
		log.Printf("Request send[%d]\n", idx)
		res, _ := client.Get(targetEndpoint)
		log.Printf("Response receive[%d]\n", idx)

		// http response check
		if code := res.StatusCode; code != http.StatusOK {
			log.Printf("Invalid status code: %d\n", code)
			resArray = append(resArray[1:], 1)
		}
		resArray = append(resArray[1:], 0)

		// calcurate
		errorCnt := 0
		for _, val := range resArray {
			errorCnt += val
		}
		if errorCnt > threshold {
			log.Fatalf("Fatal error count exceeded threshold: %d > %d\n", errorCnt, threshold)
		}

		// sleep
		time.Sleep(1 * time.Second)
	}
}
