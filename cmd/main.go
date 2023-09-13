package main

import (
	"fmt"
	"github.com/livelists/livelist-server/cmd/client"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/config/boot"
	"github.com/livelists/livelist-server/pkg/logger"
	"github.com/livelists/livelist-server/pkg/websocket"
	"io/ioutil"
	"os"
)

func main() {
	conf, err := getConfig()
	if err != nil {
		logger.Errorw("Read config error", err)
	}

	StartServer(conf)
}

func StartServer(conf *config.Config) {
	mongoClient, err := config.ConnectToMongo(conf.Mongo)

	if err != nil {
		logger.Errorw("Mongo connection error", err)
	}
	boot.SeedMongo(mongoClient, conf)

	go client.StartTwirpRPC()

	websocket.StartWS(conf.Port)
}

func getConfigString(configFile string) (string, error) {
	outConfigBody, err := ioutil.ReadFile(configFile)
	if err != nil {
		return "", err
	}

	return string(outConfigBody), nil
}

func getConfig() (*config.Config, error) {
	path, err := os.Getwd()
	fmt.Print("path::", path, "endPath")
	confString, err := getConfigString(path + "config.yaml")
	if err != nil {
		return nil, err
	}

	conf, err := config.NewConfig(confString)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
