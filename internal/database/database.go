package database

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/shal/opencars/pkg/models"
	"os"
	"strings"
)

func CreateSchema(db *pg.DB) error {
	err := db.CreateTable((*models.Transport)(nil), &orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		return err
	}

	_, err = db.Model((*models.Transport)(nil)).Exec(
		"CREATE INDEX IF NOT EXISTS NUMBERS ON transports USING btree (number)",
	)

	if err != nil {
		return err
	}

	return nil
}

func DB() (*pg.DB, error) {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	database := "opencars"

	if len(strings.TrimSpace(os.Getenv("DATABASE_HOST"))) != 0 {
		host = os.Getenv("DATABASE_HOST")
	}

	if len(strings.TrimSpace(os.Getenv("DATABASE_PORT"))) != 0 {
		port = os.Getenv("DATABASE_PORT")
	}

	fmt.Println(fmt.Sprintf("%s:%s", host, port))

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		User:     user,
		Password: password,
		Database: database,
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
