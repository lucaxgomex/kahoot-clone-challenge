package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var quizColletion *mongo.Collection

func main() {
    app := fiber.New()

	setupDb()

	app.Get("/", index)
    app.Get("/api/quizzes", getQuizzes)

    log.Fatal(app.Listen(":3000"))
}

func setupDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer cancel()

	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	quizColletion = client.Database("quiz").Collection("quizzes")
}

func index(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func getQuizzes(c *fiber.Ctx) error {
	cursor, err := quizColletion.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}

	quizzes := 	[]map[string]any{}

	err = cursor.All(context.Background(), &quizzes)
	if err != nil {
		panic(err)
	}

	return c.JSON(quizzes)
}