package middlewear

import (
	. "WDB/endpoints"
	"WDB/views"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
)

func (m *Middlewear) CheckJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims := &Claims{}
		jwtToken := r.Header.Values("Authorization")

		token, err := jwt.ParseWithClaims(jwtToken[0][strings.Index(jwtToken[0], " ")+1:], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("secret")), nil
		})

		if err != nil {
			fmt.Printf("%v\n", err)
			views.ResponseWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		if !token.Valid {
			views.ResponseWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		fmt.Printf("%v --->  %v\n", token, claims)
		if !claims.IsAdmin {
			views.ResponseWithError(w, http.StatusUnauthorized, "not an admin")
			return
		}

		next(w, r)
	}
}
