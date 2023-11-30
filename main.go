package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/cronJohn/godoro/cmd"
	db "github.com/cronJohn/godoro/db/sqlc"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger() // add file name and line number
}

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	handle := db.New(conn)

	result, err := handle.GetAll(ctx)
	if err != nil {
		panic(err)
	}

	log.Print(result)

	// cmd.Execute()
}
