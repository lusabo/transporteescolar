package models

import (
	"gorm.io/gorm"
)

type Motorista struct {
	gorm.Model
	Nome    string   `gorm:"not null"`
	Email   string   `gorm:"not null"`
	Senha   string   `gorm:"not null"`
	Escolas []Escola `gorm:"many2many:motorista_escolas;"`
	Bairros []Bairro `gorm:"many2many:motorista_bairros;"`
}
