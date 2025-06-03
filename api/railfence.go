package api

import (
	"gocipher/crypto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RailFenceHandler(r *gin.Engine) {
	r.POST("/railfence/encrypt", func(ctx *gin.Context) {
		request := struct {
			Plaintext string `json:"plaintext" binding:"required"`
			Rails     int    `json:"rails" binding:"required"`
		}{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		rail := crypto.NewRailFence(request.Rails)

		ciphertext := rail.Encrypt(request.Plaintext)

		ctx.JSON(http.StatusOK, gin.H{"ciphertext": ciphertext})

	})

	r.POST("/railfence/decrypt", func(ctx *gin.Context) {
		request := struct {
			Ciphertext string `json:"ciphertext" binding:"required"`
			Rails      int    `json:"rails" binding:"required"`
		}{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		rail := crypto.NewRailFence(request.Rails)

		plaintext := rail.Decrypt(request.Ciphertext)

		ctx.JSON(http.StatusOK, gin.H{"plaintext": plaintext})

	})

}
