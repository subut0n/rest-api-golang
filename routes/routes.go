package routes

import (
	"rest-api-golang/controllers"

	"github.com/gorilla/mux"
)

// All routes related to our Documents
func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/documents", controllers.GetDocuments).Methods("GET")
	router.HandleFunc("/documents/{id}", controllers.GetDocumentById).Methods("GET")
	router.HandleFunc("/documents", controllers.CreateDocument).Methods("POST")
	router.HandleFunc("/documents/{id}", controllers.UpdateDocument).Methods("PUT")
	router.HandleFunc("/documents/{id}", controllers.DeleteDocument).Methods("DELETE")
}
