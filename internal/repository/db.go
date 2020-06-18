package repository

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/taciomcosta/chesstournament/internal/model"
)

func init() {
	db = newDB()
}

var db *pg.DB

func newDB() *pg.DB {
	db := connect()
	if err := createSchema(db); err != nil {
		panic(err)
	}
	return db
}

func connect() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "keycloak",
		Password: "password",
		Database: "chesstournament",
		OnConnect: func(*pg.Conn) error {
			fmt.Println("Connected to Postgres successfully!")
			return nil
		},
	})
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.ChessClub)(nil),
	}
	for _, model := range models {
		db.CreateTable(model, &orm.CreateTableOptions{})
	}
	return nil
}
