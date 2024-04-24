package models

import (
	"gorm.io/gorm"
)

type Estado struct {
	gorm.Model
	Uf      string `gorm:"type:varchar(2);unique;not null"`
	Nome    string `gorm:"type:varchar(20);not null"`
	Cidades []Cidade
}
