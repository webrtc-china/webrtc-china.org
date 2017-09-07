package models

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"shou.tv/config"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"webrtc-china.org/session"
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

func CreateUser(ctx context.Context, username string, fullName string, email string, password string) (*User, error) {
	var users []User
	db := session.Database(ctx)
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

func AuthUserWithPassword(ctx context.Context, username string, email string, password string) (*User, error) {
	if user, err := FindUserByNameOrEmail(ctx, username, email); err != nil || user == nil {
		return nil, err
	} else if err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)); err != nil {
		return nil, nil
	} else {
		return user, nil
	}
}

func FindUserByName(ctx context.Context, name string) (*User, error) {
	if user, err := FindUserByNameOrEmail(ctx, name, ""); err != nil || user == nil {
		return nil, err
	} else {
		return user, nil
	}
}

func FindUserByNameOrEmail(ctx context.Context, name string, email string) (*User, error) {
	var user User
	_, err := session.Database(ctx).QueryOne(&user, `SELECT * FROM users WHERE username = ? OR email = ?`, name, email)
	return &user, err
}

func (user *User) Authentication(ctx context.Context) *http.Cookie {
	expire := time.Now().UTC().Add(time.Minute * 30)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        user.Id,
		ExpiresAt: expire.Unix(),
	})

	key := sha256.Sum256([]byte(user.EncryptedPassword))
	tokenString, _ := token.SignedString(key[:])
	cookie := http.Cookie{Name: "Authorization", Value: tokenString, Path: "/", MaxAge: int((time.Minute * 30).Seconds()), Secure: config.Environment == "production"}
	return &cookie
}

func AuthenticateUserWithToken(ctx context.Context, tokenString string) (*User, error) {
	var user *User = nil
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			description := fmt.Errorf("Unauthorized, maybe email or password error")
			return nil, description
		}
		if claims, ok := token.Claims.(jwt.MapClaims); !ok {
			return nil, fmt.Errorf("Unauthorized, token")
		} else {
			userId := claims["jti"]

			_, err := session.Database(ctx).QueryOne(user, `SELECT * FROM users WHERE id = ?`, userId)
			if err != nil || user == nil {
				return nil, fmt.Errorf("Unauthorized, transaction error")
			}
			key := sha256.Sum256([]byte(user.EncryptedPassword))
			return key[:], nil
		}
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Unauthorized, valid failed")
	}
	return user, nil
}
