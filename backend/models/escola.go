package models

import (
	"gorm.io/gorm"
)

type Escola struct {
	gorm.Model
	Nome      string      `gorm:"not null"`
	Endereco  string      `gorm:"not null"`
	BairroID  uint        `gorm:"not null"`
	Motorista []Motorista `gorm:"many2many:motorista_escolas;"`
}
