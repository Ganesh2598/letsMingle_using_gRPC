package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	//"os"
	//"github.com/joho/godotenv"
	"../model"
	"log"
)

var Db *gorm.DB = nil
var err error

func Connection()  {

	/*err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}*/
	Db, err = gorm.Open("mysql", "root:Ganesh@2598@/sample?parseTime=true")
	//Db, err = gorm.Open(os.Getenv("DB"),os.Getenv("DBUSER")+":"+os.Getenv("DBPASSWORD")+"@/"+os.Getenv("DBNAME")+"?parseTime=true")
	if err != nil {
		log.Fatal("Not Connected",err)
	}
	log.Printf("Connected")
	//Db.DropTableIfExists(&model.Post{}, &model.Comment{})
	Db.AutoMigrate(&model.User{})
	Db.AutoMigrate(&model.Post{})
	Db.AutoMigrate(&model.Comment{})
	Db.Model(&model.Comment{}).AddForeignKey("post_id", "posts(id)","cascade", "cascade")
	Db.AutoMigrate(&model.Friend{})
	Db.LogMode(true)

}