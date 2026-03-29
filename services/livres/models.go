package main

import "gorm.io/gorm"

// --- Modèles GORM ---

type Livre struct {
	gorm.Model
	Titre      string `json:"titre"`
	Auteur     string `json:"auteur"`
	ISBN       string `json:"isbn" gorm:"unique"`
	Disponible bool   `json:"disponible" gorm:"default:true"`
	Quantite   int    `json:"quantite" gorm:"default:1"`
}

// --- DTOs ---

type LivreCreateInput struct {
	Titre      string `json:"titre" binding:"required"`
	Auteur     string `json:"auteur" binding:"required"`
	ISBN       string `json:"isbn" binding:"required"`
	Disponible *bool  `json:"disponible"`
	Quantite   *int   `json:"quantite"`
}

type LivreUpdateInput struct {
	Titre      *string `json:"titre"`
	Auteur     *string `json:"auteur"`
	ISBN       *string `json:"isbn"`
	Disponible *bool   `json:"disponible"`
	Quantite   *int    `json:"quantite"`
}
