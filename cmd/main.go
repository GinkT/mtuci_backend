package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ginkt/mtuci_backend/config"
	"github.com/ginkt/mtuci_backend/internal/db"
	"github.com/ginkt/mtuci_backend/internal/handlers"
	log "github.com/sirupsen/logrus"
)

var ctx = context.TODO()

func main() {
	cfg := config.InitConfig()
	dbClient, err := db.CreateMongoClient(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	handler := handlers.CreateHandler(dbClient)

	router := gin.Default()
	router.POST("/api/v1/create-user", handler.CreateNewUserHandler)
	router.POST("/api/v1/create-ride", handler.CreateNewRideHandler)
	router.POST("/api/v1/add-point", handler.AddNewPointHandler)

	log.Fatalln(router.Run(":8080"))
}

