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
	h := NewEmpruntHandler(db)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "emprunts"})
	})

	emprunts := r.Group("/emprunts")
	{
		emprunts.GET("", h.GetAll)
		emprunts.POST("", h.Emprunter)
		emprunts.POST("/:id/retour", h.Retourner)
		emprunts.GET("/historique/:id", h.Historique)
	}

	r.Run(":8003")
}
