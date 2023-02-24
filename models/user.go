package models

import (
	"crud/databases"
	"log"
)

type User struct {
	//gorm.Model
	ID    int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	NAME  string `json:"name"`
	GRADE int    `json:"grade"`
}

var DB,err = databases.Connection()

func InitDatabase() {
	DB.AutoMigrate(&User{})
	log.Println("Migration Complete")
}

///USER///

func Show() ([]User, error){
	var users []User
	DB.Find(&users)
	return users,nil
}

func ShowOne(id int) (User,error){
	var user = User{ID: id}
	
	r := DB.First(&user)

	if r.Error != nil { //verify if exist
		return user, r.Error
	}

	return user,nil
}

func Add(name string, grade int) (User,error){
	var newUser = User{NAME:name,GRADE:grade}
	DB.Create(& User{NAME:name,GRADE:grade})
	return newUser, nil
}

func Delete(id int) (User,error){
	var delUser = User{ID: id}

	r := DB.First(&delUser) 

	if r.Error != nil { //verify if exist
		return delUser, r.Error
	}

	DB.Delete(&User{ID: id})
	return delUser,nil
}

func Update(id int, name string, grade int) (User, error){
	var upUser = User{ID:id}
	r := DB.First(&upUser)
	if r.Error != nil { //verify if exist
		return upUser, r.Error
	}

	upUser.NAME = name
	upUser.GRADE = grade
	DB.Save(&upUser)
	return upUser,nil
}