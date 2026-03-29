package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *EmpruntHandler) GetAll(c *gin.Context) {
	var emprunts []Emprunt

	if err := h.DB.Find(&emprunts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur récupération"})
		return
	}

	c.JSON(http.StatusOK, emprunts)
}

func (h *EmpruntHandler) Emprunter(c *gin.Context) {
	var input EmpruntCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int64
	h.DB.Model(&Emprunt{}).
		Where("livre_id = ? AND date_retour IS NULL", input.LivreID).
		Count(&count)

	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Livre déjà emprunté"})
		return
	}

	emprunt := Emprunt{
		LivreID:       input.LivreID,
		UtilisateurID: input.UtilisateurID,
		DateEmprunt:   time.Now(),
		EnRetard:      false,
	}

	if err := h.DB.Create(&emprunt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur création emprunt"})
		return
	}

	c.JSON(http.StatusCreated, emprunt)
}

func (h *EmpruntHandler) Retourner(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var emprunt Emprunt
	if err := h.DB.First(&emprunt, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Emprunt non trouvé"})
		return
	}

	if emprunt.DateRetour != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Livre déjà retourné"})
		return
	}

	now := time.Now()
	emprunt.DateRetour = &now
	dateLimite := emprunt.DateEmprunt.Add(14 * 24 * time.Hour)
	emprunt.EnRetard = now.After(dateLimite)

	if err := h.DB.Save(&emprunt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur retour"})
		return
	}

	c.JSON(http.StatusOK, emprunt)
}

func (h *EmpruntHandler) Historique(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var emprunts []Emprunt
	if err := h.DB.Where("utilisateur_id = ?", id).Find(&emprunts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur récupération historique"})
		return
	}

	c.JSON(http.StatusOK, emprunts)
}
