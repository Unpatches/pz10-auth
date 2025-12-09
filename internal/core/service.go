package core

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"

	"example.com/pz10-auth/internal/http/middleware"
	"example.com/pz10-auth/internal/platform/jwt"
)

type userRepo interface {
	CheckPassword(email, pass string) (*User, error)
}

type jwtSigner interface {
	Sign(userID int64, email, role string) (string, error)
}

type Service struct {
	repo       userRepo
	accessJWT  jwt.Validator // для access
	refreshJWT jwt.Validator // для refresh
	blacklist  *RefreshBlacklist
}

func NewService(r userRepo, access, refresh jwt.Validator) *Service {
	return &Service{
		repo:       r,
		accessJWT:  access,
		refreshJWT: refresh,
		blacklist:  NewRefreshBlacklist(),
	}
}

func (s *Service) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		httpError(w, http.StatusBadRequest, "bad_json")
		return
	}

	u, err := s.repo.CheckPassword(in.Email, in.Password)
	if err != nil {
		httpError(w, http.StatusUnauthorized, "invalid_credentials")
		return
	}

	access, err := s.accessJWT.Sign(u.ID, u.Email, u.Role)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "token_error")
		return
	}
	refresh, err := s.refreshJWT.Sign(u.ID, u.Email, u.Role)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "token_error")
		return
	}

	jsonOK(w, map[string]any{
		"access":  access,
		"refresh": refresh,
	})
}

func (s *Service) MeHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.CtxClaimsKey).(map[string]any)
	if !ok || claims == nil {
		httpError(w, http.StatusUnauthorized, "no_claims")
		return
	}

	jsonOK(w, map[string]any{
		"id":    claims["sub"],
		"email": claims["email"],
		"role":  claims["role"],
	})
}

func (s *Service) AdminStats(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]any{"users": 2, "version": "1.0"})

}

var mockUsers = map[int64]User{
	1: {ID: 1, Email: "admin@example.com", Role: "admin"},
	2: {ID: 2, Email: "user@example.com", Role: "user"},
}

func (s *Service) UserByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		httpError(w, http.StatusBadRequest, "bad_id")
		return
	}

	u, ok := mockUsers[id]
	if !ok {
		httpError(w, http.StatusNotFound, "not_found")
		return
	}

	jsonOK(w, u)
}

// утилиты и ключ для контекста — экспортируем из middleware
// type ctxClaims struct{}

func jsonOK(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
func httpError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

type RefreshBlacklist struct {
	mu     sync.RWMutex
	tokens map[string]int64 // token -> exp (Unix)
}

func NewRefreshBlacklist() *RefreshBlacklist {
	return &RefreshBlacklist{
		tokens: make(map[string]int64),
	}
}

func (b *RefreshBlacklist) IsRevoked(token string) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	_, ok := b.tokens[token]
	return ok
}

func (b *RefreshBlacklist) Revoke(token string, exp int64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.tokens[token] = exp
}

func (s *Service) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Refresh string `json:"refresh"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Refresh == "" {
		httpError(w, http.StatusBadRequest, "bad_refresh")
		return
	}

	// 1) Проверяем blacklist
	if s.blacklist.IsRevoked(in.Refresh) {
		httpError(w, http.StatusUnauthorized, "refresh_revoked")
		return
	}

	// 2) Парсим refresh-токен
	claims, err := s.refreshJWT.Parse(in.Refresh)
	if err != nil {
		httpError(w, http.StatusUnauthorized, "invalid_refresh")
		return
	}

	// jwt.MapClaims — это map[string]any, можно брать как обычно
	subF, _ := claims["sub"].(float64)
	id := int64(subF)
	email, _ := claims["email"].(string)
	role, _ := claims["role"].(string)
	expF, _ := claims["exp"].(float64)
	exp := int64(expF)

	// 3) Отзываем старый refresh (кладём в blacklist)
	s.blacklist.Revoke(in.Refresh, exp)

	// 4) Генерируем новую пару
	accessNew, err := s.accessJWT.Sign(id, email, role)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "token_error")
		return
	}
	refreshNew, err := s.refreshJWT.Sign(id, email, role)
	if err != nil {
		httpError(w, http.StatusInternalServerError, "token_error")
		return
	}

	jsonOK(w, map[string]any{
		"access":  accessNew,
		"refresh": refreshNew,
	})
}
