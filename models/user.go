package models

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id                string
	Username          string
	FullName          string
	Email             string
	EncryptedPassword string
	AvatarURL         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func CreateUser(db *pg.DB, username string, fullName string, email string, password string) (*User, error) {
	var users []User
	_, err := db.Query(&users, `SELECT * FROM users WHERE username = ? OR email = ?`, username, email)
	log.Println(err)
	if len(users) > 0 || err != nil {
		description := fmt.Errorf("username or email token")
		return nil, description
	}
	user := &User{
		Id:        bson.NewObjectId().Hex(),
		Username:  username,
		FullName:  fullName,
		Email:     email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	if encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10); err != nil {
		log.Println(err)
		return nil, err
	} else {
		user.EncryptedPassword = string(encryptedPassword)
	}

	err = db.Insert(user)
	if err != nil {
		description := fmt.Errorf("Insert db user failed")
		log.Println(err)
		return nil, description
	}
	return user, nil
}

func AuthUserWithPassword(db *pg.DB, username string, email string, password string) (*User, error) {
	if user, err := FindUserByNameOrEmail(db, username, email); err != nil || user == nil {
		return nil, err
	} else if err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)); err != nil {
		return nil, nil
	} else {
		return user, nil
	}
}

func FindUserByName(db *pg.DB, name string) (*User, error) {
	if user, err := FindUserByNameOrEmail(db, name, ""); err != nil || user == nil {
		return nil, err
	} else {
		return user, nil
	}
}

func FindUserByNameOrEmail(db *pg.DB, name string, email string) (*User, error) {
	var user User
	_, err := db.QueryOne(&user, `SELECT * FROM users WHERE username = ? OR email = ?`, name, email)
	return &user, err
}
