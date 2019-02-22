package app

import (
	"net/http"
	"os"
	"strings"

	"github.com/rmsubekti/ghapis/app/models"
	"github.com/rmsubekti/ghapis/app/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

// Xclaim context
var Xclaim = &models.Token{}

// JwtAuthentication middleware
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noAuth := []string{"/v1/user/new", "/v1/user/login", "/v1/product/all"}
		requestPath := r.URL.Path

		//skip authentication for request path that in whitelist
		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		//The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		response := make(map[string]interface{})
		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			response = utils.Message(false, "Access Restricted, Missing Auth key")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		splitted := strings.Split(bearerToken, " ")
		if len(splitted) != 2 {
			response = utils.Message(false, "Access restricted, Invalid key")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		tokenPart := splitted[1]
		claim := &models.Token{}
		token, err := jwt.ParseWithClaims(tokenPart, claim, func(token *jwt.Token) (i interface{}, e error) {
			return []byte(os.Getenv("TOKEN_PASWORD")), nil
		})

		if err != nil {
			response = utils.Message(false, "Access restricted, Malformed key")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		if !token.Valid {
			response = utils.Message(false, "Invalid key, Server not recognize the key")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}
		//all passed
		//ctx := context.WithValue(r.Context(),"claim",claim)
		//r.WithContext(ctx)
		context.Set(r, Xclaim, claim)
		next.ServeHTTP(w, r)
	})
}
