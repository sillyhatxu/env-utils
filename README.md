# env-utils

load evn config to struct

```

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

var config EnvConfig
fileName := ""
if os.Getenv("APPLICATION_ENV") == "" {
    fileName = "local.env"
}
err := ParseConfig(&config, FileName(fileName))
```