package api

import (
	"gocipher/crypto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PolybiusHandler(r *gin.Engine) {
	r.POST("/polybius/encrypt", func(ctx *gin.Context) {
		request := struct {
			Alphabet  string `json:"alphabet" binding:"required"`
			Key       string `json:"key" binding:"required"`
			Chars     string `json:"chars" binding:"required"`
			Plaintext string `json:"plaintext" binding:"required"`
		}{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		if len(request.Alphabet) == 0 || len(request.Key) == 0 || len(request.Chars) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no field must have length 0"})
			return
		}

		poly := crypto.NewPolybius(request.Alphabet, request.Key, request.Chars)

		ciphertext := poly.Encrypt(request.Plaintext)

		ctx.JSON(http.StatusOK, gin.H{"ciphertext": ciphertext})

	})

	r.POST("/polybius/decrypt", func(ctx *gin.Context) {
		request := struct {
			Alphabet   string `json:"alphabet" binding:"required"`
			Key        string `json:"key" binding:"required"`
			Chars      string `json:"chars" binding:"required"`
			Ciphertext string `json:"ciphertext" binding:"required"`
		}{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		if len(request.Alphabet) == 0 || len(request.Key) == 0 || len(request.Chars) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no field must have length 0"})
			return
		}

		poly := crypto.NewPolybius(request.Alphabet, request.Key, request.Chars)

		plaintext := poly.Decrypt(request.Ciphertext)

		ctx.JSON(http.StatusOK, gin.H{"plaintext": plaintext})

	})

}
