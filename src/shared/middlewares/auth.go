package middlewares

import (
	"net/http"
	"os"

	"github.com/eoussama/anusic-api/src/shared/models"
	"github.com/eoussama/anusic-api/src/shared/utils"
)

// Auth checks if the correct access token was passed before resuming the request
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Getting the access token
		token := r.Header.Get("x-access-token")

		// Checking the validity of the token and acting accordingly
		if token == os.Getenv("SECRET") {
			next.ServeHTTP(w, r)
		} else {
			utils.ReturnResponse(w, r, nil, models.Error.InvalidAccessToken(models.Error{}), http.StatusUnauthorized)
		}
	})
}
