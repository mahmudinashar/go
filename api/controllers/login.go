package controllers

import (
	"net/http"

	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/mahmudinashar/go/api/auth"
	"github.com/mahmudinashar/go/api/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginInputParam struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type LoginOutputParam struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (server *Server) Login(c echo.Context) error {
	var output LoginOutputParam
	body := new(LoginInputParam)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// convert stuct to json
	b, err := json.Marshal(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user := models.User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	output.Email = user.Email[0:3] + ` **** ` + user.Email[len(user.Email)-5:]
	output.Token = token
	return c.JSON(http.StatusOK, output)
}

func (server *Server) SignIn(email, password string) (string, error) {
	var err error
	user := models.User{}
	err = server.DB.Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
