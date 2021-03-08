package api

import (
	"cinkin/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// APITestGET is the way we check if our API router works
func APITestGET(c echo.Context) (err error) {
	vm := make(map[string]interface{})
	vm["DocTitle"] = "SneakyMommies | Home"
	vm["Data"] = "Ase Deliri's Version of a PIMPED out router"
	c.JSONPretty(http.StatusOK, vm, " ")
	return nil
}

func FireBaseTest(c echo.Context) (err error) {

	// Get a Context.
	ctx := context.Background()
	// Get a Client
	client := models.FsClient(ctx)
	// defer and close client to prevent resource leaks
	defer client.Close()

	// time it takes
	start := time.Now()
	// initialize struct for form content

	dsnap, err := client.Collection("smb").Doc("iputNbJXKcbzavhhXelBsiput").Get(ctx)
	if err != nil {
		fmt.Println(err)
	}
	var x models.Blog
	dsnap.DataTo(&x)
	fmt.Printf("Document data: %#v\n", c)
	m := make(map[string]interface{})
	m["x"] = x
	m["time"] = time.Since(start)

	c.JSONPretty(http.StatusOK, m, " ")
	return nil
}
