package app

import (
	"github.com/rmsubekti/ghapis/app/utils"
	"net/http"
)

// NotFoundHandler exception
var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		utils.Respond(w, utils.Message(false, "Resource not found"))
		next.ServeHTTP(w, r)
	})
}
