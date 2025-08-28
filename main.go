package main

import (
	"Pichat/api/routes"
	"Pichat/api/validator"
	"Pichat/pkg/auth"
	"Pichat/pkg/users"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"

	"Pichat/pkg/ent"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s users=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.Background()

	// repository & service
	userRepo := users.NewRepo(client, ctx)
	userService := users.NewService(userRepo)
	authService := auth.NewService(userRepo)

	validator.InitValidator()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	api := app.Group("/api")
	routes.UserRouter(api, userService)

	log.Fatal(app.Listen(":3000"))
}
