# env-utils

load evn config to struct

```
var config EnvConfig
fileName := ""
if os.Getenv("APPLICATION_ENV") == "" {
    fileName = "local.env"
}
err := ParseConfig(&config, FileName(fileName))
```