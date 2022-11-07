package middleware

import (
	"encoding/json"
	"erp/constants"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func checkAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var mySigningKey = []byte(constants.SECRET_KEY)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			json.NewEncoder(w).Encode("Your Token has been expired")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["status"] == "Admin" {
				r.Header.Set("Status", "Admin")
				next.ServeHTTP(w, r)
				return
			} else {
				http.Error(w, http.StatusText(403), 403)
				return
			}
		}
		return
	})
}
