package main

import (
	"Sunat/internal/configs"
	"Sunat/internal/handlers"
	"Sunat/internal/repositories"
	"Sunat/internal/services"
	"Sunat/loggers"
	"Sunat/pkg/database"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	log.Println("Start connection")
	log.Println("Step 1")
	logger, err := loggers.InitLogger()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(logger2 *zap.Logger) {
		err = logger2.Sync()
		if err != nil {
			log.Println(err)
			return
		}
	}(logger)

	log.Println("Step 2")
	config, err := configs.InitConfigs()
	if err != nil {
		logger.Error("Error in configs")
		return
	}

	log.Println("Step 3")
	db, err := database.InitDatabase(config)
	if err != nil {
		logger.Error("Error in database")
		return
	}

	log.Println("Step 4")
	repository := repositories.NewRepository(db)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service, logger)
	mux := handlers.InitRouter(handler)

	srv := http.Server{
		Addr:    config.Server.Host + config.Server.Port,
		Handler: mux,
	}

	log.Println("start")
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("Issue in ListeningAndServe!", zap.Error(err))
		return
	}
}
