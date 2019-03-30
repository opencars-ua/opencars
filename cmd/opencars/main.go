package main

import (
	"github.com/go-pg/pg"
	"github.com/opencars-ua/opencars/internal/http"
	"github.com/opencars-ua/opencars/internal/sql"
)

type RealDB struct {
	*pg.DB
}

func (r *RealDB) SelectWhere(model interface{}, limit int, condition string, params ...interface{}) error {
	query := r.DB.Model(model).Where(condition, params...)
	return query.Limit(limit).Select()
}

func main() {
	http.DB = &RealDB{sql.Must(sql.DB())}
	http.Run()
}
