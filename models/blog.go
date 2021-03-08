package models

import (
	"time"
)

// Blog is the struct that includes the Unique ID which will set the
// document ID in fireStore
type Blog struct {
	ID string   `firestore:"ID" json:"id"`
	P  BlogForm `firestore:"Form" json:"blog_form"`
}

// Blogs is a slice of blog elements
type Blogs []Blog

// Post is the struct for the form that captures the content
type BlogForm struct {
	Slug             string            `firestore:"Slug" json:"slug"`
	Author           string            `firestore:"Author" json:"author"`
	Date             string            `firestore:"Date" json:"date"`
	Year             string            `firestore:"Year" json:"year"`
	Month            string            `firestore:"Month" json:"month"`
	Day              string            `firestore:"Day" json:"day"`
	CreatedAt        time.Time         `firestore:"CreatedAt" json:"created_at"`
	UpdatedAt        time.Time         `firestore:"UpdatedAt" json:"updated_at"`
	DeletedAt        time.Time         `firestore:"DeletedAt" json:"deleted_at"`
	Rating           string            `firestore:"Rating" json:"rating"`
	PrepTime         string            `firestore:"PrepTime" json:"prep_time"`
	CookTime         string            `firestore:"CookTime" json:"cook_time"`
	TotalTime        string            `firestore:"TotalTime" json:"total_time"`
	RecipeYield      string            `firestore:"RecipeYield" json:"recipe_yield"`
	Calories         string            `firestore:"Calories" json:"calories"`
	RecipeCategory   string            `firestore:"RecipeCategory" json:"recipe_category"`
	Keywords         []string          `firestore:"Keywords" json:"key_words"`
	Title            string            `firestore:"Title" json:"title"`
	RecipeCuisine    string            `firestore:"RecipeCuisine" json:"recipe_cuisine"`
	MetaTag          string            `firestore:"MetaTag" json:"meta_tag"`
	Post             string            `firestore:"Post" json:"post"`
	Amount           []string          `firestore:"Amount" json:"amount"`
	Ingredients      []string          `firestore:"Ingredient" json:"ingredient"`
	IngAmounts       map[string]string `firestore:"IngAmounts" json:"ing_amounts"`
	PrepSteps        []string          `firestore:"PrepSteps" json:"prep_steps"`
	HeroImage        string            `firestore:"HeroImageURL" json:"hero_image_url"`
	HeroAlt          string            `firestore:"HeroImageALT" json:"hero_image_alt"`
	SpinningPlateURL string            `firestore:"SpinningPlateURL" json:"spinning_plate_url"`
	SpinningPlateAlt string            `firestore:"SpinningPlateAlt" json:"spinning_plate_atl"`
	SliderImages     []string          `firestore:"SliderImagesURLs" json:"slider_images_urls"`
	SponsoredImages  map[string]string `firestore:"SponsoredImagesURL" json:"sponsored_image_urls"`
	URLNames         []string          `firestore:"URLNames" json:"url_names"`
	URLs             []string          `firestore:"URLs" json:"urls"`
	UNames           map[string]string `firestore:"NUrls" json:"named_urls"`
	Visible          bool              `firestore:"Visible" json:"visible"`
}

func NewBlog() *Blog {
	return &Blog{
		ID: "",
		P: BlogForm{
			Slug:             "",
			Author:           "",
			Date:             "",
			Year:             "",
			Month:            "",
			Day:              "",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Time{}, // time.zero date
			DeletedAt:        time.Time{}, // time.zero date
			Rating:           "",
			PrepTime:         "",
			CookTime:         "",
			TotalTime:        "",
			RecipeYield:      "",
			Calories:         "",
			RecipeCategory:   "",
			Keywords:         []string{},
			Title:            "",
			RecipeCuisine:    "",
			MetaTag:          "",
			Post:             "",
			Amount:           []string{},
			Ingredients:      []string{},
			IngAmounts:       make(map[string]string),
			PrepSteps:        []string{},
			HeroImage:        "",
			HeroAlt:          "",
			SpinningPlateURL: "",
			SpinningPlateAlt: "",
			SliderImages:     []string{},
			SponsoredImages:  make(map[string]string),
			URLNames:         []string{},
			URLs:             []string{},
			UNames:           make(map[string]string),
			Visible:          false,
		},
	}
}
