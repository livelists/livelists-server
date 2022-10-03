package main

import (
	"fmt"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/websocket"
	"io/ioutil"
)

func getConfigString(configFile string) (string, error) {
	outConfigBody, err := ioutil.ReadFile(configFile)
	if err != nil {
		return "", err
	}

	return string(outConfigBody), nil
}

func getConfig() (*config.Config, error) {
	confString, err := getConfigString("C:\\Users\\nikra\\Desktop\\livelists\\livelists-server\\pkg\\config\\config.yaml")
	if err != nil {
		return nil, err
	}

	conf, err := config.NewConfig(confString)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func main() {
	conf, err := getConfig()
	websocket.StartWS(conf.Port)

	if err == nil {
		fmt.Println(conf)
	}
}
