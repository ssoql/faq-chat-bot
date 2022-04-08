package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "APP_ENV", appEnv)
	assert.EqualValues(t, "8083", port)
	assert.EqualValues(t, "prod", production)
}

func TestIsProduction(t *testing.T) {
	assert.EqualValues(t, false, IsProduction())
}

func TestGetPort(t *testing.T) {
	assert.EqualValues(t, ":8083", GetPort())
}
