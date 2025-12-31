package main

import (
	"api-management/internal/db"
	"api-management/internal/utils"
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	jwt := utils.NewJWTManager("super-secret")

	ctx := context.Background()

	dbPool := db.NewDB(ctx)
	defer dbPool.Close()

	r := NewRouter(dbPool, jwt)

	log.Println("server running on :8004")
	if err := http.ListenAndServe(":8004", r); err != nil {
		log.Fatal(err)
	}
}
