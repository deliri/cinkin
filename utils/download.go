package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DlFile is a function that takes in a url and a name
// downloads a file and returns an error if something goes wrong
func DlFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	// close the create handler to prevent memory leaks
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// close the response or you'll leave memory
	defer resp.Body.Close()
	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
