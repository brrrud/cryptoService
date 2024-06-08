package main

import (
	"cryptoService/contollers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/encryptMsg", contollers.EncryptDataHandler)
	router.POST("/decryptMsg", contollers.DecryptDataHandler)
	router.POST("/encryptFile", contollers.EncryptFileFromComputerHandler)
	router.POST("/decryptFile", contollers.DecryptFileFromComputerHandler)

	err := router.Run(":8090")
	if err != nil {
		_ = fmt.Errorf("error with start GO seervice %s", err)
		return
	}
}
