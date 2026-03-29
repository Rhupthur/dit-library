package main

import "gorm.io/gorm"

type LivreHandler struct {
	DB *gorm.DB
}

func NewLivreHandler(db *gorm.DB) *LivreHandler {
	return &LivreHandler{DB: db}
}
