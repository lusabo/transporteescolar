package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

func ListarEstados(db *gorm.DB) ([]models.Estado, error) {
	var estados []models.Estado
	if err := db.Find(&estados).Error; err != nil {
		return nil, err
	}
	return estados, nil
}
