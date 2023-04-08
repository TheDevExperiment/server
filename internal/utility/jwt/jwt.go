package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type Claims struct {
	jwt.StandardClaims
	CountryCode string
	CityID      string
	IsGuest     bool
}

func CreateToken(id string, countryCode string, cityId string, isGuest bool) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(time.Hour * 876600).Unix()

	// Create token claims
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    viper.GetString("jwt.issuer"),
			Subject:   id,
		},
		CountryCode: countryCode,
		CityID:      cityId,
		IsGuest:     isGuest,
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret
	secret := []byte(viper.GetString("jwt.secret"))
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*Claims, error) {

	var handlerFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Return secret for validation
		return []byte(viper.GetString("jwt.secret")), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, handlerFunc)
	if err != nil {
		return nil, fmt.Errorf("could not parse token: %w", err)
	}

	// Check if token is valid
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
