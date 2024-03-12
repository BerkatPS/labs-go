package middleware

import (
	"github.com/BerkatPS/go-jwt-testing/config"
	"github.com/BerkatPS/go-jwt-testing/helper"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

// dia akan menerima parameter http handler dan mengembalikan http handler

func JWTMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Silahkan Login Terlebih Dahulu!!!"}
				helper.Writejson(w, http.StatusUnauthorized, response)
				return
			}
		}
		// jika ada kita mengambil value token tersebut.

		value := cookie.Value

		// kita gunakan type jwt claims yang ada di config jwt
		claims := &config.JWTClaims{}

		token, err := jwt.ParseWithClaims(value, claims, func(token *jwt.Token) (interface{}, error) {
			// kita mengembalikan Config dari jwt key dan juga nil
			return config.JWT_KEY, nil
		})

		if err != nil {
			// mengambil validation error terlebih dahulu.
			validationError, _ := err.(*jwt.ValidationError)
			// kita switch errornya
			switch validationError {

			case jwt.ErrSignatureInvalid:
				response := map[string]string{"message": "Unauthorized"}
				helper.Writejson(w, http.StatusUnauthorized, response)
				return
			case jwt.ErrTokenExpired:
				response := map[string]string{"message": "Token Expired!"}
				helper.Writejson(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unauthorized"}
				helper.Writejson(w, http.StatusUnauthorized, response)
				return
			}
		}
		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helper.Writejson(w, http.StatusUnauthorized, response)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
