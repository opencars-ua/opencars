package adapter

import (
	"github.com/go-pg/pg"
)

// Adapter implements Storage interface from storage package.
type Adapter struct {
	db *pg.DB
}

// New returns new storage adapter.
func New(db *pg.DB) *Adapter {
	adapter := new(Adapter)
	adapter.db = db
	return adapter
}

// Select returns set of objects searched by SQL SELECT.
func (adapter *Adapter) Select(
	model interface{},
	limit int,
	condition string,
	params ...interface{},
) error {
	query := adapter.db.Model(model).Where(condition, params...)
	return query.Order("id ASC").Limit(limit).Select()
}

// Update changes state of existing model.
func (adapter *Adapter) Update(
	model interface{},
) error {
	return adapter.db.Update(model)
}

// Insert creates new record in the storage.
func (adapter *Adapter) Insert(
	model interface{},
) error {
	_, err := adapter.db.Model(model).Insert()
	return err
}

// Healthy performs storage health check.
func (adapter *Adapter) Healthy() bool {
	_, err := adapter.db.Exec("SELECT 1")
	return err != nil
}
