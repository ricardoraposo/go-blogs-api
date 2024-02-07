package middlewares

import (
	"context"
	"net/http"

	"github.com/ricardoraposo/blogs-api-go/internal/auth"
)

func VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if len(tokenStr) == 0 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := auth.VerifyToken(tokenStr)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
