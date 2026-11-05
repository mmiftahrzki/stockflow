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
	logger := GetLogger()

	cfg, err := LoadConfig()
	if err != nil {
		logger.Fatalln(err)
	}

	db, err := NewDB(cfg.Database)
	if err != nil {
		logger.Fatalf("Database Error: %v\n", err)
	}
	defer db.Close()

	app := NewApp(cfg.App, db)
	sigInt := make(chan os.Signal, 1)
	ctx := context.Background()

	go func() {
		signal.Notify(sigInt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := app.server.Shutdown(ctx); err != nil {
			app.log.Fatalf("graceful shutdown failed: %v", err)
		}

		defer close(sigInt)
	}()

	log.Infof("server runnning on %s", app.server.Addr)
	if err := app.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		app.log.Fatalln(err)
	}

	<-sigInt
	log.Info("server stopped gracefully")
}
