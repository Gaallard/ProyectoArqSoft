package db

import (
	"backend/clients/courses"
	"backend/clients/inscriptions"
	"backend/clients/users"
	"backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "arqsoft1"  //Nombre de la base de datos local de ustedes
	DBUser := "root"      //usuario de la base de datos, habitualmente root
	DBPass := "root"      //password del root en la instalacion
	DBHost := "127.0.0.1" //host de la base de datos. hbitualmente 127.0.0.1
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3307)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	users.Db = db
	courses.Db = db
	inscriptions.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Course{})
	db.AutoMigrate(&models.Inscription{})

	log.Info("Finishing Migration Database Tables")
}
