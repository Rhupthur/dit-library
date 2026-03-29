package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

func (h *LivreHandler) GetAll(c *gin.Context) {
	var livres []Livre

	query := h.DB
	if titre := c.Query("titre"); titre != "" {
		query = query.Where("titre ILIKE ?", "%"+titre+"%")
	}
	if auteur := c.Query("auteur"); auteur != "" {
		query = query.Where("auteur ILIKE ?", "%"+auteur+"%")
	}
	if isbn := c.Query("isbn"); isbn != "" {
		query = query.Where("isbn = ?", isbn)
	}

	if err := query.Find(&livres).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur récupération"})
		return
	}

	c.JSON(http.StatusOK, livres)
}

func (h *LivreHandler) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var livre Livre
	if err := h.DB.First(&livre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livre non trouvé"})
		return
	}

	c.JSON(http.StatusOK, livre)
}

func (h *LivreHandler) Create(c *gin.Context) {
	var input LivreCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	livre := Livre{
		Titre:  input.Titre,
		Auteur: input.Auteur,
		ISBN:   input.ISBN,
	}

	if input.Disponible != nil {
		livre.Disponible = *input.Disponible
	} else {
		livre.Disponible = true
	}

	if input.Quantite != nil {
		livre.Quantite = *input.Quantite
	} else {
		livre.Quantite = 1
	}

	if err := h.DB.Create(&livre).Error; err != nil {
		if isUniqueViolation(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "Un livre avec cet ISBN existe déjà"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur création"})
		return
	}

	c.JSON(http.StatusCreated, livre)
}

func (h *LivreHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var livre Livre
	var input LivreUpdateInput

	if err := h.DB.First(&livre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livre non trouvé"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Titre != nil {
		livre.Titre = *input.Titre
	}
	if input.Auteur != nil {
		livre.Auteur = *input.Auteur
	}
	if input.ISBN != nil {
		livre.ISBN = *input.ISBN
	}
	if input.Disponible != nil {
		livre.Disponible = *input.Disponible
	}
	if input.Quantite != nil {
		livre.Quantite = *input.Quantite
	}

	if err := h.DB.Save(&livre).Error; err != nil {
		if isUniqueViolation(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "Un livre avec cet ISBN existe déjà"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur mise à jour"})
		return
	}

	c.JSON(http.StatusOK, livre)
}

func (h *LivreHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var livre Livre
	if err := h.DB.First(&livre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livre non trouvé"})
		return
	}

	if err := h.DB.Delete(&livre).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur suppression"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livre supprimé"})
}
