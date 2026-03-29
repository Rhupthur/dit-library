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
	h := NewUtilisateurHandler(db)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "utilisateurs"})
	})

	utilisateurs := r.Group("/utilisateurs")
	{
		utilisateurs.GET("", h.GetAll)
		utilisateurs.GET("/:id", h.GetOne)
		utilisateurs.POST("", h.Create)
		utilisateurs.PATCH("/:id", h.Update)
		utilisateurs.DELETE("/:id", h.Delete)
	}

	r.Run(":8002")
}
