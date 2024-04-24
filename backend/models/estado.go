package models

import (
	"gorm.io/gorm"
)

type Estado struct {
	gorm.Model
	Uf      string `gorm:"not null;unique;size:2"`
	Nome    string `gorm:"not null;size:20"`
	Cidades []Cidade
}
