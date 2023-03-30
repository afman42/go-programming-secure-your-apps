package helpers

import (
	"sesi_2_authentication_middleware/enums"

	"github.com/dgrijalva/jwt-go"
)

type RoleUserClaims struct {
	jwt.StandardClaims
	ID    uint           `json:"id,omitempty"`
	Email string         `json:"email,omitempty"`
	Role  enums.RoleUser `json:"role,omitempty"`
}

var secretKey = "abc"

func GenerateToken(id uint, email string, role enums.RoleUser) string {
	claims := RoleUserClaims{
		ID:    id,
		Email: email,
		Role:  role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}
