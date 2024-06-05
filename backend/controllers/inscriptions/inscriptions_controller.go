package inscriptions

import (
	dtoInscription "backend/dto/inscription"
	inscriptions "backend/services/inscriptions"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserInscription is a function that allows to register a user in a course

func UserInscription(context *gin.Context) {
	var inscriptionRequest dtoInscription.InscriptionDto
	context.BindJSON(&inscriptionRequest)
	response, err := inscriptions.UserInscription(inscriptionRequest)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusCreated, response)
}
