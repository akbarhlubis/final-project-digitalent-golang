package helpers

import "github.com/golang-jwt/jwt/v5"

var secretKey = "secret"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := parseToken.SignedString([]byte(secretKey))

	return token
}
