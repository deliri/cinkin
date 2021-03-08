// api is the package for creating the api mux
package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/unrolled/secure"
)

// createAPIMux is how we create a muxer with all the default nice to have features plugged in and ready to be passed around
func createAPIMux() *echo.Echo {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny: true,
	})
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	// Server header
	e.Use(middleware.Recover())

	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	e.Use(middleware.CORS())
	return e
}
