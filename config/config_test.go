package config_test

import (
	"go-scrum-poker-bot/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	c := config.NewConfig()

	assert.Equal(t, "0.0.0.0", c.Redis.Host)
	assert.Equal(t, 6379, c.Redis.Port)
	assert.Equal(t, 0, c.Redis.DB)
	assert.Equal(t, []string{"?", "0", "0.5", "1", "2", "3", "5", "8", "13", "20", "40", "100"}, c.App.PokerRanks)
}

func TestNewConfigIncorrectIntFromEnv(t *testing.T) {
	os.Setenv("REDIS_PORT", "-")

	assert.Panics(t, func() { config.NewConfig() })
}
