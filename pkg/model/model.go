package model

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

type User struct {
	ID        string    `json:"Id"`
	Username  string    `json:"username"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserTable ...
func CreateUserTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := db.CreateTable(&User{}, opts)

	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)

		return createError
	}

	log.Printf("Todo table created")

	return nil
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db

	// config.Connect()
	// db = config.GetDB()
	// db.AutoMigrate(&Book{})
}

// CreateUser model
func CreateUser(c *gin.Context){
	var user User
	c.BindJSON(&user)
	username := user.Username
	firstname := user.FirstName
	lastname := user.LastName
	email := user.Email
	password := user.Password
	id := guuid.New().String()


	// var Users []User
	// var prevUserID = 0
	// prevUserID++
	// u.ID = strconv.Itoa(prevUserID)

	insertError := dbConnect.Insert(&User{
		ID:        id,
		Username:  username,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new ue into db, Reason: %v\n", insertError)

		// return
	}

	// Users = append(Users, u)
	fmt.Println("Create new User")
	return 
}

// GetAllUsers model
func GetAllUsers() []User {
	var Users []User
	err := dbConnect.Model(&Users).Select()
	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
	}
	fmt.Println("Get all Users")
	return Users
}

// func GetUserByID(Id int64) (*User , *gorm.DB){
// 	var getBook Book
// 	db:=db.Where("ID = ?", Id).Find(&getUser)
// 	return &getUser, db
// }

func DeleteUser(ID int64) User {
	var user User
	err := dbConnect.Delete(user)

	if err != nil {
		log.Printf("Error while deleting a user, Reason: %v\n", err)

		// return
	}

	return user
}
