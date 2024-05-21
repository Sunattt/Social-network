package configs

import (
	"Sunat/pkg/models"
	"encoding/json"
	"os"
)

var Setting *models.ConfigsModels

func InitConfigs() (*models.ConfigsModels, error) {
	file, err := os.OpenFile("./internal/configs/configs.json", os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	var configs *models.ConfigsModels
	err = json.NewDecoder(file).Decode(&configs)
	if err != nil {
		return nil, err
	}

	return configs, nil
}
