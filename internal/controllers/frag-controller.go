package controllers

import (
	"encoding/json"
	"github.com/cruzD21/frag-backend/internal/database"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetFragrances(w http.ResponseWriter, r *http.Request) {
	fragrances, err := database.DBConn.GetAllFragrances()
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(fragrances)
	w.Header().Set("Access-Control-Allow-Origin", "*	")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFragranceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fragID := vars["fragID"]
	_, err := strconv.ParseInt(fragID, 0, 0)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}
	fragrance, err := database.DBConn.GetFragranceByID(fragID)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(fragrance)

	w.Header().Set("Access-Control-Allow-Origin", "*	")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFragNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fragID := vars["fragID"]
	res := database.DBConn.GetFragranceNotesRPC(fragID)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*	")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFragAccords(w http.ResponseWriter, r *http.Request) {

	res := database.DBConn.GetFragranceAccordRPC()
	//if err := json.Unmarshal(res, &fragrances); err != nil { //rpc returned error, bad query
	//
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*	")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
