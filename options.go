package envutils

type Config struct {
	Filenames []string
}

type Option func(*Config)

func Filenames(Filenames []string) Option {
	return func(c *Config) {
		c.Filenames = Filenames
	}
}
