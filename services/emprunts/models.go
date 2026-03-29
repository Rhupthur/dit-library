package main

import (
	"time"

	"gorm.io/gorm"
)

// --- Modèles GORM ---

type Emprunt struct {
	gorm.Model
	LivreID       uint       `json:"livre_id"`
	UtilisateurID uint       `json:"utilisateur_id"`
	DateEmprunt   time.Time  `json:"date_emprunt"`
	DateRetour    *time.Time `json:"date_retour"`
	EnRetard      bool       `json:"en_retard" gorm:"default:false"`
}

// --- DTOs ---

type EmpruntCreateInput struct {
	LivreID       uint `json:"livre_id" binding:"required,gt=0"`
	UtilisateurID uint `json:"utilisateur_id" binding:"required,gt=0"`
}
