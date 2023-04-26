package main

import (
	"context"
	"crypto/rand"
	"log"
	"math/big"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var version = "none"

type Response struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func main() {
	// サーバ停止シグナル待ち受けチャネル定義
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	gin.SetMode(gin.ReleaseMode)

	// ルーティング定義
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		if version == "error" {
			val, _ := rand.Int(rand.Reader, big.NewInt(10))
			if val.Int64() == 0 {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}
		c.JSON(http.StatusOK, Response{
			Message: "Hello World!!!",
			Version: version,
		})
	})

	// HTTPサーバ定義
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// HTTPサーバ起動(非同期)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// サーバ停止シグナル待ち受け開始
	<-ctx.Done()

	// サーバ停止シグナル待ち受け停止
	stop()
	log.Println("shutting down gracefully...")

	// HTTPサーバ停止(タイムアウト: 5秒)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		// HTTPサーバ強制停止
		log.Fatal("server forced to shutdown: ", err)
	}

	log.Println("server shutdown completed.")
}
