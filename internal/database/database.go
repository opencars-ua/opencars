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
	Update(model interface{}) error
	Insert(model interface{}) error
	Select(
		model interface{},
		limit int,
		condition string,
		params ...interface{},
	) error
}

// TODO: Make this func clear.
func CreateSchema(db *pg.DB) error {
	tables := []interface{}{
		(*model.Operation)(nil), (*model.Registration)(nil),
	}

	for _, table := range tables {
		err := db.CreateTable(table, &orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}

	err := db.CreateTable((*model.Registration)(nil), &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}

	queries := []string{
		"CREATE INDEX IF NOT EXISTS NUMBERS ON operations USING btree (number)",
		"CREATE INDEX IF NOT EXISTS ops_vin_codes ON operations USING btree (vin)",
		"CREATE INDEX IF NOT EXISTS reg_numbers ON registrations USING btree (number)",
		"CREATE INDEX IF NOT EXISTS reg_vin_codes ON registrations USING btree (vin)",
		"CREATE INDEX IF NOT EXISTS reg_codes ON registrations USING btree (code)",
	}

	for _, query := range queries {
		if _, err = db.Exec(query); err != nil {
			return err
		}
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

	// TODO: Move secrets to toml configuration file.
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
