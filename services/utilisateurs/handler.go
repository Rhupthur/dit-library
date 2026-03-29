package main

import "gorm.io/gorm"

type UtilisateurHandler struct {
	DB *gorm.DB
}

func NewUtilisateurHandler(db *gorm.DB) *UtilisateurHandler {
	return &UtilisateurHandler{DB: db}
}
