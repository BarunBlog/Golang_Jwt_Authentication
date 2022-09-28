package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	// The StandardClaims type is designed to be embedded into your custom types to provide standard validation features.
	jwt.StandardClaims
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey) // Now we can generate a tokenString using the secret_key
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims( // parse the JWT into claims
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim) // From the parsed token, we extract the claims
	if !ok {
		err = errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
