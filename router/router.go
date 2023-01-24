package router

import (
	"UserAuth/controller"
	"UserAuth/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	authController := controller.AuthController{}
	userController := controller.UserController{}

	authRouter := router.PathPrefix("/api/v1").Subrouter()
	authRouter.HandleFunc("/authenticate", authController.GenerateToken).Methods("POST")

	userRouter := router.PathPrefix("/api/v1").Subrouter()
	userRouter.Use(middleware.CheckAuthentication)

	userRouter.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	//router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	//router.HandleFunc("/api/movie/{id}", controllers.DeleteAMovie).Methods("DELETE")
	//router.HandleFunc("/api/deleteallmovie", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}
