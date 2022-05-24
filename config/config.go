package config

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/iam1912/gemseries/gemleaf/utils"
	"github.com/spf13/viper"
)

type AppConfig struct {
	DB           string
	Port         string
	ReadTimeOut  int
	WriteTimeOut int
}

var (
	once       sync.Once
	_AppConfig *AppConfig
)

func MustGetAppConfig() AppConfig {
	once.Do(func() {
		path := utils.InferRootDir("conf/config.yml")
		fmt.Println(path)
		viper.SetConfigFile(filepath.Join(path, "conf/config.yml"))
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Fail to parse 'conf/config.yml': %v", err)
		}
		viper.Unmarshal(&_AppConfig)
	})
	return *_AppConfig
}
