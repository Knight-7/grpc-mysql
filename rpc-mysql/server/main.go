package main

import (
	"fmt"
	"os"
	"rpc-mysql/pkg/config"
	"rpc-mysql/rpc-mysql/server/engine"
)

func main() {
	filePath := os.Getenv("grpc_mysql_config")
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
