package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

func GetAllCidades(db *gorm.DB) ([]models.Cidade, error) {
	var cidades []models.Cidade
	if err := db.Find(&cidades).Error; err != nil {
		return nil, err
	}
	return cidades, nil
}
