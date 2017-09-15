package main

import (
	"time"

	"crypto/sha1"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func createToken(p *SessionPayload) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["userId"] = p.UserID

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func userFromContext(c echo.Context) (*User, error) {
	token := c.Get("user").(*jwt.Token)
	uid := int(token.Claims.(jwt.MapClaims)["userId"].(float64))
	u := &User{Id: uid}
	if err := GetDB().Select(u); err != nil {
		return nil, err
	}
	return u, nil
}

func testToken(c echo.Context) error {
	u, err := userFromContext(c)
	if err != nil {
		return err
	}
	c.Logger().Printf("%+v", u)
	return nil
}

func getHashFromPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}

type SessionPayload struct {
	UserID int
}
