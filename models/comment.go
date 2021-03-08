// models is where we struct out the varios data structures for the app
package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/thedevsaddam/govalidator"
)

// Comment is the for ratings section of her recipes
type Comment struct {
	ID         string    `firestore:"CommentID" json:"comment_id"`
	Name       string    `firestore:"Name" json:"name"`
	Email      string    `firestore:"Email" json:"email"`
	Rating     int       `firestore:"Rating" json:"rating"`
	Body       string    `firestore:"Body" json:"body"`
	IPAddress  string    `firestore:"IpAddress" json:"ip_address"`
	CreatedAt  time.Time `firestore:"CreatedAt" json:"created_at"`
	Approved   bool      `firestore:"Approved" json:"approved"`
	Visible    bool      `firestore:"Visible" json:"visible"`
	ApprovedAT time.Time `firestore:"ApprovedAt" json:"approved_at"`
	Agent      string    `firestore:"Agent" json:"agent"`
}

//ValidateComment is a function to check incoming data for correctness
func ValidateComment(c *Comment) {
	rules := govalidator.MapData{
		"id":     []string{"required", "min:3"},
		"name":   []string{"required", "min:4", "max:20"},
		"email":  []string{"required", "min:4", "email"},
		"body":   []string{"required", "min:3"},
		"rating": []string{"required"},
	}
	opts := govalidator.Options{
		Data:  &c,
		Rules: rules,
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) >= 0 {
		data, _ := json.MarshalIndent(e, "", " ")
		fmt.Println(string(data))
	}
}

// NewComment Creates and initializes a new struct for comments
func NewComment() *Comment {
	return &Comment{
		ID:         "",
		Name:       "",
		Email:      "",
		Rating:     0,
		Body:       "",
		IPAddress:  "",
		CreatedAt:  time.Now(),
		Approved:   false,
		Visible:    false,
		ApprovedAT: time.Time{},
		Agent:      "",
	}
}
