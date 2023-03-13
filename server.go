package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pchchv/golog"
)

// Checks that the server is up and running
func pingHandler(c echo.Context) error {
	message := "User balance service. Version 0.0.1"
	return c.String(http.StatusOK, message)
}

func createUserHandler(c echo.Context) error {
	var jsonMap map[string]interface{}
	if err := c.Bind(&jsonMap); err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}
	user, err := createUser(jsonMap)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func updateBalanceHandler(c echo.Context) error {
	var jsonMap map[string]interface{}
	if err := c.Bind(&jsonMap); err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}
	user, err := updateBalance(jsonMap)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func deleteUserHandler(c echo.Context) error {
	id := c.QueryParam("id")
	err := deleteUser(id)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func getUserHandler(c echo.Context) error {
	id := c.QueryParam("id")
	user, err := getUser(id)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// The declaration of all routes comes from it
func routes(e *echo.Echo) {
	e.GET("/", pingHandler)
	e.GET("/ping", pingHandler)
	e.GET("/user", getUserHandler)
	e.POST("/user", createUserHandler)
	e.PATCH("/balance", updateBalanceHandler)
	e.DELETE("/user", deleteUserHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	golog.Fatal(e.Start(":" + getEnvValue("PORT")).Error())
}
