package configs

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	MongodbHost     string
	MongodbPort     string
	MongodbUser     string
	MongodbPassword string
	MongodbName     string
	MongodbAuth     string
	WebServerPort   int
}

func New() (*Config, error) {
	mongodbHost := os.Getenv("MongodbHost")
	if len(mongodbHost) == 0 {
		return nil, errors.New("please specify MongodbHost variable in env")
	}

	mongodbPort := os.Getenv("MongodbPort")
	if len(mongodbPort) == 0 {
		return nil, errors.New("please specify MongodbPort variable in env")
	}

	mongodbName := os.Getenv("MongodbName")

	if len(mongodbName) == 0 {
		return nil, errors.New("please specify MongodbName variable in env")
	}

	webServerPortString := os.Getenv("WebServerPort")
	if len(webServerPortString) == 0 {
		return nil, errors.New("please specify WebServerPort variable in env")
	}
	webServerPort, err := strconv.Atoi(webServerPortString)
	if err != nil {
		return nil, errors.New("WebServerPort must be integer")
	}

	return &Config{
		MongodbHost:     mongodbHost,
		MongodbPort:     mongodbPort,
		MongodbUser:     os.Getenv("MongodbUser"),
		MongodbPassword: os.Getenv("MongodbPassword"),
		MongodbName:     mongodbName,
		MongodbAuth:     os.Getenv("MongodbAuth"),
		WebServerPort:   webServerPort,
	}, nil
}
