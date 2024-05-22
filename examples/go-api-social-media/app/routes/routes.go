package routes

import (
	"api-social-media/app/internal/auth"
	"api-social-media/app/internal/followers"
	"api-social-media/app/internal/users"

	"github.com/gorilla/mux"
)

func SetupAppRoutes() *mux.Router {
	// users.SetupUserRoutes(router)
	router := mux.NewRouter()
	auth.SetupAuthRoutes(router)
	users.SetupUserRoutes(router)
	followers.SetupFollowersRoutes(router)
	return router
}
