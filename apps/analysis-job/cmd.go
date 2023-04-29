package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	// environment value
	targetEndpoint string
	canaryFlag     bool
	threshold      int
)

type Response struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func main() {
	log.Println("[START] analysis-job")

	// クライアント停止シグナル待ち受けチャネル定義
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		log.Println("[END] analysis-job")
		os.Exit(0)
	}()

	// set environment value
	targetEndpoint = os.Getenv("TARGET_ENDPOINT")
	canaryFlag, _ = strconv.ParseBool(os.Getenv("CANARY_FLAG"))
	threshold, _ = strconv.Atoi(os.Getenv("ERROR_THRESHOLD"))

	// execute analysis
	execute()

	log.Println("[END] analysis-job")
}

func execute() {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	// initialize value
	var resArray []int
	for i := 0; i < 60; i++ {
		resArray = append(resArray, 0)
	}

	idx := 0
	for {
		idx++

		// send http request
		retVal := loop_run(client, idx)
		resArray = append(resArray[1:], retVal)

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

func loop_run(httpClient *http.Client, idx int) int {
	req, _ := http.NewRequest("GET", targetEndpoint, nil)
	if canaryFlag {
		req.Header.Set("X-Canary", "true")
	}

	// execute
	log.Printf("Send request[%d]\n", idx)
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Fatal http request send: %v\n", err)
	}

	body, _ := io.ReadAll(res.Body)
	res.Body.Close()

	// http response check
	if code := res.StatusCode; code != http.StatusOK {
		log.Printf("Invalid status code: %d\n", code)
		return 1
	}

	var resJson Response
	json.Unmarshal(body, &resJson)
	log.Printf("Receive response[%d]: {version: %s}\n", idx, resJson.Version)
	return 0
}
