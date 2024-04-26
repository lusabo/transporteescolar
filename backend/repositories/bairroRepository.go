package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

func ListarBairrosPorCidade(db *gorm.DB, cidadeId uint) ([]models.Bairro, error) {
	var bairros []models.Bairro
	if err := db.Where("cidade_id = ?", cidadeId).Find(&bairros).Error; err != nil {
		return nil, err
	}

	return bairros, nil
}
