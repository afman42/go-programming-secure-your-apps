package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// https://github.com/dgrijalva/jwt-go/issues/403

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.ParseWithClaims(stringToken, &RoleUserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	// fmt.Println(token.Claims.(*RoleUserClaims))
	if _, ok := token.Claims.(*RoleUserClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(*RoleUserClaims), nil
}
