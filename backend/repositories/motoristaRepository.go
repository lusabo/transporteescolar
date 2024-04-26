package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

func ListarMotoristas(db *gorm.DB, bairroId uint, escolaId uint) ([]models.Motorista, error) {

	query := db.Model(&models.Motorista{})

	if bairroId != 0 {
		query = query.Joins("JOIN motorista_bairros ON motorista_bairros.motorista_id = motorista.id").
			Where("motorista_bairros.bairro_id = ?", bairroId)
	}

	if escolaId != 0 {
		query = query.Joins("JOIN motorista_escolas ON motorista_escolas.motorista_id = motorista.id").
			Where("motorista_escolas.escola_id = ?", escolaId)
	}

	var motoristas []models.Motorista
	if err := query.Find(&motoristas).Error; err != nil {
		return nil, err
	}

	return motoristas, nil
}

func BuscarMotorista(db *gorm.DB, motoristaId uint) (models.Motorista, error) {
	var motorista models.Motorista
	if err := db.First(&motorista, motoristaId).Error; err != nil {
		return models.Motorista{}, err
	}
	return motorista, nil
}
