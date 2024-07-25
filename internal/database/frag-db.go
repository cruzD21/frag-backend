package database

import (
	"encoding/json"
	"github.com/cruzD21/frag-backend/internal/models"
	"log"
)

type Note struct {
	Name string `json:"name"`
}

type FragranceNotes struct {
	TopNotes    []Note `json:"top_notes"`
	MiddleNotes []Note `json:"middle_notes"`
	BaseNotes   []Note `json:"base_notes"`
}

func (db *DB) GetAllFragrances() ([]models.Fragrance, error) {
	var err error
	var res []models.Fragrance

	data, _, err := db.Supabase.
		From("fragrance").
		Select("*", "1", false).
		Execute()
	if err != nil {
		log.Printf("db.Supabase.From('fragrance').Select(*) ->  %v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Printf("json.Unmarshal -> %v", err)
		return nil, err
	}
	return res, nil
}

func (db *DB) GetFragranceByID(ID string) ([]models.Fragrance, error) {
	var err error
	var res []models.Fragrance
	//id := strconv.FormatInt(ID, 0)

	data, _, err := db.Supabase.
		From("fragrance").
		Select("*", "1", false).
		Eq("fragrance_id", ID).
		Execute()
	if err != nil {
		log.Printf("error selecting fragrance -> %v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Printf("json.Unmarshal -> %v", err)
		return nil, err
	}
	return res, nil
}

func (db *DB) SupabaseRPC(query string) []byte {
	str := db.Supabase.
		Rpc("search_fragrances8", "0", map[string]interface{}{"searchterm": query})
	return []byte(str)
}

func (db *DB) GetFragranceNotesRPC(fid string) []byte {

	str := db.Supabase.
		Rpc("get_fragrance_notes3", "0", map[string]interface{}{"fid": fid})
	return []byte(str)
}

func (db *DB) GetFragranceAccordRPC() []byte {
	str := db.Supabase.
		Rpc("get_fragrance_details3", "0", map[string]interface{}{"fid": 595})
	return []byte(str)
}
