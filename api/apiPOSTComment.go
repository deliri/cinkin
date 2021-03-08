package api

import (
	"cinkin/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/iterator"
)

// CommentPOSTAPI handles any comments submitted
func CommentPOSTAPI(c echo.Context) (err error) {
	// Get a Context.
	ctx := context.Background()
	// Get a Client
	client := models.FsClient(ctx)
	// defer and close client to prevent resource leaks
	defer client.Close()
	type Comment struct {
		ID         string    `firestore:"CommentID" json:"comment_id" form:"comment_id"`
		Name       string    `firestore:"Name" json:"name" form:"name"`
		Email      string    `firestore:"Email" json:"email" form:"email"`
		Rating     string    `firestore:"Rating" json:"rating" form:"rating"`
		Body       string    `firestore:"Body" json:"body" form:"body"`
		IPAddress  string    `firestore:"IpAddress" json:"ip_address" form:"ip_address"`
		CreatedAt  time.Time `firestore:"CreatedAt" json:"created_at"`
		Approved   bool      `firestore:"Approved" json:"approved"`
		Visible    bool      `firestore:"Visible" json:"visible"`
		ApprovedAT time.Time `firestore:"ApprovedAt" json:"approved_at"`
		Agent      string    `firestore:"Agent" json:"agent"`
	}
	u := &Comment{}
	// Read the Body content
	// var bodyBytes []byte
	// if c.Request().Body != nil {
	// 	bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	// }

	// Restore the io.ReadCloser to its original state
	// c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	err2 := c.Bind(u)
	if err2 != nil {
		c.Logger().Error(err2)
		fmt.Println(err2)
	}
	//title := u.ID
	// blogRef := client.Collection("sneakymommiesblog")
	// commentRef := client.Collection("comments").NewDoc()
	cref := client.Collection("sneakymommiesblog").Doc("afghan-salat").Collection("comments").NewDoc()

	_, err3 := cref.Set(ctx, u)
	if err3 != nil {
		c.Logger().Error(err3)
		fmt.Println(err3)
	}

	fmt.Println("All coments:")
	iter := client.Collection("sneakymommiesblog").Doc("afghan-salat").Collection("comments").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(doc.Data())
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	c.JSON(http.StatusCreated, u)
	return nil
}
