package api

import (
	"github.com/labstack/echo/v4"
)

// APIRoutes is where we register any API routes the application should know about
func APIRoutes() *echo.Echo {
	a := createAPIMux()

	// API TEST ENDPOINT
	a.GET("/", APITestGET)
	a.GET("/fb", FireBaseTest)
	// // comment API endpoint
	a.POST("/api/v1/nov-24-2020/submitComment", CommentPOSTAPI)

	return a
}
