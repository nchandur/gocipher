package api

import (
	"gocipher/crypto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AutokeyHandler(r *gin.Engine) {
	r.POST("/autokey/encrypt", func(ctx *gin.Context) {
		request := struct {
			Plaintext string `json:"plaintext" binding:"required"`
			Key       string `json:"key" binding:"required"`
		}{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		auto := crypto.NewAutokey(request.Key)
		ciphertext := auto.Encrypt(request.Plaintext)

		ctx.JSON(http.StatusOK, gin.H{"ciphertext": ciphertext})

	})

	r.POST("/autokey/decrypt", func(ctx *gin.Context) {
		request := struct {
			Ciphertext string `json:"ciphertext" binding:"required"`
			Key        string `json:"key" binding:"required"`
		}{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		auto := crypto.NewAutokey(request.Key)

		plaintext := auto.Decrypt(request.Ciphertext)

		ctx.JSON(http.StatusOK, gin.H{"plaintext": plaintext})

	})

}
