package middlewares

import "../../../../../go-api-social-media/app/core/middlewares/net/http"

func SetupHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Definir o cabeçalho Content-Type como application/json
		w.Header().Set("Content-Type", "application/json")

		// Chamar o próximo middleware ou manipulador de rota
		next.ServeHTTP(w, r)
	})
}
