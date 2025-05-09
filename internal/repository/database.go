package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func SetClient() (*supabase.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	projectUrl := os.Getenv("SUPABASE_URL")
	anonKey := os.Getenv("SUPABASE_ANON_KEY")

	client, err := supabase.NewClient(projectUrl, anonKey, nil)
	if err != nil {
		fmt.Println("cannot initialise client", err)
		return nil, err
	} else {
		fmt.Println("successfully connected to database!")
	}

	return client, nil
}
