package inventory_service

import (
	"testing"
)

const (
	dsn string = "root@tcp(127.0.0.1:3306)/glive_inventory_development"
)

func TestGet(t *testing.T) {
	repo := NewConn(dsn)
	product, _ := repo.Get("5cd90647-d60c-40df-7bfa-17063328e98a")

	if product.Uuid != "5cd90647-d60c-40df-7bfa-17063328e98a" {
		t.Errorf("Expect Product.uuid eql: db125b76-f3f4-4063-a48d-49dc102190d2 but got: %s", product.Uuid)
	}

	if product.GlobalMax != 1000 {
		t.Errorf("Expect Product.global_max eql: 1000 but got: %d", product.GlobalMax)
	}
}

func BenchmarkGet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		repo := NewConn(dsn)
		repo.Get("5cd90647-d60c-40df-7bfa-17063328e98a")
	}
}
