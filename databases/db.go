package databases

import (
	//"crud/models"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

func Connection() (*gorm.DB,error){
	user := "root"
	pass := ""
	bd := "registro"
	dns := user + ":" + pass + "@tcp(127.0.0.1)/" + bd + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dns),&gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to BD")
	}

	log.Println("Database connected...")
	return DB,nil

}



