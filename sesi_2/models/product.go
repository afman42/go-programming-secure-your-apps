package models

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Product struct {
	GormModel
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	User        User
}

func GetOneProductByUserId(db *pgx.Conn, productID uint) (product Product, err error) {
	query := "SELECT user_id FROM products id = $1 limit 1"
	err = db.QueryRow(context.Background(), query, productID).Scan(&product.ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func AllProducts(c *gin.Context) (products []Product, err error) {
	query := "SELECT * FROM product"
	rows, err := c.MustGet("db").(*pgx.Conn).Query(context.Background(), query)
	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product = Product{}

		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.CreatedAt, &product.UpdatedAt)

		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return products, err
	}

	return products, nil
}
