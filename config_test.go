package envutils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type EnvConfig struct {
	TestString  string  `env:"TEST.STRING"`
	TestInt     int     `env:"TEST.INT"`
	TestInt8    int8    `env:"TEST.INT8"`
	TestInt32   int32   `env:"TEST.INT32"`
	TestInt64   int64   `env:"TEST.INT64"`
	TestBoolean bool    `env:"TEST.BOOLEAN"`
	TestFloat32 float32 `env:"TEST.FLOAT32"`
	TestFloat64 float64 `env:"TEST.FLOAT64"`
}

func TestParseConfig(t *testing.T) {
	var config EnvConfig
	fileName := ""
	if os.Getenv("APPLICATION_ENV") == "" {
		fileName = "local.env"
	}
	err := ParseConfig(&config, FileName(fileName))
	assert.Nil(t, err)
	assert.EqualValues(t, "this is string", config.TestString)
	assert.EqualValues(t, 1, config.TestInt)
	assert.EqualValues(t, 8, config.TestInt8)
	assert.EqualValues(t, 32, config.TestInt32)
	assert.EqualValues(t, 64, config.TestInt64)
	assert.EqualValues(t, true, config.TestBoolean)
	assert.EqualValues(t, 0.32, config.TestFloat32)
	assert.EqualValues(t, 64.64, config.TestFloat64)
}
