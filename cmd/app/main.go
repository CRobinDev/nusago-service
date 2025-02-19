package main

import (
	"log"
	"os"
	"github.com/CRobinDev/nusago-service/config"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load env variables: %v", err)
	}
}

func main() {
	db, err := config.NewDB()
	if err != nil {
		log.Fatalf("message: %v", err)
	}

	err = config.Migrate(db)
	if err != nil {
		log.Fatalf("message: %v", err)
	}
	
	app := config.NewFiber()
	config.StartApp(&config.AppConfig{
		DB:  db,
		App: app,
	})
   
	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
