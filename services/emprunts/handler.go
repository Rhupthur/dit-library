package main

import "gorm.io/gorm"

type EmpruntHandler struct {
	DB *gorm.DB
}

func NewEmpruntHandler(db *gorm.DB) *EmpruntHandler {
	return &EmpruntHandler{DB: db}
}
