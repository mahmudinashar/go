package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mahmudinashar/go/api/auth"
	"github.com/mahmudinashar/go/api/models"
)

type UpdateInputParam struct {
	Nickname string `json:"nickname" form:"nickname" query:"nickname"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type GetInputParam struct {
	Id string `json:"id" form:"id" query:"id"`
}

func (server *Server) CreateUser(c echo.Context) error {

	body := new(UpdateInputParam)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	b, err := json.Marshal(body)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, userCreated)
}

func (server *Server) GetUsers(c echo.Context) error {

	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (server *Server) GetUser(c echo.Context) error {

	body := new(GetInputParam)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uid, err := strconv.Atoi(body.Id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, userGotten)
}

func (server *Server) UpdateUser(c echo.Context) error {
	// bearerToken := c.Request().Header.Get("Authorization")
	r := c.Request()
	body := new(UpdateInputParam)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	b, err := json.Marshal(body)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func (server *Server) DeleteUser(c echo.Context) error {

	r := c.Request()
	body := new(GetInputParam)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uid, err := strconv.Atoi(body.Id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	if tokenID != 0 && tokenID != uint32(uid) {
		return c.JSON(http.StatusUnauthorized, "Unauthorized, only allow delete your {current} username!")
	}
	_, err = user.DeleteAUser(server.DB, uint32(uid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var response string
	response = "User ID : " + strconv.Itoa(uid) + " status deleted"
	return c.JSON(http.StatusOK, response)
}
