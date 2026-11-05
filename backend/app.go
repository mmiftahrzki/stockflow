package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type app struct {
	server *http.Server
	log    *logrus.Entry
}

func NewApp(cfg AppConfig, db *pgxpool.Pool) *app {
	app_logger := GetLogger().WithField("component", "app")

	return &app{
		log: app_logger,
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Port),
			Handler:      newMux(db, cfg),
			WriteTimeout: time.Second * 30,
			ReadTimeout:  time.Second * 10,
		},
	}
}