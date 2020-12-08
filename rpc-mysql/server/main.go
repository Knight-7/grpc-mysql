package main

import (
	"flag"
	"fmt"
	"log"
	"rpc-mysql/pkg/config"
	"rpc-mysql/rpc-mysql/server/engine"
)

func main() {
	var (
		filePath string
		logBool  bool
		authBool bool
	)

	flag.StringVar(&filePath, "config", "config.yaml", "config name")
	flag.BoolVar(&logBool, "log", false, "whether use log")
	flag.BoolVar(&authBool, "auth", false, "whether user auth")

	flag.Parse()

	if filePath == "" {
		log.Panic("config.yaml not found")
	}

	err := config.LoadYAMLConfig(filePath)
	if err != nil {
		//TODO: add log
		fmt.Println(err)
		return
	}
	cfg := config.GetConfig()

	daoEngine, err := engine.NewEngine(cfg)
	if err != nil {
		//TODO: add log
		fmt.Println(err)
		return
	}
	daoEngine.Run()
}
