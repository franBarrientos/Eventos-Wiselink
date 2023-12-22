package main

import (
	"context"
	"github.com/franBarrientos/infrastructure/config"
	"github.com/franBarrientos/infrastructure/gorm"
	"github.com/franBarrientos/infrastructure/rest/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func main() {

	env := config.GetEnv()
	db, err := gorm.InitDbConnection(env)

	if err != nil {
		log.Fatal("Failed to initialize database connection:", err)
		return
	}

	app := fiber.New()
	app.Use(cors.New())
	routes.InitRoutes(app, db, env)

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal("Failed to initialize amqp.Dia:", err)
		return
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to initialize conn.Channel():", err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				body := "Hi I'm Franco"
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				err := ch.PublishWithContext(ctx,
					"",
					q.Name,
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(body),
					})
				if err != nil {
					log.Printf("Failed to publish message: %v", err)
				} else {
					log.Println(" [x] Sent ", body)
				}
			}
		}
	}()

	app.Listen(":" + env.Port)
}
