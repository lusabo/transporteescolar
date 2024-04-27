package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

func ListarCidadesPorEstado(db *gorm.DB, estadoId string) ([]models.Cidade, error) {

	var cidades []models.Cidade
	if err := db.Where("estado_id = ?", estadoId).Find(&cidades).Error; err != nil {
		return nil, err
	}

	return cidades, nil
}
