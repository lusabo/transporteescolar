package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

func ListarEscolasPorCidade(db *gorm.DB, cidadeId uint) ([]models.Escola, error) {
	var escolas []models.Escola

	if err := db.Where("bairro_id IN (SELECT id FROM bairros WHERE cidade_id = ?)", cidadeId).Find(&escolas).Error; err != nil {
		return nil, err
	}

	return escolas, nil
}
