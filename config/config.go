package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	WSPort       string `mapstructure:"WS_PORT"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
}

func LoadConfig(path string) *conf {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg *conf
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return cfg
}

func (c *conf) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
