package productdb

import (
	"github.com/google/uuid"
	"github.com/ruhollahh/go-ecom/internal/entity/product"
	"time"
)

type dbProduct struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	ImageID     uuid.UUID `db:"image_id"`
	Price       int       `db:"price"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func toDBProduct(ent product.Product) dbProduct {
	db := dbProduct{
		ID:          ent.ID,
		Name:        ent.Name,
		Description: ent.Description,
		ImageID:     ent.ImageID,
		Price:       ent.Price,
		CreatedAt:   ent.CreatedAt.UTC(),
		UpdatedAt:   ent.UpdatedAt.UTC(),
	}

	return db
}

func toProduct(db dbProduct) product.Product {
	ent := product.Product{
		ID:          db.ID,
		Name:        db.Name,
		Description: db.Description,
		ImageID:     db.ImageID,
		Price:       db.Price,
		CreatedAt:   db.CreatedAt.In(time.Local),
		UpdatedAt:   db.UpdatedAt.In(time.Local),
	}

	return ent
}

func toProducts(dbs []dbProduct) []product.Product {
	bus := make([]product.Product, len(dbs))

	for i, db := range dbs {
		bus[i] = toProduct(db)
	}

	return bus
}
