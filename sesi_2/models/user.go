package models

import (
	"context"
	"errors"
	"sesi_2_authentication_middleware/enums"
	"sesi_2_authentication_middleware/helpers"
	"sesi_2_authentication_middleware/input"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type User struct {
	BaseModel
	Email    string
	Password string
	Role     enums.RoleUser `sql:"type:enum('admin','user')"`
}

func RegisterUser(input input.RegisterUser, c *gin.Context) (user User, err error) {
	passwordHash := helpers.HashPass(input.Password)
	query := "INSERT INTO users (email,password,role) values($1,$2,$3) returning id, email"
	err = c.MustGet("db").(*pgx.Conn).QueryRow(context.Background(), query, input.Email, passwordHash, enums.RoleUser(input.Role)).Scan(&user.ID, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func LoginUser(input input.LoginUser, c *gin.Context) (user User, err error) {
	query := "SELECT id,email,password,role FROM users WHERE email = $1"
	err = c.MustGet("db").(*pgx.Conn).QueryRow(context.Background(), query, input.Email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return user, errors.New("invalid email or password")
	}
	return user, nil
}
