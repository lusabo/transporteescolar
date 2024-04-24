package models

import (
	"gorm.io/gorm"
)

type Bairro struct {
	gorm.Model
	Nome      string      `gorm:"not null;size:50"`
	CidadeID  uint        `gorm:"not null"`
	Motorista []Motorista `gorm:"many2many:motorista_bairros;"`
}
