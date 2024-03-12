package config

import "github.com/golang-jwt/jwt/v4"

// memmbuat secret key jwt
// dengan key nya byte
var JWT_KEY = []byte("eq2ijo2qe09jo2eqre2jie2rijo")

// menyimpan jwt user
// expired
type JWTClaims struct {
	Username string
	jwt.RegisteredClaims
}
