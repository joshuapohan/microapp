package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"

	model "github.com/joshuapohan/microapp/model"
	http_util "github.com/joshuapohan/microapp/util"
)

func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		var authToken string

		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			http_util.RespondWithError(w, http.StatusBadRequest, http_util.Error{Message: "Forbidden"})
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			authToken = bearerToken[1]
		}

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				http_util.RespondWithError(w, http.StatusBadRequest, http_util.Error{Message: "Forbidden"})
				return nil, fmt.Errorf("Error signing method")
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			http_util.RespondWithError(w, http.StatusForbidden, http_util.Error{Message: "Forbidden"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			u.KSUID = claims["ksuid"].(string)
			u.Email = claims["email"].(string)
			ctx := model.NewUserContext(r.Context(), &u)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		} else {
			http_util.RespondWithError(w, http.StatusForbidden, http_util.Error{Message: "Forbidden"})
		}
	}

	return http.HandlerFunc(fn)
}
