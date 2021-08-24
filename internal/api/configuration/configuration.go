package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	DBHost string `mapstructure:"DB_HOST"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBPort string `mapstructure:"DB_PORT"`
	DBName string `mapstructure:"DB_NAME"`
}

func (c *Configuration) GetDBUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}

func LoadConfig(path string) (config Configuration, err error) {
	// Set default values
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASS", "password")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_NAME", "todos")

	viper.AddConfigPath(path)

	var filename = os.Getenv("CONFIG_FILE_NAME")
	if filename == "" {
		filename = "app"
	}

	viper.SetConfigName(filename)

	var filetype = os.Getenv("CONFIG_FILE_TYPE")
	if filetype == "" {
		filetype = "env"
	}

	viper.SetConfigType(filetype)

	viper.AutomaticEnv()

	err1 := viper.ReadInConfig()

	err2 := viper.Unmarshal(&config)

	if err2 != nil {
		err = err2
		return
	}

	err = err1
	return
}
