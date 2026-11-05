package main

import (
	"backend/internal"
	"backend/internal/logger"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log := logger.GetLogger()

	cfg, err := internal.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	db, err := internal.NewDB(ctx, cfg.Database)
	if err != nil {
		log.Fatalf("Database Error: %v\n", err)
	}
	defer db.Close()

	app := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      newMux(db),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
	}

	stopCh := make(chan struct{})
	go func() {
		sigInt := make(chan os.Signal, 1)
		signal.Notify(sigInt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		signal := <-sigInt

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		log.Info(fmt.Sprintf("received %s signal, shutting down gracefully", signal.String()))
		if err := app.Shutdown(ctx); err != nil {
			log.Fatalf("graceful shutdown failed: %v", err)
		}

		close(stopCh)
	}()

	log.Infof("server runnning on %s", app.Addr)
	if err := app.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}

	<-stopCh
	log.Info("server stopped gracefully")
}
