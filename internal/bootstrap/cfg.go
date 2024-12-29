package bootstrap

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfig struct {
	LogPath      string
	DbConnStr    string `mapstructure:"DB_CONN_STR"`
	RedisConnStr string `mapstructure:"REDIS_CONN_STR"`
	RedisPwd     string `mapstructure:"REDIS_PWD"`
	JwtSecret    string `mapstructure:"JWT_SECRET"`
	AppName      string `mapstructure:"APP_NAME"`
}

const (
	LOCAL_ENV = "local"
	DEV_ENV   = "dev"
	PROD_ENV  = "prod"
)

func ReadEnv() string {
	return os.Getenv("ENV")
}

func ReadAppConfig(env string) (*AppConfig, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName(env)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if env == LOCAL_ENV {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	viper.AutomaticEnv()
	if err := viper.MergeInConfig(); err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
