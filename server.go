package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pchchv/golog"
)

func server() {
	e := echo.New()
	golog.Fatal(e.Start(":" + getEnvValue("PORT")).Error())
}
