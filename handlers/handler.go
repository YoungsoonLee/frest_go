package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func JSONHTTPErrorHandler(err error, c echo.Context) {
	var msg interface{}
	code := http.StatusInternalServerError
	msg = http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		//logrus.Debug(he.Message)
	}

	if !c.Response().Committed {
		c.JSON(code, map[string]interface{}{
			"status":     "fail",
			"statusCode": code,
			"message":    msg,
		})
	}
}
