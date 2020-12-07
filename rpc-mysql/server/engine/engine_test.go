package engine

import (
	"github.com/stretchr/testify/assert"
	"os"
	"rpc-mysql/pkg/config"
	"testing"
)

func TestEngine_Run(t *testing.T) {
	filePath := os.Getenv("test_config_path")
	if filePath == "" {
		panic("path not found")
	}

	err := config.LoadYAMLConfig(filePath)
	assert.NoError(t, err)

	cfg := config.GetConfig()

	_, err = NewEngine(cfg)
	assert.NoError(t, err)
}
