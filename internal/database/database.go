package database

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"github.com/opencars/opencars/pkg/model"
)

// Database interface makes handler testable.
type Adapter interface {
	Healthy() bool
	Select(
		model interface{},
		limit int,
		condition string,
		params ...interface{},
	) error
}

func CreateSchema(db *pg.DB) error {
	err := db.CreateTable((*model.Operation)(nil), &orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		return err
	}

	_, err = db.Model((*model.Operation)(nil)).Exec(
		"CREATE INDEX IF NOT EXISTS NUMBERS ON operations USING btree (number)",
	)

	if err != nil {
		return err
	}

	return nil
}

func DB() (*pg.DB, error) {
	host := "localhost"
	port := "5432"

	if os.Getenv("DATABASE_HOST") != "" {
		host = os.Getenv("DATABASE_HOST")
	}

	if os.Getenv("DATABASE_PORT") != "" {
		port = os.Getenv("DATABASE_PORT")
	}

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		User:     "postgres",
		Password: "postgres",
		Database: "opencars",
	})

	err := CreateSchema(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Must(db *pg.DB, err error) *pg.DB {
	if err != nil {
		panic(err)
	}

	return db
}
