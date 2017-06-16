package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"

	"frest_echo/models"
)

func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		u := new(models.User)

		if err = c.Bind(&u); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if err = u.Validate(); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		tx := c.Get("Tx").(*dbr.Tx)

		user := models.NewUser(u.Username, u.Email, u.Password, u.Permission)

		//check unique email, username

		if err := user.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusCreated, user)
	}
}

func GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*dbr.Tx)

		//position := c.QueryParam("position")
		users := new(models.Users)
		if err = users.Load(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound, "Member does not exists.")
		}

		return c.JSON(http.StatusOK, users)
	}
}
