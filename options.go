package envutils

type Config struct {
	FileName string
}

type Option func(*Config)

func FileName(FileName string) Option {
	return func(c *Config) {
		c.FileName = FileName
	}
}
