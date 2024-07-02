package database

import (
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
	"os"
)

type DB struct {
	Supabase *supabase.Client
}

var DBConn = &DB{}

func init() {
	_ = godotenv.Load()
	dbApiUrl := os.Getenv("DB_API_URL")
	dbApiKey := os.Getenv("DB_API_KEY")
	client, _ := supabase.NewClient(dbApiUrl, dbApiKey, nil)

	DBConn.Supabase = client
}
