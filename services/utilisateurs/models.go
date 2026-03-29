package main

import "gorm.io/gorm"

// --- Modèles GORM ---

type Utilisateur struct {
	gorm.Model
	Nom   string `json:"nom"`
	Email string `json:"email" gorm:"unique"`
	Type  string `json:"type"`
}

// --- DTOs ---

type UtilisateurCreateInput struct {
	Nom   string `json:"nom" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Type  string `json:"type" binding:"required"`
}

type UtilisateurUpdateInput struct {
	Nom   *string `json:"nom"`
	Email *string `json:"email" binding:"omitempty,email"`
	Type  *string `json:"type"`
}
