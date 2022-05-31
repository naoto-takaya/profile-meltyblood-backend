package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/naoto-takaya/profile-meltyblood-backend/service"
	"github.com/naoto-takaya/profile-meltyblood-backend/validators"
)

func Upload(c echo.Context) error {
	image := new(validators.Image)
	// var image validators.Image
	var err error

	if err = c.Bind(image); err != nil {
		return err
	}
	if err = c.Validate(image); err != nil {
		return err
	}

	// fmt.Print(image.Name)
	// service.Upload(image)

	return c.String(http.StatusOK, "uploaded")
}
