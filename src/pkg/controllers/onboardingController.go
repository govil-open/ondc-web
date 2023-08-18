package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openbank-ondc-web/src/pkg/crypto"
	"github.com/openbank-ondc-web/src/pkg/models"
	"github.com/openbank-ondc-web/src/pkg/utils"
)

func OnSubscribe(context *gin.Context) {
	var request models.OnSubscribe

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	onSubscribeReq := models.OnSubscribe{
		SubscriberId: request.SubscriberId,
		Challenge:    request.Challenge,
	}

	decryptedText, err := decryptChallenge(onSubscribeReq.Challenge)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, &decryptedText)
}

func decryptChallenge(challenge string) (models.OnSubscribeResponse, error) {
	privateKey, err := utils.ReadPEMFileToString("/Users/govil.kumar/Projects/openbank-ondc-web/src/resources/encryptionkey.pem")
	if err != nil {
		return models.OnSubscribeResponse{}, err
	}

	publicKey := utils.ONDC_PUB_ENCRYPTION_KEY_PRE_PROD
	//Remove below code that is only for testing purpose
	// publicKey, err := utils.ReadPEMFileToString("/Users/govil.kumar/Projects/openbank-ondc-web/src/resources/encryption-pub.pem")
	// if err != nil {
	// 	return models.OnSubscribeResponse{}, err
	// }
	// encryptedText, err := crypto.Encrypt(privateKey, publicKey, challenge)
	// if err != nil {
	// 	return models.OnSubscribeResponse{}, err
	// }

	decryptedText, err := crypto.Decrypt(privateKey, publicKey, challenge)
	if err != nil {
		return models.OnSubscribeResponse{}, err
	}

	return models.OnSubscribeResponse{
		Answer: decryptedText,
	}, nil
}
