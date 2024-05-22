package middlewares

import (
	auth_token "api-social-media/app/core/helpers/auth"
	responses "api-social-media/app/core/helpers/response"
	"net/http"
)

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth_token.ValidateToken(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
