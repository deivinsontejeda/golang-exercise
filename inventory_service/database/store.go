package inventory_service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type ProductsService interface {
	Get(uuid string) (*Product, error)
}

type Product struct {
	Uuid      string
	GlobalMax int64
}

type dbStore struct{ db *sql.DB }

func NewConn(dsn string) *dbStore {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}
	return &dbStore{db: db}
}

func (s *dbStore) Get(uuid string) (*Product, error) {
	var product = &Product{}
	rows, err := s.db.Query("SELECT uuid, global_max FROM products WHERE uuid=? LIMIT 1", uuid)

	if err != nil {
		log.Fatal(err.Error())
	}
	for rows.Next() {
		if err := rows.Scan(&product.Uuid, &product.GlobalMax); err != nil {
			log.Fatal(err.Error())
		}
	}
	return product, err
}
