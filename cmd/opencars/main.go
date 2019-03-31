package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opencars-ua/opencars/internal/http"
)

type RealDB struct {
	*sqlx.DB
}

func (r *RealDB) Select(
	model interface{},
	limit int,
	condition string,
	params ...interface{},
) error {
	query := "SELECT * FROM transport WHERE " + condition + " LIMIT " + strconv.Itoa(limit)

	return r.DB.Select(model, query)
}

func main() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	database := "opencars"

	if os.Getenv("DATABASE_HOST") != "" {
		host = os.Getenv("DATABASE_HOST")
	}

	if os.Getenv("DATABASE_PORT") != "" {
		port = os.Getenv("DATABASE_PORT")
	}

	info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database,
	)

	DB := sqlx.MustConnect("postgres", info)
	http.DB = &RealDB{DB}

	http.Run()

	defer DB.Close()
}
