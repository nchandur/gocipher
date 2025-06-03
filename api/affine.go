package api

import (
	"gocipher/crypto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AffineHandler(r *gin.Engine) {
	r.POST("/affine/encrypt", func(ctx *gin.Context) {
		request := struct {
			Plaintext string `json:"plaintext" binding:"required"`
			Scale     int    `json:"scale" binding:"required"`
			Shift     int    `json:"shift" binding:"required"`
		}{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		affine := crypto.NewAffine(request.Scale, request.Shift)

		ciphertext := affine.Encrypt(request.Plaintext)

		ctx.JSON(http.StatusOK, gin.H{"ciphertext": ciphertext})

	})

	r.POST("/affine/decrypt", func(ctx *gin.Context) {
		request := struct {
			Ciphertext string `json:"ciphertext" binding:"required"`
			Scale      int    `json:"scale" binding:"required"`
			Shift      int    `json:"shift" binding:"required"`
		}{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		affine := crypto.NewAffine(request.Scale, request.Shift)

		plaintext := affine.Decrypt(request.Ciphertext)

		ctx.JSON(http.StatusOK, gin.H{"plaintext": plaintext})

	})

}
