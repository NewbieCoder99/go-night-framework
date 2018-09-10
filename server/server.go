package server

import(
	"github.com/NewbieCoder99/go-night-framework/config"
	"github.com/joho/godotenv"
	"os"
	"log"
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
