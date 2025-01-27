package models

type Fragrance struct {
	FragranceID  int    `json:"fragrance_id"`
	Name         string `json:"name"`
	HouseID      int    `json:"house_id"`
	Description  string `json:"description"`
	FraganticaID string `json:"fraganticaID"`
	FragIMG      string `json:"frag_img"`
	HouseName    string `json:"house_name"`
}

type FragranceHouse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Note struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FragranceToNote struct {
	FragID   int    `json:"fragrance_id"`
	NoteID   int    `json:"note_id"`
	NoteType string `json:"note_type"`
}

type Accord struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FragranceToAccord struct {
	FragID   int `json:"fragrance_id"`
	AccordID int `json:"accord_id"`
}
