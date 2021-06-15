package gqldb

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name      string
	Age       int32
	Contact Contact
	Pets      []Pet
}

type Contact struct {
	gorm.Model
	Email string
	Phone string
	UserID int
}

type Pet struct {
	gorm.Model
	Name    string
	Species string
	Age     int32
	User    User
	UserID int
}

type BasicDb struct {
	db *gorm.DB
}

func ConnectDB() BasicDb {

	db, err := gorm.Open(mysql.Open(
		"leighton:123456@tcp(127.0.0.1:3307)/graphqlproject?parseTime=true"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Contact{}, &Pet{})
	
	return BasicDb{db: db}
}

// Creates user and contact in the database
func (ctx BasicDb) Seed() {
	db := ctx.db

	users := []User{
		{
			Name:    "Leighton",
			Age:     60,
			Contact: Contact{Email: "leighton@email.com", Phone: "0279446788"},
			Pets: []Pet{
				{Name: "Jack", Species: "Dog", Age: 7},
				{Name: "Roxy", Species: "Dog", Age: 17},
			},
		},
		{
			Name:    "Tim",
			Age:     22,
			Contact: Contact{Email: "tim@email.com", Phone: "027954178"},
			Pets: []Pet{
				{Name: "Pip", Species: "Cat", Age: 1},
				{Name: "Zoomer", Species: "Mouse", Age: 13},
			},
		},
		{
			Name:    "Eric",
			Age:     60,
			Contact: Contact{Email: "eric@email.com", Phone: "0278369755"},
			Pets: []Pet{
				{Name: "Bob", Species: "Dog", Age: 3},
				{Name: "Hutch", Species: "Dog", Age: 12},
			},
		},
	}

	db.Create(&users)
}
