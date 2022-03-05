package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/wasinaphatlilawatthananan/go-postgres/api"
	db "github.com/wasinaphatlilawatthananan/go-postgres/db/sqlc"
	"github.com/wasinaphatlilawatthananan/go-postgres/util"
)

func main() {
	config,err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAdress)

	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
