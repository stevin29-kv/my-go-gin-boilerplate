package server

import (
	cfg "employee-app/config"
	"employee-app/database"
	"employee-app/internal/controller"
	"employee-app/internal/repository"
	"employee-app/internal/service"
	"employee-app/logger"
	"employee-app/server/router"
)

func Start() {
	config := cfg.GetConfig()

	db, dbErr := database.PrepareDatabase()
	if dbErr != nil {
		panic(dbErr)
	}

	_, err := logger.InitLogger(config.Env)
	if err != nil {
		logger.Errorf("error initializing logger", err)
	}
	logger.Info(db)

	repo := repository.InitRepository(db)
	serv := service.InitService(repo)
	cont := controller.InitController(serv)

	router := router.PrepareRouter(&router.CapsuleRouter{
		DB:         db,
		Repository: repo,
		Service:    serv,
		Controller: cont,
	})

	logger.Infof("server running at port %s", config.Port)
	err = router.Run(":" + config.Port)
	if err != nil {
		logger.Fatalf("error running server - %s", err.Error())
	}
}
