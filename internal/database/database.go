package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/shal/opencars/pkg/models"
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
	db := pg.Connect(&pg.Options{
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
