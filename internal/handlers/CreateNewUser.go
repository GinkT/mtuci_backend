package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ginkt/mtuci_backend/internal/db"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)


func (h *Handler) CreateNewUserHandler(ctx *gin.Context) {
	var user db.UserRecord
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.WithField("error", err.Error()).Infoln("Interrupting create user request")
		return
	}

	user.Id = primitive.NewObjectID()
	resultId, err := db.InsertNewUserRecordInUsers(h.mongo, context.TODO(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error()})
		log.WithField("error", err.Error()).Infoln("Interrupting create user request")
		return
	}

	log.Infoln("Successfully created new user!", user.Name, resultId.Hex())
	ctx.JSON(http.StatusOK, user)
}
