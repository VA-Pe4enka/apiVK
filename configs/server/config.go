package server

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	VKToken string `mapstructure:"VkOpenApiToken"`
}

type Service struct {
	Service *VkApiConfig
}

type VkApiConfig interface {
	LoadApiConfig(filename string) (config Config, err error)
}

func (service *Service) LoadApiConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("err in config.go: ", err.Error())
	}

	err = viper.Unmarshal(&config)
	return
}
