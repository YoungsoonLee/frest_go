package api

import (
	"net/http"

	"github.com/labstack/echo"

	"frest_echo/config"
)

type Default struct {
	API_VERSION string `json:"API_VERSION"`
	API_NAME    string `json:"API_NAME"`
}

func GetDefault() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		//example for avoke Error
		//return echo.NewHTTPError(500, "test")

		d := &Default{
			API_VERSION: config.API_VERSION,
			API_NAME:    config.API_NAME,
		}

		return c.JSONPretty(http.StatusCreated, d, " ")

	}
}
