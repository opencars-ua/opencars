package main

import (
	"github.com/go-pg/pg"
	"github.com/opencars-ua/opencars/internal/database"
	"github.com/opencars-ua/opencars/internal/http"
)

type RealDB struct {
	*pg.DB
}

func (r *RealDB) Select(
	model interface{},
	limit int,
	condition string,
	params ...interface{},
) error {
	query := r.DB.Model(model).Where(condition, params...)
	return query.Limit(limit).Select()
}

func main() {
	sql := database.Must(database.DB())
	http.DB = &RealDB{sql}
	http.Run()
	defer sql.Close()
}
