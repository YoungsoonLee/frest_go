package route

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"

	"frest_echo/api"
	"frest_echo/db"
	"frest_echo/handlers"
	myMw "frest_echo/middleware"
)

//check header
func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") == "application/json" {
			h.ServeHTTP(w, r)
		} else {
			data := make(map[string]interface{})

			data["status"] = "fail"
			data["statusCode"] = "400"
			data["message"] = "need to set Header Content-Type: application/json"
			doc, _ := json.MarshalIndent(data, "", " ")

			logrus.Debug(string(doc))

			w.WriteHeader(http.StatusBadRequest)
			w.Write(doc)
		}

	})
}

func Init() *echo.Echo {

	e := echo.New()

	e.Debug = true

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Recover())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	e.HTTPErrorHandler = handler.JSONHTTPErrorHandler

	// Set Custom MiddleWare
	e.Use(myMw.TransactionHandler(db.Init()))
	e.Use(echo.WrapMiddleware(middleware))

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/", api.GetDefault())
		v1.GET("/users", api.GetAllUsers())
		v1.POST("/users", api.CreateUser())

		//v1.GET("/members", api.GetMembers())
		//v1.GET("/members/:id", api.GetMember())
	}

	return e
}
