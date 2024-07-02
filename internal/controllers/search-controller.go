package controllers

import (
	"encoding/json"
	"github.com/cruzD21/frag-backend/internal/database"
	"github.com/gorilla/mux"
	"net/http"
)

type SearchOutput struct {
	FragID     int    `json:"frag_id"`
	FragName   string `json:"frag_name"`
	FragDesc   string `json:"frag_description"`
	FragHouse  string `json:"frag_house"`
	FragImgUrl string `json:"frag_img_url"`
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	var fragrances []SearchOutput

	vars := mux.Vars(r)
	query := vars["query"]
	res := database.DBConn.SupabaseRPC(query)                //postgresql rpc call
	if err := json.Unmarshal(res, &fragrances); err != nil { //rpc returned error, bad query

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*	")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//search fragrance table call to database?
//return dior savague + more dior fragrances

//search brand call to database?
//return brand (Dior)

//merge results?
