package config

import "time"

type Loader interface {
	LoadConfig(path string, config Config) error
}

type Config interface {
	Name() string
	Defaults() map[string]string
}

type AppConfig struct {
	ApplicationAPITimeout  time.Duration `mapstructure:"APPLICATION_API_TIMEOUT"`
	ApplicationName        string        `mapstructure:"APPLICATION_NAME"`
	ApplicationEnvironment string        `mapstructure:"APPLICATION_ENV"`
	ApplicationVersion     string        `mapstructure:"APPLICATION_VERSION"`
	LogLevel               string        `mapstructure:"LOG_LEVEL"`
	CORSAllowedOrigins     string        `mapstructure:"CORS_ALLOWED_ORIGINS"`
	DatabaseURL            string        `mapstructure:"DATABASE_URL"`
	DatabaseUsername       string        `mapstructure:"DATABASE_USERNAME"`
	DatabasePassword       string        `mapstructure:"DATABASE_PASSWORD"`
	DatabaseDB             string        `mapstructure:"DATABASE_DB"`
	DatabasePrintQueries   bool          `mapstructure:"DATABASE_PRINT_QUERIES"`
}

func (ac *AppConfig) Name() string {
	return "app"
}

func (ac *AppConfig) Defaults() map[string]string {
	return map[string]string{
		"APPLICATION_API_TIMEOUT": "30s",
		"APPLICATION_VERSION":     "0.0.1",
		"APPLICATION_NAME":        "it-test",
		"APPLICATION_ENV":         "local",
		"LOG_LEVEL":               "debug",
		"CORS_ALLOWED_ORIGINS":    "*",
		"DATABASE_URL":            "localhost:5432",
		"DATABASE_USERNAME":       "it",
		"DATABASE_PASSWORD":       "it",
		"DATABASE_DB":             "it",
		"DATABASE_PRINT_QUERIES":  "true",
	}
}
