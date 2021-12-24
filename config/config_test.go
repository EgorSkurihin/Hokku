package config_test

import (
	"testing"

	"github.com/EgorSkurihin/Hokku/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	//Config exists
	conf, err := config.New("config.toml")
	assert.Nil(t, err)
	assert.NotNil(t, conf)

	//Config not exists
	conf, err = config.New("qwe.toml")
	assert.NotNil(t, err)
	assert.Nil(t, conf)
}
