package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

var typesValides = map[string]struct{}{
	"etudiant":   {},
	"professeur": {},
	"admin":      {},
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

func (h *UtilisateurHandler) GetAll(c *gin.Context) {
	var utilisateurs []Utilisateur

	if err := h.DB.Find(&utilisateurs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur récupération"})
		return
	}

	c.JSON(http.StatusOK, utilisateurs)
}

func (h *UtilisateurHandler) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var utilisateur Utilisateur
	if err := h.DB.First(&utilisateur, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	c.JSON(http.StatusOK, utilisateur)
}

func (h *UtilisateurHandler) Create(c *gin.Context) {
	var input UtilisateurCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := typesValides[input.Type]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type invalide : etudiant, professeur ou admin"})
		return
	}

	utilisateur := Utilisateur{
		Nom:   input.Nom,
		Email: input.Email,
		Type:  input.Type,
	}

	if err := h.DB.Create(&utilisateur).Error; err != nil {
		if isUniqueViolation(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "Un utilisateur avec cet email existe déjà"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur création"})
		return
	}

	c.JSON(http.StatusCreated, utilisateur)
}

func (h *UtilisateurHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var utilisateur Utilisateur
	var input UtilisateurUpdateInput

	if err := h.DB.First(&utilisateur, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Nom != nil {
		utilisateur.Nom = *input.Nom
	}
	if input.Email != nil {
		utilisateur.Email = *input.Email
	}
	if input.Type != nil {
		if _, ok := typesValides[*input.Type]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Type invalide : etudiant, professeur ou admin"})
			return
		}
		utilisateur.Type = *input.Type
	}

	if err := h.DB.Save(&utilisateur).Error; err != nil {
		if isUniqueViolation(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "Un utilisateur avec cet email existe déjà"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur mise à jour"})
		return
	}

	c.JSON(http.StatusOK, utilisateur)
}

func (h *UtilisateurHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	var utilisateur Utilisateur
	if err := h.DB.First(&utilisateur, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	if err := h.DB.Delete(&utilisateur).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur suppression"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Utilisateur supprimé"})
}
