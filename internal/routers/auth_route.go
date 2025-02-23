package routers

import (
	"github.com/gorilla/mux"
	"csi-accounts/internal/controllers"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/auth", controllers.AuthHandler).Methods("GET")
}
