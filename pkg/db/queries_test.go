package db_test

import (
	"fmt"
	"golang_challenge/pkg/db"
	"golang_challenge/pkg/handles"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopulateUsersTable(t *testing.T) {
	db.InitDatabase()

	var userTest0 handles.UserResponse
	userTest0.ID = 123
	userTest0.Name = "John Doe"
	userTest0.Username = "jdoe"
	userTest0.Email = "jdoe@example.com"
	userTest0.Address.Street = "123 Main St"
	userTest0.Address.Suite = "Apt 4B"
	userTest0.Address.City = "Anytown"
	userTest0.Address.Zipcode = "12345"
	userTest0.Address.Geo.Lat = "37.7749"
	userTest0.Address.Geo.Lng = "-122.4194"
	userTest0.Phone = "555-123-4567"
	userTest0.Website = "example.com"
	userTest0.Company.Name = "Example Inc."
	userTest0.Company.CatchPhrase = "Building better solutions"
	userTest0.Company.Bs = "Innovative"

	var userTest1 handles.UserResponse
	userTest1.ID = 124
	userTest1.Name = "Joca"
	userTest1.Username = "joca"
	userTest1.Email = "joca@example.com"
	userTest1.Address.Street = "124 Main St"
	userTest1.Address.Suite = "Apt 4A"
	userTest1.Address.City = "Anyhometown"
	userTest1.Address.Zipcode = "12346"
	userTest1.Address.Geo.Lat = "37.7748"
	userTest1.Address.Geo.Lng = "-122.4193"
	userTest1.Phone = "555-123-4568"
	userTest1.Website = "example2.com"
	userTest1.Company.Name = "Example2 Inc."
	userTest1.Company.CatchPhrase = "Building better better solutions"
	userTest1.Company.Bs = "Innovative2"

	// Here I'am adding the previous handles.UserResponse objects to an slice test PopulateUsersTable.
	var usersTest []handles.UserResponse
	usersTest = append(usersTest, userTest0)
	usersTest = append(usersTest, userTest1)
	db.PopulateUsersTable(usersTest)

	var resultTest0, resultTest1 db.User
	db.DB.Where("id = ?", userTest0.ID).First(&resultTest0)
	db.DB.Where("id = ?", userTest1.ID).First(&resultTest1)

	// This auxiliar variables are used to concat information about user's address and user's company.
	expectAddress := fmt.Sprintf("%s, %s, %s, %s", userTest0.Address.Street, userTest0.Address.Suite, userTest0.Address.City, userTest0.Address.Zipcode)
	expectCompany := fmt.Sprintf("%s, %s, %s", userTest0.Company.Name, userTest0.Company.CatchPhrase, userTest0.Company.Bs)
	// Verify that insertion of userTest0 was succeful and all the information is right.
	assert.Equal(t, userTest0.Username, resultTest0.Username)
	assert.Equal(t, userTest0.Email, resultTest0.Email)
	assert.Equal(t, expectAddress, resultTest0.Address)
	assert.Equal(t, userTest0.Address.Geo.Lat, resultTest0.GeoLat)
	assert.Equal(t, userTest0.Address.Geo.Lng, resultTest0.GeoLng)
	assert.Equal(t, userTest0.Phone, resultTest0.Phone)
	assert.Equal(t, userTest0.Website, resultTest0.Website)
	assert.Equal(t, expectCompany, resultTest0.Company)

	expectAddress = fmt.Sprintf("%s, %s, %s, %s", userTest1.Address.Street, userTest1.Address.Suite, userTest1.Address.City, userTest1.Address.Zipcode)
	expectCompany = fmt.Sprintf("%s, %s, %s", userTest1.Company.Name, userTest1.Company.CatchPhrase, userTest1.Company.Bs)
	// Verify that insertion of userTest was succeful and all the information is right.
	assert.Equal(t, userTest1.Username, resultTest1.Username)
	assert.Equal(t, userTest1.Email, resultTest1.Email)
	assert.Equal(t, expectAddress, resultTest1.Address)
	assert.Equal(t, userTest1.Address.Geo.Lat, resultTest1.GeoLat)
	assert.Equal(t, userTest1.Address.Geo.Lng, resultTest1.GeoLng)
	assert.Equal(t, userTest1.Phone, resultTest1.Phone)
	assert.Equal(t, userTest1.Website, resultTest1.Website)
	assert.Equal(t, expectCompany, resultTest1.Company)

	// Delete changes in the database.
	db.DB.Unscoped().Delete(&resultTest0)
	db.DB.Unscoped().Delete(&resultTest1)
}

func TestCreateUser(t *testing.T) {
	err := db.InitDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	user := db.User{
		ID:       11,
		Username: "Jocabed Rios",
		Email:    "joca@example.com",
		Address:  "Av. Chapultepec 123",
		GeoLat:   "37.7749",
		GeoLng:   "-122.4194",
		Phone:    "555-555-1212",
		Website:  "www.example.com",
		Company:  "Example Inc.",
	}

	fmt.Print(user)

	db.CreateUser(user.ID, user.Username, user.Email, user.Address, user.GeoLat, user.GeoLng, user.Phone, user.Website, user.Company)

	var result db.User
	db.DB.Where("id = ?", user.ID).First(&result)

	assert.Equal(t, user.Username, result.Username)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.Address, result.Address)
	assert.Equal(t, user.GeoLat, result.GeoLat)
	assert.Equal(t, user.GeoLng, result.GeoLng)
	assert.Equal(t, user.Phone, result.Phone)
	assert.Equal(t, user.Website, result.Website)
	assert.Equal(t, user.Company, result.Company)

	db.DB.Unscoped().Delete(&result)
}
