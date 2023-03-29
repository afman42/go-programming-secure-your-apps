package helpers

import (
	"sesi_2_authentication_middleware/enums"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "abc"

func GenerateToken(id int, email string, role enums.RoleUser) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}
