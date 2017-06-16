package models

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/bwmarrin/snowflake"
	"github.com/gocraft/dbr"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id           int64     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Permission   string    `json:"permission"` //banned user or normal user
	Confirmed    bool      `json:"confirmed"`
	Is_active    bool      `json:"is_active"`    //active or close
	Is_anonymous bool      `json:"is_anonymous"` //guest or not (for future)
	Created_at   time.Time `json:"created_at"`
	//Updated_at   time.Time `json:"updated_at"`
	//Closed_at    time.Time `json:"closed_at"`
	Updated_at dbr.NullTime `json:"updated_at"`
	Closed_at  dbr.NullTime `json:"updated_at"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.Length(5, 20)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 20)),
		validation.Field(&u.Permission, validation.Required),
	)
}

func NewUser(username string, email string, password string, permission string) *User {
	// snowflake
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		logrus.Debug(err)
		return nil
	}

	// Generate a snowflake ID.
	id := node.Generate()

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	} else {
		return &User{
			Id:         id.Int64(),
			Username:   username,
			Email:      email,
			Password:   string(hashedPassword),
			Permission: permission,
			Created_at: time.Now(),
		}
	}
}

func (u *User) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("user").
		Columns("id", "username", "email", "password", "permission", "created_at").
		Record(u).
		Exec()

	return err
}

type Users []User

func (u *Users) Load(tx *dbr.Tx) error {
	/*
		var condition dbr.Condition
		if position != "" {
			condition = dbr.Eq("position", position)
		}
	*/
	return tx.Select("*").
		From("user").
		LoadStruct(u)
}
