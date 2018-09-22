/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-23 01:44:05
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
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func DBConf(driver string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error open .env file")
	}

	db_driver 	:= os.Getenv("DB_DRIVER")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_name 	:= os.Getenv("DB_DATABASE")

	if driver == "setup" {
		if db_driver == "mysql" {
			return db_username +":"+ db_password +"@/" + db_name + "?charset=utf8&parseTime=True&loc=Local"
		}
	}

	return db_driver
}

func BaseUrl() (string) {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	base_url 	:= os.Getenv("BASE_URL")
	port 		:= os.Getenv("HOST_PORT")

	return base_url+":"+port+"/"
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

