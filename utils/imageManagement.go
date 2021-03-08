package utils

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/dchest/uniuri"
	"github.com/labstack/gommon/log"
)

// FileToDirectory takes an uploaded file and saves in the directory
func FileToDirectory(file *multipart.FileHeader, dir string) (fname string, err error) {
	// open the file that's been handed to you
	fileSource, err := file.Open()
	if err != nil {
		log.Error("Error get the file: ", err)
	}
	// get the file name
	fname = file.Filename
	// Close the file handle
	defer fileSource.Close()
	// Create RANDOM String to add to the fileName because Kraken needs UNIQUE requests
	rand := uniuri.New()
	newRandName := rand + "_" + fname

	// get the directory which will include the full path to the directory and the filename
	fdir := "./" + dir + "/" + newRandName

	// destination to save the file
	fdest, err := os.Create(fdir)
	if err != nil {
		log.Error("Error with saving the file into the server directory", err)
	}
	// close the handle to prevent memory leaks
	defer fdest.Close()
	// Copy the file
	if _, err := io.Copy(fdest, fileSource); err != nil {
		log.Error("Error with copying the file from the source to the desintation ", err)
	}

	return newRandName, nil
}

// UploadImageToGCloud is a function to quickly upload an image to gcloud storage bucket
// returns the URL where that image will be hosted or an error
func UploadImageToGCloud(filename string) (string, error) {
	// ProjectID to make sure it goes in the right bucket
	projectID := "sneakymommies"
	// get a context
	ctx := context.Background()
	// get a client handle for google cloud storage
	gclient, err := storage.NewClient(ctx)
	if err != nil {
		log.Error("Couldn't get a client for google cloud storage: ", err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	// defer the cancel call to turn the handle free
	defer cancel()
	// open the file
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Error("Couldn't Open the file to upload: ", err)
	}
	//FilePath gets you the directory and the file name
	FilePath := file.Name()
	name := FilePath[10:]
	// close the file handle to avoid leaking resources
	defer file.Close()
	// get a writer handle
	wc := gclient.Bucket(projectID).Object(name).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		log.Error("Failed to copy from server to gcloud storage ", err)
	}
	// we close the handle for the gcloud writer
	if err := wc.Close(); err != nil {
		log.Error("Failed to close the writer handle", err)
	}
	// create a url for the name of the saved file
	url := "https://storage.googleapis.com/sneakymommies/" + name
	return url, nil
}

// Kraken is a function that takes a file and sends it to the KRAKEN service
// returns a compressed version of the file
// func Kraken(directory string, filename string, processName string) (string, error) {
// 	// Key
// 	krakenKey := "dc426d7eda279d91c58422f87a3717fc"
// 	// Secret
// 	krakenSecret := "7ec23e9e3555ba8dc804f320cc54258c5bb84ff2"
// 	kr, err := kraken.New(krakenKey, krakenSecret)
// 	if err != nil {
// 		log.Error("Error with kraken.io creation of new instance", err)
// 	}
// 	// Set some parameters for Kraken
// 	params := map[string]interface{}{
// 		"wait":   true,
// 		"lossy":  true,
// 		"format": "jpg",
// 	}
// 	// call the API and send the image
// 	api, err := kr.Upload(params, directory)
// 	if err != nil {
// 		log.Error("Error with kraken.io api call with payload", err)
// 	}
// 	// check the status of the API
// 	if api["success"] != true {
// 		log.Error("Failed kraken Call: ", api["message"])
// 		return "", err
// 	}
// 	log.Print("Success, optimized image URL:", api["kraked_url"])
// 	// if we made it this far, we need to get the url from the image
// 	krakenImageUrl := fmt.Sprintf("%v", api["kraked_url"])
// 	// now we submit a get request to get the new optimized image back
// 	krakenHeroFileName := "./uploads/" + "kraken_" + processName + filename

// 	derr := DlFile(krakenHeroFileName, krakenImageUrl)
// 	if derr != nil {
// 		log.Error("Error with kraken.io download request", derr)
// 	}
// 	// once that file has been saved back on the server we need to open it
// 	krakenCompressedImage, err := os.OpenFile(krakenHeroFileName, os.O_RDWR, 0644)
// 	if err != nil {
// 		log.Error("Couldn't Open the kraken hero file: ", err)
// 	}
// 	// kraken hero image handle has to be closed not to leak memory
// 	defer krakenCompressedImage.Close()
// 	// return the name of the compressed filename and nil error
// 	name := krakenCompressedImage.Name()
// 	return name, nil
// }
