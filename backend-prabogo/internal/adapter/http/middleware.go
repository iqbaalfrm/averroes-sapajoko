package http

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	ContextUserID contextKey = "id_pengguna"
	ContextRole   contextKey = "peran"
)

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")
			if authorization == "" {
				ResponGagal(w, http.StatusUnauthorized, "Token tidak ditemukan", "harap login terlebih dahulu")
				return
			}
			parts := strings.SplitN(authorization, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				ResponGagal(w, http.StatusUnauthorized, "Format token tidak valid", "gunakan Bearer token")
				return
			}

			tokenStr := parts[1]
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				ResponGagal(w, http.StatusUnauthorized, "Token tidak valid", "silakan login ulang")
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				ResponGagal(w, http.StatusUnauthorized, "Token tidak valid", "silakan login ulang")
				return
			}

			if exp, ok := claims["exp"].(float64); ok {
				expTime := time.Unix(int64(exp), 0)
				if time.Now().After(expTime) {
					ResponGagal(w, http.StatusUnauthorized, "Token kadaluarsa", "silakan login ulang")
					return
				}
			}

			idFloat, ok := claims["id_pengguna"].(float64)
			if !ok {
				ResponGagal(w, http.StatusUnauthorized, "Token tidak valid", "silakan login ulang")
				return
			}

			role, _ := claims["peran"].(string)
			ctx := context.WithValue(r.Context(), ContextUserID, int64(idFloat))
			ctx = context.WithValue(ctx, ContextRole, role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
