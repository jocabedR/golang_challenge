package db

import (
	"fmt"
	"golang_challenge/pkg/handles"
)

type User struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT:false"`
	Username string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Address  string `gorm:"not null"`
	GeoLat   string `gorm:"not null"`
	GeoLng   string `gorm:"not null"`
	Phone    string `gorm:"not null"`
	Website  string `gorm:"not null"`
	Company  string `gorm:"not null"`
}

type TODO struct {
	ID        int  `gorm:"primary_key;AUTO_INCREMENT:false"`
	User      User `gorm:"foreignkey:UserID"` // use UserID as foreign key
	UserID    int
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"not null"`
}

type Album struct {
	ID     int  `gorm:"primary_key;AUTO_INCREMENT:false"`
	User   User `gorm:"foreignkey:UserID"` // use UserID as foreign key
	UserID int
	Title  string `gorm:"not null"`
}

type Post struct {
	ID     int  `gorm:"primary_key;AUTO_INCREMENT:false"`
	User   User `gorm:"foreignkey:UserID"` // use UserID as foreign key
	UserID int
	Title  string `gorm:"not null"`
	Body   string `gorm:"not null"`
}

type Comment struct {
	ID     int  `gorm:"primary_key;AUTO_INCREMENT:false"`
	Post   Post `gorm:"foreignkey:PostID"` // use UserID as foreign key
	PostID int
	Name   string `gorm:"not null"`
	Email  string `gorm:"not null"`
	Body   string `gorm:"not null"`
}

func Migration() {
	db := DB
	// This table creation its going to have the constrains that are specified in the struct User.

	DB = db.AutoMigrate(User{})
	DB = db.AutoMigrate(&TODO{})
	DB = db.AutoMigrate(&Album{})
	DB = db.AutoMigrate(&Post{})
	DB = db.AutoMigrate(&Comment{})

	DB.Model(&TODO{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&Album{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&Comment{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE")
}

func CreateUser(id int, name, email, address, geoLat, geoLng, phone, website, company string) {
	db := DB

	user := &User{
		ID:       id,
		Username: name,
		Email:    email,
		Address:  address,
		GeoLat:   geoLat,
		GeoLng:   geoLng,
		Phone:    phone,
		Website:  website,
		Company:  company,
	}

	db.Create(user)
}

func PopulateUsersTable(users []handles.UserResponse) error {
	for _, u := range users {
		CreateUser(
			u.ID,
			u.Username,
			u.Email,
			// To concat Street, Suite, City, and Zipcode and save them in the address field.
			fmt.Sprintf("%s, %s, %s, %s", u.Address.Street, u.Address.Suite, u.Address.City, u.Address.Zipcode),
			u.Address.Geo.Lat,
			u.Address.Geo.Lng,
			u.Phone,
			u.Website,
			// To concat Company's name, CatchPhrase, and BS and save them in the company field.
			fmt.Sprintf("%s, %s, %s", u.Company.Name, u.Company.CatchPhrase, u.Company.Bs),
		)
	}

	return nil
}
