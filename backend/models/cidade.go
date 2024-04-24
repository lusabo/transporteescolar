package models

import (
	"gorm.io/gorm"
)

type Cidade struct {
	gorm.Model
	Nome     string `gorm:"type:varchar(30);not null"`
	EstadoID uint
	Estado   Estado
	Bairros  []Bairro
}
