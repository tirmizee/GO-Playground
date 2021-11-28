package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type (
	UserRequest struct {
		Id int `json:"id"`
	}

	UserResponse struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UserValidator struct {
		Validator *validator.Validate
	}
)

func HandlerGet(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func HandlerGetParam(c echo.Context) error {

	_name := c.QueryParam("name")
	_type := c.QueryParam("type")

	response := map[string]string{
		"name": _name,
		"type": _type,
	}

	return c.JSON(http.StatusOK, response)
}

func HandlerGetPath(c echo.Context) error {

	cid := c.Param("cid")
	method := c.Param("method")

	response := map[string]string{
		"cid":    cid,
		"method": method,
	}

	return c.JSON(http.StatusOK, response)
}

func HandlerUserGet(c echo.Context) error {

	defer c.Request().Body.Close()

	userRequest := UserRequest{}

	err := json.NewDecoder(c.Request().Body).Decode(&userRequest)
	if err != nil {
		return err
	}

	userResponse := UserResponse{
		Id:    userRequest.Id,
		Name:  "tirmizee",
		Email: "zee@hotmail.com",
	}

	return c.JSON(http.StatusOK, userResponse)

}

func main() {

	var (
		e = echo.New()
	)

	e.GET("/", HandlerGet)
	e.GET("/param", HandlerGetParam)
	e.GET("/path/:cid/:method", HandlerGetPath)
	e.POST("/user", HandlerUserGet)

	err := e.Start(":8080")
	e.Logger.Fatal(err)

}
