package databaseconnector

import (
	"fmt"

	"github.com/jackc/pgx/pgtype"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define struct to create DB Table
type User struct {
	CustomerID string `gorm:"primaryKey"`
	Username   string `gorm:"unique"`
	Email      string
	Age        int
	MetaData   pgtype.JSONB `gorm:"type:jsonb" json:"fieldnameofjsonb"`
}

// Create User table in database
func autoMigrateDB(db *gorm.DB) {
	// Perform Database migration
	var err error = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}
}

// Create connection to the database
func connectToPostgreSQL() (*gorm.DB, error) {
	var dsn string = "host=localhost user=postgres password=1234 dbname=Temp port=5433 sslmode=disable"
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Create new user and call insert method to insert new user into table
func createuserWithMetaData(db *gorm.DB, customerId string, username string, email string, age int, metaData string) (*User, error) {
	var jsonData pgtype.JSONB = pgtype.JSONB{}
	var err error = jsonData.Set([]byte(metaData))
	if err != nil {
		fmt.Println("Error converting metaData to bytestream")
		return nil, err
	}
	var newUser User = User{CustomerID: customerId, Username: username, Email: email, Age: age, MetaData: jsonData}
	err = createUser(db, &newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

// Insert user into the table
func createUser(db *gorm.DB, user *User) error {
	var result = db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Connect to DB and initialize new table creation
func CreateDB(tablename string) error { // Why is this tablename
	var db, err = connectToPostgreSQL()
	if err != nil {
		return err
	}
	autoMigrateDB(db)
	return nil
}

// Create user
func CreateUser(customerId string, username string, email string, age int, metaData string) (*User, error) {
	db, err := connectToPostgreSQL()
	if err != nil {
		return nil, err
	}
	user, err := createuserWithMetaData(db, customerId, username, email, age, metaData)
	return user, err
}

// Get user by user id
func GetUserByID(userId string) (*User, error) {
	var db, err = connectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var user User
	var result *gorm.DB = db.First(&user, userId)
	if result.Error != nil {
		return nil, err
	}
	return &user, nil
}

// Get user by meta data
func GetUserByMetaData(metaDataFilter string) (*User, error) {
	db, err := connectToPostgreSQL()
	if err != nil {
		return nil, err
	}
	var user User
	// result := db.First(&user, userID)
	result := db.Where(metaDataFilter).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
