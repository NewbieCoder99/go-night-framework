/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 07:35:30
*/
package server

import(
	"os"
	"log"
	"github.com/NewbieCoder99/go-night-framework/config"
	"github.com/joho/godotenv"
)

func Init() {

	config.GetConfig()
	router := NewRouter()

	err := godotenv.Load()
		if err != nil { log.Fatal("Error loading .env file") }

	host_name 	:= os.Getenv("HOST_NAME")
	host_port 	:= os.Getenv("HOST_PORT")

	router.Run(host_name + ":" + host_port)
}
