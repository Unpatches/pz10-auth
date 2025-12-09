package middleware

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func AuthZRoles(allowed ...string) func(http.Handler) http.Handler {
	set := map[string]struct{}{}
	for _, a := range allowed {
		set[a] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, _ := r.Context().Value(CtxClaimsKey).(map[string]any)
			role, _ := claims["role"].(string)
			if _, ok := set[role]; !ok {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func AuthZUserSelfOrAdmin() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, _ := r.Context().Value(CtxClaimsKey).(map[string]any)
			role, _ := claims["role"].(string)

			if role == "" {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			// admin — без ограничений
			if role == "admin" {
				next.ServeHTTP(w, r)
				return
			}

			// для user нужно сравнить id из URL и sub из токена
			idStr := chi.URLParam(r, "id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				http.Error(w, "bad id", http.StatusBadRequest)
				return
			}

			subF, _ := claims["sub"].(float64)
			userIDFromToken := int64(subF)

			if role == "user" && id != userIDFromToken {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
