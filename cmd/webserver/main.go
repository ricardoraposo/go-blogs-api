package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ricardoraposo/blogs-api-go/config"
	"github.com/ricardoraposo/blogs-api-go/internal/server"
)

func main() {
	c := config.LoadConfig("./")
	server := server.NewServer(c.WSPort)

	go func() {
        fmt.Println("Server started on port", c.WSPort)
		if err := server.ListenAndServe(); err != nil {
            log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

	fmt.Println("\nServer gracefully stopped")
}
