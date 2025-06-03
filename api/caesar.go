package api

import (
	"gocipher/crypto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CaesarHandler(r *gin.Engine) {
	r.POST("/caesar/encrypt", func(ctx *gin.Context) {
		request := struct {
			Plaintext string `json:"plaintext" binding:"required"`
			Shift     int    `json:"shift" binding:"required"`
		}{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		caesar := crypto.NewCaesar(request.Shift)

		ciphertext := caesar.Encrypt(request.Plaintext)

		ctx.JSON(http.StatusOK, gin.H{"ciphertext": ciphertext})

	})

	r.POST("/caesar/decrypt", func(ctx *gin.Context) {
		request := struct {
			Ciphertext string `json:"ciphertext" binding:"required"`
			Shift      int    `json:"shift" binding:"required"`
		}{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		caesar := crypto.NewCaesar(request.Shift)

		plaintext := caesar.Decrypt(request.Ciphertext)

		ctx.JSON(http.StatusOK, gin.H{"plaintext": plaintext})

	})

}
