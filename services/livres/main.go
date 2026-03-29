package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Pas de fichier .env, utilisation des variables d'environnement système")
	}

	db := ConnectDB()
	h := NewLivreHandler(db)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "livres"})
	})

	livres := r.Group("/livres")
	{
		livres.GET("", h.GetAll)
		livres.GET("/:id", h.GetOne)
		livres.POST("", h.Create)
		livres.PATCH("/:id", h.Update)
		livres.DELETE("/:id", h.Delete)
	}

	r.Run(":8001")
}
