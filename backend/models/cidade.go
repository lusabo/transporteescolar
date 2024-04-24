package models

import (
	"gorm.io/gorm"
)

type Cidade struct {
	gorm.Model
	Nome     string `gorm:"not null;size:30"`
	EstadoID uint   `gorm:"not null"`
	Bairros  []Bairro
}
