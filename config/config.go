/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 09:45:07
*/
package config

import (
	"os"
	"log"
	"path/filepath"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var config *viper.Viper

func Init() {
	var err error
	config = viper.New()
	// config.SetConfigType("yaml")
	config.SetConfigName("test")
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func GetMode() {
	err := godotenv.Load()
		if err != nil { log.Fatal("Error loading .env file") }
	 // Set from .env to "release" or "debug" 
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}

