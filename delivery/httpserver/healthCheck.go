package httpserver

import "github.com/labstack/echo/v4"

func (s Server) healthChech(c echo.Context) error {

	return c.JSON(200, echo.Map{
		"message": "everithings is OK",
	})
}
