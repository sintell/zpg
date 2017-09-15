package main

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

func userFromToken(t *jwt.Token) (*User, error) {
	uid := int(t.Claims.(jwt.MapClaims)["userId"].(float64))
	u := &User{Id: uid}
	if err := GetDB().Select(u); err != nil {
		return nil, err
	}
	return u, nil
}

func testToken(c echo.Context) error {
	u, err := userFromToken(c.Get("user").(*jwt.Token))
	if err != nil {
		return err
	}
	c.Logger().Printf("%+v", u)
	return nil
}

type SessionPayload struct {
	UserID int
}
