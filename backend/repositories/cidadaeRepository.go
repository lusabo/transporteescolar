package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

func ListarCidadesPorEstado(db *gorm.DB, uf string) ([]models.Cidade, error) {

	var estado models.Estado
	if err := db.Where("uf = ?", uf).First(&estado).Error; err != nil {
		return nil, err
	}

	var cidades []models.Cidade
	if err := db.Where("estado_id = ?", estado.ID).Find(&cidades).Error; err != nil {
		return nil, err
	}

	return cidades, nil
}
