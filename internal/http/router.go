package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"example.com/pz10-auth/internal/core"
	"example.com/pz10-auth/internal/http/middleware"
	"example.com/pz10-auth/internal/platform/config"
	"example.com/pz10-auth/internal/platform/jwt"
	"example.com/pz10-auth/internal/repo"
)

func Build(cfg config.Config) http.Handler {
	r := chi.NewRouter()

	// DI
	userRepo := repo.NewUserMem() // храним заранее захэшированных юзеров (email, bcrypt)
	accessJWT := jwt.NewHS256(cfg.JWTSecret, cfg.AccessTTL)
	refreshJWT := jwt.NewHS256(cfg.JWTSecret, cfg.RefreshTTL)

	svc := core.NewService(userRepo, accessJWT, refreshJWT)

	// Публичные маршруты
	r.Post("/api/v1/login", svc.LoginHandler) // выдаёт JWT по email+password
	r.Post("/api/v1/refresh", svc.RefreshHandler)

	// Защищённые маршруты
	r.Group(func(priv chi.Router) {
		priv.Use(middleware.AuthN(accessJWT))            // аутентификация JWT
		priv.Use(middleware.AuthZRoles("admin", "user")) // базовая RBAC
		priv.Get("/api/v1/me", svc.MeHandler)            // вернёт профиль из токена
		priv.With(middleware.AuthZUserSelfOrAdmin()).
			Get("/api/v1/users/{id}", svc.UserByIDHandler)
	})

	// Пример только для админов
	r.Group(func(admin chi.Router) {
		admin.Use(middleware.AuthN(accessJWT))
		admin.Use(middleware.AuthZRoles("admin"))
		admin.Get("/api/v1/admin/stats", svc.AdminStats)
	})

	return r
}
