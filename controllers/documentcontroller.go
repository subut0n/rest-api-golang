package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rest-api-golang/database"
	"github.com/rest-api-golang/entities"

	"github.com/gorilla/mux"
)

//CREATE | Create a new document
func CreateDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var document entities.Document
	json.NewDecoder(r.Body).Decode(&document)
	fmt.Printf("Request body : %v", r.Body) // test
	database.Instance.Create(&document)
	json.NewEncoder(w).Encode(document)
	fmt.Printf("Document content: %v", document) //test
}

// Used for READ UPDATE and DELETE Operations
// READ: GetDocumentById(...)
// UPDATE: UpdateDocument(...)
// DELETE: DeleteDocument(...)
func checkIfDocumentExists(documentId string) bool {
	var document entities.Document

	database.Instance.First(&document, documentId)
	return document.ID != 0 // if docId=0 ? return False : return True

}

// GET BY ID | Get document details by document ID
func GetDocumentById(w http.ResponseWriter, r *http.Request) {
	documentId := mux.Vars(r)["id"] // MUX extracts the id from the URL and assigns the value to the id variable
	if !checkIfDocumentExists(documentId) {
		json.NewEncoder(w).Encode("Document not found!")
		return
	}

	var document entities.Document
	database.Instance.First(&document, documentId) // fill the document details to the variable 'document'
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(document)
}

// GET ALL | Get all documents details
func GetDocuments(w http.ResponseWriter, r *http.Request) {
	var documents []entities.Document
	database.Instance.Find(&documents)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(documents)
}

// UPDATE
func UpdateDocument(w http.ResponseWriter, r *http.Request) {
	documentId := mux.Vars(r)["id"]
	if !checkIfDocumentExists(documentId) {
		json.NewEncoder(w).Encode("Document not found!")
		return
	}

	var document entities.Document
	database.Instance.First(&document, documentId) // GORM queries the document record to document variable
	json.NewDecoder(r.Body).Decode(&document)      // Decoder converts the request body to document variable
	database.Instance.Save(&document)              // Saved to the database table
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(document)
}

// DELETE | Delete documents by passing-in the document id
func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	documentId := mux.Vars(r)["id"]
	if !checkIfDocumentExists(documentId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Document not found!")
		return
	}

	var document entities.Document
	database.Instance.Delete(&document, documentId) // GORM deletes the document by ID
	json.NewEncoder(w).Encode("Document deleted successfully!")
}
