package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ginkt/mtuci_backend/internal/db"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler)AddNewPointHandler(ctx *gin.Context) {
	var point db.PointRecord
	if err := ctx.ShouldBindJSON(&point); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.WithField("error", err.Error()).Infoln("Interrupting add point request")
		return
	}

	point.Id = primitive.NewObjectID()
	insertedId, err := db.AddNewPointRecordToPoints(h.mongo, context.TODO(), point)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.WithField("error", err.Error()).Infoln("Interrupting add point request")
		return
	}

	log.Infoln("Point was successfully added", point.Timestamp, insertedId.Hex())
	ctx.JSON(http.StatusOK, point)
}
