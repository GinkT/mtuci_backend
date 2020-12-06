package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ginkt/mtuci_backend/internal/db"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) CreateNewRideHandler(ctx *gin.Context) {
	var ride db.RideRecord
	if err := ctx.ShouldBindJSON(&ride); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.WithField("error", err.Error()).Infoln("Interrupting create ride request")
		return
	}

	exists, err := db.IsUserExists(h.mongo, ctx, ride.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error()})
		log.WithField("error", err.Error()).Infoln("Interrupting create user request")
		return
	}
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{ "error": "userId"})
		log.WithField("error", err.Error()).Infoln("Interrupting create user request")
		return
	}

	ride.Id = primitive.NewObjectID()
	insertedId, err := db.InsertNewRideRecordInRides(h.mongo, context.TODO(), ride)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error()})
		log.WithField("error", err.Error()).Infoln("Interrupting create user request")
		return
	}

	log.Infoln("Successfully created new ride!", ride.Position, insertedId.Hex())
	ctx.JSON(http.StatusOK, ride)
}
