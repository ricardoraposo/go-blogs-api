package database

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ricardoraposo/blogs-api-go/config"
)

type Database struct {
	*sql.DB
}

func NewDatabase() *Database {
	c := config.LoadConfig("./")

	db, err := sql.Open("mysql", c.GetDSN())
	if err != nil {
		panic(err)
	}

	return &Database{DB: db}
}

func (d *Database) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := d.DB.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
