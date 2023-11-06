package api

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gobase.com/base/pkg/registering"
)

func RegisterUser(service registering.RegisteringInterface) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := new(registering.User)
		//values := echo.QueryParamsBinder(c)
		//values.String("name", &user.FirstName)

		err := c.Bind(user)
		if err != nil {
			writeError(c, errors.New("Failed to read content"), nil)
			return nil
		}
		id, regErrors := service.RegisterUser(*user)
		if len(regErrors) > 0 {
			writeErrors(c, regErrors, nil)
			return nil
		}
		writeJSON(c, id)
		return nil
	}
}
