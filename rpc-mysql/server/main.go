package main

import (
	"flag"
	"fmt"
	"rpc-mysql/pkg/config"
	"rpc-mysql/rpc-mysql/server/engine"
)

func main() {
	var filePath string

	flag.StringVar(&filePath, "config", "config.yaml", "config name")
	flag.Parse()

	if filePath == "" {
		panic("config.yaml not found")
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
