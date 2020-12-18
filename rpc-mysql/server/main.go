package main

import (
	"flag"
	"log"
	"rpc-mysql/pkg/config"
	"rpc-mysql/rpc-mysql/server/engine"
)

func main() {
	var filePath string

	flag.StringVar(&filePath, "config", "config.yaml", "config name")

	flag.Parse()

	if filePath == "" {
		log.Panic("config.yaml not found")
	}

	err := config.LoadYAMLConfig(filePath)
	if err != nil {
		log.Fatalln("load yaml config failed")
		return
	}
	cfg := config.GetConfig()

	daoEngine, err := engine.NewEngine(cfg)
	if err != nil {
		log.Fatalln("engine start failed")
		return
	}
	daoEngine.Run()
}
