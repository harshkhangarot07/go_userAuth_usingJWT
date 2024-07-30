package middleware

import (
	"context"
	"go_userAuth/internal/auth"
	"net/http"
	"strings"
)

type ContextKey string

const ClaimsContextKey ContextKey = "claims"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add claims to the request context
		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
