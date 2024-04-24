package models

import (
	"gorm.io/gorm"
)

type Bairro struct {
	gorm.Model
	Nome       string `gorm:"type:varchar(50);not null"`
	CidadeID   uint
	Cidade     Cidade
	Escolas    []Escola
	Motoristas []Motorista `gorm:"many2many:motorista_bairros;"`
}
