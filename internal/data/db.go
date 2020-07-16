package data

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/taciomcosta/chesstournament/internal/config"
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
		Addr:     config.String("DB_ADDR"),
		User:     config.String("DB_USER"),
		Password: config.String("DB_PASSWORD"),
		Database: config.String("DB_DATABASE"),
		OnConnect: func(*pg.Conn) error {
			log.Println("Connected to Postgres successfully!")
			return nil
		},
	})
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.ChessClub)(nil),
		(*model.Player)(nil),
	}
	for _, model := range models {
		db.CreateTable(model, &orm.CreateTableOptions{})
	}
	return nil
}
