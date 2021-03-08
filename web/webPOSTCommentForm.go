package web

import (
	"cinkin/models"
	"cinkin/utils"
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// BlogCommentFormPOST handles any comments submitted
func BlogCommentFormPOST(c echo.Context) (err error) {
	// Get a Context.
	ctx := context.Background()
	// Get a Client
	client := models.FsClient(ctx)
	// defer and close client to prevent resource leaks
	defer client.Close()

	// time it takes
	//start := time.Now()
	// initialize struct for form content
	comment := models.NewComment()

	// capture Raw values
	rawName := c.FormValue("author")
	rawemail := c.FormValue("email")
	rawRating := c.FormValue("rating")
	n := strings.TrimSpace(rawName)
	e := strings.TrimSpace(rawemail)
	goodEmail := utils.ValidateEmail(e)
	var validEmail string
	if goodEmail == true {
		validEmail = e
	}
	rate, err := strconv.Atoi(rawRating)
	if err != nil {
		log.Println("couldn't get the rating number")
	}
	rawBody := c.FormValue("comment")
	b := strings.TrimSpace(rawBody)
	comment.ID = c.FormValue("BlogTitle")
	comment.Body = b
	comment.Name = n
	comment.Email = validEmail
	comment.IPAddress = c.RealIP()
	comment.Agent = c.Request().UserAgent()
	comment.Rating = rate

	c.JSONPretty(http.StatusOK, comment, " ")
	return nil
}
