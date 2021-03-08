package web

import "github.com/labstack/echo/v4"

// WebRoutes is where we register any Web routes the application should know about
func WebRoutes() *echo.Echo {
	w := createWebMux()
	// assets for the webapp
	w.Static("/public", "public")
	// Home landing page
	w.GET("/", HomeGET)
	// not found handler to deal with bad route requests
	echo.NotFoundHandler = NotFoundGET
	return w
}
