package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	AtbashHandler(r)
	ROT13Handler(r)
	CaesarHandler(r)
	AffineHandler(r)
	RailFenceHandler(r)
	PolybiusHandler(r)
	SubstitutionHandler(r)
	ColumnarHandler(r)
	AutokeyHandler(r)

	return r

}
