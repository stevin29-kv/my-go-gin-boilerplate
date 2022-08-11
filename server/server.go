package server

import (
	cfg "employee-app/config"
	"employee-app/database"
	"employee-app/logger"
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
	// logger.Infof("server runnin at port %s", config.Port)
	// err :=
}
