package models

import (
	"gorm.io/gorm"
)

type Escola struct {
	gorm.Model
	Nome       string `gorm:"not null"`
	Endereco   string `gorm:"not null"`
	BairroID   uint
	Bairro     Bairro
	Motoristas []Motorista `gorm:"many2many:escola_motoristas;"`
}
