package middleware

import (
	"github.com/sergejpm/product/internal/domain/service/authorization"
	"github.com/sergejpm/product/internal/infra/log"
	"net/http"
	"strings"
	"sync"
)

const (
	msgMissingAuthHeader       = "request header does not contain an authorization key"
	msgTokenVerificationFailed = "token verification failed"
)

var (
	securePaths = map[string]struct{}{"/api/v1/product/info": {}}
	mu          = &sync.Mutex{}
)

func AuthHandler(h http.Handler, authService *authorization.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuthRequired(r) {
			h.ServeHTTP(w, r)
			return
		}
		authToken := r.Header.Get("Authorization")
		splitToken := strings.Split(authToken, "Bearer ")

		if len(splitToken) < 2 {
			http.Error(w, msgMissingAuthHeader, http.StatusUnauthorized)
			return
		}

		authToken = splitToken[1]

		_, err := authService.Authorize(r.Context(), authToken)
		if err != nil {
			if err.Error() == "unauthorized" {
				http.Error(w, msgTokenVerificationFailed, http.StatusUnauthorized)
				return
			}

			log.Logger().Errorf("error while authorizing user: %v", err)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func isAuthRequired(r *http.Request) bool {
	mu.Lock()
	defer mu.Unlock()
	_, ok := securePaths[strings.ToLower(r.URL.Path)]
	return ok
}
