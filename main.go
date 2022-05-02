package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type car struct {
	Model string
	Year  string
}

var cars []car = []car{
	{"Aventador", "2011"},
	{"Huracan", "2014"},
	{"Urs", "2018"},
}

func main() {
	rand.Seed(time.Now().Unix())

	api := echo.New()

	api.GET("/cars", getcar)
	api.POST("/cars", getcar)
	api.PUT("/cars", getcar)

	port := os.Getenv("PORT")

	if port == "" {
		port = "25255"
	}

	api.Start(":" + port)

}

func getcar(c echo.Context) error {
	index := rand.Intn(len(cars))
	return c.JSON(http.StatusOK, cars[index])
}
