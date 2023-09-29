package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()

	server, err := InitializeServer(ctx)
	if err != nil {
		panic(err)
	}

	shutdown := make(chan struct{})
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			panic(err)
		}
		close(shutdown)
	}()

	if err = server.Start(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}

	<-shutdown
}
