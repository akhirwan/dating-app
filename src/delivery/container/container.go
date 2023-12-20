package container

import (
	"dating-app/src/config"
	"dating-app/src/infrastructure/database"
	"log"
)

type Container struct {
	EnvironmentConfig config.EnvironmentConfig
}

func SetupContainer() Container {
	log.Println("Starting new container...")

	log.Println("Loading config...")
	config, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Loading database...")
	_, err = database.NewPostgreSQLDBConnection(&config.Database)
	if err != nil {
		log.Panic(err)
	}

	return Container{
		EnvironmentConfig: config,
	}
}
