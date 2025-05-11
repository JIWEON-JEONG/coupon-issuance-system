package configuration

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"DATABASE_HOST"`
	DBUserName string `mapstructure:"DATABASE_USER"`
	DBPassword string `mapstructure:"DATABASE_PASSWORD"`
	DBName     string `mapstructure:"DATABASE_DATABASE"`
	DBPort     string `mapstructure:"DATABASE_Port"`

	RedisUrl string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
