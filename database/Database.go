/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 07:37:04
*/
package Database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/joho/godotenv"
	"os"
)

var db *sql.DB
var err error

func Init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error open .env file")
	}

	db_host 	:= os.Getenv("DB_HOST")
	db_driver 	:= os.Getenv("DB_DRIVER")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_name 	:= os.Getenv("DB_DATABASE")

	db, err = sql.Open(db_driver, db_username +":" + db_password + "@tcp(" + db_host + ")/" + db_name)

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
	}
}
