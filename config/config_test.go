package config

import (
	"testing"

	"github.com/test-go/testify/assert"
)

func TestLoad(t *testing.T) {
	// create a temporary file for testing
	testFile := "./config_test.yml"

	// load the config from the temporary file
	cfg, err := Load(testFile)
	if err != nil {
		t.Fatalf("Failed to load config from temporary file: %v", err)
	}

	// check that the config values are correct
	if cfg.Serial != "test-serial" {
		t.Errorf("Expected serial=test-serial, got %s=%s", "serial", cfg.Serial)
	}
	if cfg.Chat.Name != "cai.zhihong" {
		t.Errorf("Expected name=cai.zhihong, got %s=%s", "name", cfg.Chat.Name)
	}
	if cfg.Chat.Age != 44 {
		t.Errorf("Expected age=44, got %s=%d", "age", cfg.Chat.Age)
	}
}

func TestLoad_FileNotExists(t *testing.T) {
	testFile := "./not_exists.yml"

	// load the config from the temporary file
	cfg, err := Load(testFile)
	assert.NotNil(t, err)
	assert.Nil(t, cfg)
}
