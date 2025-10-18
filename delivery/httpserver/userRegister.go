package httpserver

import (
	"fmt"
	"net/http"
	"suggestApp/service/userservice"

	"github.com/labstack/echo/v4"
)

func (s Server) userRegister(c echo.Context) error {
	var UReq userservice.RegisterRequest

	if err := c.Bind(&UReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("error on UReq %s", err.Error()))

	}

	Resp, Uerr := s.userSvc.Register(UReq)
	if Uerr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("error on user %s", Uerr.Error()))
	}
	fmt.Printf("%+v", Resp)
	return c.JSON(http.StatusCreated, Resp)

}
