package web

import (
	"cinkin/models"
	"cinkin/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
)

// UploadPath is where the uploads folder resides.
// If that folder is missing the site breaks.
const UploadPath = "uploads"

// BlogFormPOST handles the file submission.
func BlogFormPOST(c echo.Context) (err error) {

	// Get a Context.
	ctx := context.Background()
	// Get a Client
	client := models.FsClient(ctx)
	// defer and close client to prevent resource leaks
	defer client.Close()

	// time it takes
	// start := time.Now()
	// initialize struct for form content
	p := models.NewBlog()

	// capture Raw values
	rawName := c.FormValue("author")
	date := c.FormValue("date")
	rawRating := c.FormValue("aggregateRating")
	rawPrepTime := c.FormValue("prepTime")
	rawCookTime := c.FormValue("cookTime")
	rawRecipeYield := c.FormValue("recipeYield")
	rawCalories := c.FormValue("calories")
	rawCategory := c.FormValue("recipeCategory")
	rawKeywords := c.FormValue("keywords")
	rawHeroAlt := c.FormValue("heroImageTag")
	rawSpinningPlateAlt := c.FormValue("spinningImageTag")
	rawTitle := c.FormValue("blogTitle")
	rawCuisineRecipe := c.FormValue("recipeCuisine")
	rawPost := c.FormValue("blogPost")
	rawMetaTag := c.FormValue("metaTag")
	metaTag := strings.TrimSpace(rawMetaTag)
	// Immediately check for blogTitle
	// if it's missing we can't insert it into firestore
	// return
	if rawTitle == "" {
		return c.Redirect(http.StatusUnprocessableEntity, "/admin/blog/create")
	}
	// sanitize and validate inputs
	if rawName == "" {
		rawName = "Natalia Maurer"
	}
	if len(rawName) > 25 {
		rawName = "Natalia Maurer"
	}

	name := strings.TrimSpace(rawName)
	p.P.Author = name
	p.P.Date = date
	rating := strings.TrimSpace(rawRating)
	p.P.Rating = rating

	pTime := strings.TrimSpace(rawPrepTime)
	p.P.PrepTime = pTime

	cTime := strings.TrimSpace(rawCookTime)
	p.P.CookTime = cTime

	// convert prepTime into int
	pIntTime, errPtime := strconv.Atoi(pTime)
	if err != nil {
		// handle error
		fmt.Println(errPtime)
		// os.Exit(2)
	}
	// convert cook time into int
	cIntTime, errCtime := strconv.Atoi(cTime)
	if err != nil {
		// handle error
		fmt.Println(errCtime)
		// os.Exit(2)
	}

	tTime := strconv.FormatInt(int64(pIntTime+cIntTime), 10)
	p.P.TotalTime = tTime

	yield := strings.TrimSpace(rawRecipeYield)
	p.P.RecipeYield = yield

	cals := strings.TrimSpace(rawCalories)
	p.P.Calories = cals

	cat := (strings.Title(strings.ToLower(rawCategory)))
	p.P.RecipeCategory = cat

	kw := strings.TrimSpace(rawKeywords)

	p.P.Keywords = strings.Split(kw, ",")

	t := (strings.Title(strings.ToLower(rawTitle)))
	p.P.Title = t

	cr := (strings.Title(strings.ToLower(rawCuisineRecipe)))
	p.P.RecipeCuisine = cr

	d := strings.TrimSpace(rawPost)
	p.P.Post = d

	p.P.MetaTag = metaTag

	// now we need to call for FormParams to capture
	// the dynamic js fields from the front end
	submittedForm, err := c.FormParams()
	if err != nil {
		c.Logger().Error("Error with form Params method call", err.Error())
	}

	// result is a slice of a map to a string to capture
	// "[0]ingredients", "[1]amount", etc
	var result []map[string]string
	var typeName string
	for k, v := range submittedForm {
		// this regular expression compiles for the pattern of brackets and numbers
		re := regexp.MustCompile(typeName + "\\[([0-9]+)\\]\\[([a-zA-Z]+)\\]")
		// matches checks to see a key within the form against the regular expression
		matches := re.FindStringSubmatch(k)
		if len(matches) >= 3 {
			index, _ := strconv.Atoi(matches[1])
			for index >= len(result) {
				result = append(result, map[string]string{})
			}
			result[index][matches[2]] = v[0]
		}
	}

	// once we've separated the results we need to iterate through them
	// to find individual ingredients, steps, urls, amounts, names for the
	// struct to match
	for _, v := range result {
		for pkg, item := range v {
			if pkg == "ingredients" {
				str := strings.TrimSpace(item)
				s := (strings.Title(strings.ToLower(str)))
				p.P.Ingredients = append(p.P.Ingredients, s)
			}
			if pkg == "amount" {
				str := strings.TrimSpace(item)
				s := (strings.Title(strings.ToLower(str)))
				p.P.Amount = append(p.P.Amount, s)
			}
			if pkg == "steps" {
				str := strings.TrimSpace(item)
				s := (strings.Title(strings.ToLower(str)))
				p.P.PrepSteps = append(p.P.PrepSteps, s)
			}
			if pkg == "urlName" {
				str := strings.TrimSpace(item)
				p.P.URLNames = append(p.P.URLNames, str)
			}
			if pkg == "url" {
				s := (strings.ToLower(strings.TrimSpace(item)))
				p.P.URLs = append(p.P.URLs, s)
			}
		}
	}
	// Let's convert the ingredients and amounts into a single map
	// carefully, we use the ingredients as a key since it's likely
	// that there will be 2 amounts of the same value and thus creating
	// a key collision
	mAIngs, err2 := utils.Mapify(p.P.Ingredients, p.P.Amount)
	if err2 != nil {
		c.Logger().Error("Error from Mapify function", err2.Error())
	}
	// copy the map we just created into our field in the struct
	for k, v := range mAIngs {
		p.P.IngAmounts[k] = v
	}

	// Let's convert the URLNames and Urls into a single map
	// carefully, we use the url as a key since it's likely
	// that there will be 2 names of the same value and thus creating
	// a key collision
	nu, err3 := utils.Mapify(p.P.URLs, p.P.URLNames)
	if err3 != nil {
		c.Logger().Error("Error from Mapify function", err3.Error())
	}

	for k, v := range nu {
		p.P.UNames[k] = v
	}

	// capture the file for the HERO Image
	file, err := c.FormFile("heroImage")
	if err != nil {
		c.Logger().Error("Error getting the hero file: ", err.Error())
	}
	// grab the file and save it in the uploads directory
	// give the file a random name in case two files have the same name
	filename, err := utils.FileToDirectory(file, UploadPath)
	if err != nil {
		c.Logger().Error("Error storing the file from the FileToDirectory function: ", err.Error())
	}
	imDir := UploadPath + "/" + filename

	// adding the file to gCloudStorage
	gcloudURL, err := utils.UploadImageToGCloud(imDir)
	if err != nil {
		c.Logger().Error("Error pushing the file up to gCloud ", err.Error())
	}
	p.P.HeroImage = gcloudURL
	// Hero ALT tag
	heroAlt := strings.TrimSpace(rawHeroAlt)
	p.P.HeroAlt = heroAlt

	// capture file for Spinning Plate IMAGE
	sfile, serr := c.FormFile("spinningPlate")
	if serr != nil {
		c.Logger().Error("Error getting the spinning Plate file: ", err.Error())
	}
	sfilename, serr := utils.FileToDirectory(sfile, UploadPath)
	if serr != nil {
		c.Logger().Error("Error storing the spinning plate file from the FileToDirectory function: ", serr.Error())
	}
	spi := UploadPath + "/" + sfilename
	// adding the file to gCloud Storage
	sgcloudURL, serr := utils.UploadImageToGCloud(spi)
	if serr != nil {
		c.Logger().Error("Error pushing the file up to gCloud ", serr.Error())
	}
	p.P.SpinningPlateURL = sgcloudURL

	// Hero ALT tag
	SpinningPlateAlt := strings.TrimSpace(rawSpinningPlateAlt)
	p.P.SpinningPlateAlt = SpinningPlateAlt

	// READ FILES for mulitple images for the slider
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.Logger().Error("Multiple forms for Slider were not included", err.Error())
	}
	// get the FILES
	files := form.File["sliderImages"]
	// get the length of it
	numFiles := len(files)
	// Lock mechanism to make writes work without data races
	var lock = sync.RWMutex{}

	// if there are no multiple images uploaded, we can skip this step
	if numFiles > 0 {
		for _, file := range files {
			filename, err := utils.FileToDirectory(file, UploadPath)
			if err != nil {
				c.Logger().Error("Error storing the file from the FileToDirectory function: ", err.Error())
			}
			sis := UploadPath + "/" + filename
			// adding the file to gCloud Storage
			gcloudURL, err := utils.UploadImageToGCloud(sis)
			if err != nil {
				c.Logger().Error("Error pushing the file up to gCloud ", err.Error())
			}
			// set lock to prevent multiple submissions
			lock.Lock()
			p.P.SliderImages = append(p.P.SliderImages, gcloudURL)
			lock.Unlock()
		}
	}

	// delete sponsored images upload folder
	srferr := os.RemoveAll(UploadPath)
	if srferr != nil {
		c.Logger().Error("Couldn't remove all files from kUploads directory", serr.Error())
	}
	// recreate the kUploads directory
	newpath := filepath.Join(".", UploadPath)
	os.MkdirAll(newpath, os.ModePerm)

	// Create the SLUG for the post
	tSlug := utils.Slug(p.P.Title)
	// split the date entered in the form and separate the parts
	SplitDate := strings.Split(p.P.Date, "/")
	p.P.Year = SplitDate[2]
	p.P.Month = SplitDate[0]
	p.P.Day = SplitDate[1]
	// send to FireStore
	cref := client.Collection("smb").Doc(tSlug)
	_, err5 := cref.Set(ctx, p)
	if err5 != nil {
		c.Logger().Error(err5)
		fmt.Println(err5)
	}

	c.Redirect(http.StatusSeeOther, "/")
	return nil
}
