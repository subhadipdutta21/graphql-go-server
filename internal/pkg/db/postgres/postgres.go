package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func InitDbPool() {
	host := "localhost"
	port := "5432"
	user := "dockerpg"
	password := "dockerpg"
	dbname := "graphql-go-server"

	databaseUrl := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Print(err)
		log.Print("Error in config.")
		//return &pgxpool.Pool{}
	}

	pool, err = pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Print(err)
		log.Print("Could not connect to Postgres.")
	} else {
		log.Println("Postgres connected!")
	}
}

func GetPool() *pgxpool.Pool {
	if pool == nil {
		InitDbPool()
	}

	connectedPoolSize := pool.AcquireAllIdle(context.Background())
	for connectedPoolSize == nil {
		log.Println("Pg Connection Lost")
		pool.Close()
		time.Sleep(2 * time.Second)
		log.Print("Reconnecting...")
		InitDbPool()
		connectedPoolSize = pool.AcquireAllIdle(context.Background())
	}
	return pool
}
