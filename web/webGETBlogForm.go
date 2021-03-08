package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// BlogFormGET loads the creating a blog form
func BlogFormGET(c echo.Context) (err error) {

	vm := make(map[string]interface{})
	vm["DocTitle"] = "SneakyMommies | Admin | Create"
	pusher, ok := c.Response().Writer.(http.Pusher)
	if ok {
		if err = pusher.Push("/assets/**/*.css", nil); err != nil {
			log.Error("Failed to push: %v", err)
		}
		if err = pusher.Push("/assets/**/*.js", nil); err != nil {
			log.Error("Failed to push: %v", err)
		}
		if err = pusher.Push("/assets/**/*.jpg", nil); err != nil {
			log.Error("Failed to push: %v", err)
		}
	}
	// Set the correct headers
	c.Response().Header().Add(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	c.Response().Header().Set(echo.HeaderAccessControlMaxAge, "max-age=3600")
	c.Render(http.StatusOK, "blog_create_template.html", vm)
	return nil
}
