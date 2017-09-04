package models

import (
	"log"

	"github.com/go-pg/pg"
)

func InitDb() *pg.DB {

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "",
		Database: "postgres",
	})
	if db == nil {
		log.Fatal("DB connect fail")
	}
	return db
}
