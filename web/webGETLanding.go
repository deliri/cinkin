package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HomeGET loads the landing page
func HomeGET(c echo.Context) (err error) {
	// Get a Context.
	// ctx := context.Background()
	// // Get a Client
	// client := models.FsClient(ctx)
	// // defer and close client to prevent resource leaks
	// defer client.Close()

	// // create a blank data type
	// type d map[string]interface{}
	// ds := make([]d, 0, 0)
	// // var bs models.Blogs
	// iter := client.Collection("smb").Documents(ctx)
	// defer iter.Stop() // add this line to ensure resources cleaned up
	// for {
	// 	doc, err := iter.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	var b models.Blog
	// 	if err := doc.DataTo(&b); err != nil {
	// 		fmt.Println("couldn't unmarshal into b")
	// 	}
	// 	// bs = append(bs, b)
	// 	// slice := []string{}
	// 	a := make(map[string]interface{})
	// 	id := b.ID
	// 	author := b.P.Author
	// 	date := b.P.Date
	// 	sp := b.P.SpinningPlateURL
	// 	spalt := b.P.SpinningPlateAlt
	// 	title := utils.Slug(b.P.Title)
	// 	cat := b.P.RecipeCategory
	// 	year := b.P.Year
	// 	month := b.P.Month
	// 	day := b.P.Day
	// 	a["id"] = id
	// 	a["name"] = author
	// 	a["date"] = date
	// 	a["splateurl"] = sp
	// 	a["spaltealt"] = spalt
	// 	a["year"] = year
	// 	a["month"] = month
	// 	a["day"] = day
	// 	a["cat"] = cat
	// 	a["title"] = title
	// 	ds = append(ds, a)
	// }

	// m := make(map[string]interface{})
	// m["DocTitle"] = "SneakyMommies | Landing | Page"
	// m["data"] = ds
	// Set the correct headers
	c.Response().Header().Add(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	c.Response().Header().Set(echo.HeaderAccessControlMaxAge, "max-age=3600")
	c.Render(http.StatusOK, "index.html", nil)
	return nil

}
