package contollers

import (
	"cryptoService/cryptography/utils"
	"cryptoService/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EncryptDataHandler(c *gin.Context) {
	var request models.Message
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pad := utils.GetPaddingByName(request.Padding)
	algo := utils.GetAlgorithmByName(request.CryptoAlgorithm, utils.StringToByteArray(request.Key))
	mod, _ := utils.GetModeByName(request.CipherMode, pad, algo)
	encrypted, err := mod.Encrypt(utils.StringToByteArray(request.Content))

	if err != nil {
		fmt.Println(err)
		return
	}

	request.Content = utils.BytesToString(encrypted)

	c.JSON(http.StatusOK, gin.H{
		"message": "File data received successfully",
		"data":    request,
	})
}

func DecryptDataHandler(c *gin.Context) {
	var request models.Message
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pad := utils.GetPaddingByName(request.Padding)
	algo := utils.GetAlgorithmByName(request.CryptoAlgorithm, utils.StringToByteArray(request.Key))
	mod, _ := utils.GetModeByName(request.CipherMode, pad, algo)
	decrypted, err := mod.Decrypt(utils.StringToByteArray(request.Content))

	if err != nil {
		fmt.Println(err)
		return
	}

	request.Content = utils.BytesToString(decrypted)

	c.JSON(http.StatusOK, gin.H{
		"message": "File data received successfully",
		"data":    request,
	})
}

func EncryptFileFromComputerHandler(c *gin.Context) {
	var request models.Message
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var resp models.MessageLikeFile

	pad := utils.GetPaddingByName(request.Padding)
	algo := utils.GetAlgorithmByName(request.CryptoAlgorithm, utils.StringToByteArray(request.Key))
	mod, _ := utils.GetModeByName(request.CipherMode, pad, algo)

	plaintext, err := utils.ReadFileToBytes(request.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	encrypted, err := mod.Encrypt(plaintext)

	resp.Format = utils.GetFileExtension(request.Content)
	resp.Content = utils.BytesToString(encrypted)
	c.JSON(http.StatusOK, gin.H{
		"message": "File data received successfully",
		"data":    resp,
	})
}

func DecryptFileFromComputerHandler(c *gin.Context) {
	var request models.MessageLikeFile
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pad := utils.GetPaddingByName(request.Padding)
	algo := utils.GetAlgorithmByName(request.CryptoAlgorithm, utils.StringToByteArray(request.Key))
	mod, _ := utils.GetModeByName(request.CipherMode, pad, algo)

	encrypted, err := utils.ReadFileToBytes(request.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	decrypted, err := mod.Decrypt(encrypted)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	decryptedText := utils.BytesToString(decrypted)

	request.Content = decryptedText
	c.JSON(http.StatusOK, gin.H{
		"message": "File data received successfully",
		"data":    request,
	})
}
