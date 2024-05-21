package loggers

import (
	"go.uber.org/zap"
	"log"
)

func InitLogger() (*zap.Logger, error) {
	logConfig := zap.NewProductionConfig()
	logConfig.OutputPaths = []string{"./loggers/logs.log"}

	logConfig.Level.SetLevel(zap.InfoLevel)

	logg, err := logConfig.Build()
	if err != nil {
		log.Println("Issue in InitLogger!")
		return nil, err
	}

	return logg, nil
}
