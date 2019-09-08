package storage

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"github.com/opencars/opencars/internal/config"
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
func Migrate(db *pg.DB) error {
	tables := []interface{}{
		(*model.Operation)(nil),
		(*model.Registration)(nil),
	}

	// Create tables.
	for _, table := range tables {
		err := db.CreateTable(table, &orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}

	queries := []string{
		"CREATE INDEX IF NOT EXISTS ops_numbers ON operations USING btree (number)",
		"CREATE INDEX IF NOT EXISTS ops_vins ON operations USING btree (vin)",
		"CREATE INDEX IF NOT EXISTS reg_numbers ON registrations USING btree (number)",
		"CREATE INDEX IF NOT EXISTS reg_vin_codes ON registrations USING btree (vin)",
		"CREATE INDEX IF NOT EXISTS reg_codes ON registrations USING btree (code)",
	}

	// Create indices.
	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	return nil
}

// New returns newly created database connection.
func New(config *config.TOML) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:       config.Database.Address(),
		User:       config.Database.User,
		Password:   config.Database.Password,
		Database:   config.Database.Name,
		PoolSize:   config.Database.Pool,
		MaxRetries: config.Database.MaxRetries,
		Network:    config.Database.Network,
	})

	return db, nil
}
