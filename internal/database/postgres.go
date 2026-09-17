package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	var ctx  context.Context =  context.Background()

	var config *pgxpool.Config
	var err error
	
	config, err = pgxpool.ParseConfig(databaseURL)

	if err != nil {
		log.Println("Error parsing database URL:", err)
		return nil, err
	}

	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}

	err = pool.Ping(ctx)

	if err != nil {
		log.Println("Error pinging the database:", err)
		pool.Close()
		return nil, err
	}


	log.Println("Successfully connected to the database")
	return pool, nil
}