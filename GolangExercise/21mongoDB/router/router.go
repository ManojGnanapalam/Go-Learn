package router

import (
	"curdMongoDB/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.CreateMovie).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/movie", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
