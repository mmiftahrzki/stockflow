package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var once sync.Once
var db *pgxpool.Pool
var log *logrus.Entry = GetLogger().WithField("component", "database")

func new(cfg DatabaseConfig) (*pgxpool.Pool, error) {
	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
	)

	ctx := context.Background()
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new db pool %s: %w", cfg.Host, err)
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", cfg.Host, err)
	}

	return db, nil
}

func NewDB(cfg DatabaseConfig) (*pgxpool.Pool, error) {
	var err error

	if db != nil {
		return db, nil
	}

	once.Do(func() {
		db, err = new(cfg)
	})

	return db, err
}
