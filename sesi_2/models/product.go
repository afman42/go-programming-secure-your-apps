package models

import (
	"context"
	"fmt"
	"sesi_2_authentication_middleware/helpers"
	"sesi_2_authentication_middleware/input"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Product struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

func GetOneProductByUserId(db *pgx.Conn, productID uint) (product Product, err error) {
	query := "SELECT user_id FROM products WHERE id = $1 limit 1"
	err = db.QueryRow(context.Background(), query, productID).Scan(&product.UserID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func AllByUserIDProducts(c *gin.Context) (products []Product, err error) {
	userData := c.MustGet("userData").(*helpers.RoleUserClaims)
	query := "SELECT * FROM products WHERE user_id = $1"
	rows, err := c.MustGet("db").(*pgx.Conn).Query(context.Background(), query, userData.ID)
	return fetchAllProduct(rows, err, products)
}

func AllProducts(c *gin.Context) (products []Product, err error) {
	query := "SELECT * FROM products"
	rows, err := c.MustGet("db").(*pgx.Conn).Query(context.Background(), query)

	return fetchAllProduct(rows, err, products)
}

func GetByIdProduct(c *gin.Context, productID uint) (product Product, err error) {
	query := "SELECT * FROM products WHERE id = $1"
	err = c.MustGet("db").(*pgx.Conn).QueryRow(context.Background(), query, productID).Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.UserID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return product, fmt.Errorf("product id %v not found", productID)
	}
	return product, nil
}

func CreateProduct(c *gin.Context, input input.CreateOrUpdateProduct) (product Product, err error) {
	userID := c.MustGet("userData").(*helpers.RoleUserClaims).ID
	query := "INSERT INTO products (title,description,user_id) values($1,$2,$3) returning *"
	err = c.MustGet("db").(*pgx.Conn).QueryRow(
		context.Background(),
		query,
		input.Title,
		input.Description,
		userID,
	).Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.UserID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return product, fmt.Errorf("error creating data : %g", err)
	}
	return product, nil
}

func DeleteProductByID(c *gin.Context, productID uint) (err error) {
	query := "DELETE FROM products WHERE id = $1"
	_, err = c.MustGet("db").(*pgx.Conn).Exec(context.Background(), query, productID)
	if err != nil {
		return fmt.Errorf("error deleting data id %g", err)
	}
	return nil
}

func UpdateProductByID(c *gin.Context, input input.CreateOrUpdateProduct, productID uint) (product Product, err error) {
	sqlStatement := `
		UPDATE products
		SET title = $2, description = $3, updated_at = $4
		WHERE id = $1 RETURNING *;
	`

	err = c.MustGet("db").(*pgx.Conn).QueryRow(
		context.Background(),
		sqlStatement,
		productID,
		input.Title,
		input.Description,
		time.Now()).Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.UserID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return product, fmt.Errorf("error updating data %g", err)
	}

	return product, nil
}

func fetchAllProduct(rows pgx.Rows, err error, products []Product) ([]Product, error) {
	if err != nil {
		return products, fmt.Errorf("error selecting data : %g", err)
	}

	defer rows.Close()

	for rows.Next() {
		var product = Product{}

		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.UserID, &product.CreatedAt, &product.UpdatedAt)

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
