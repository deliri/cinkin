package utils

import (
	"fmt"
	"os"
	"path"
)

func ClearDir(dir string) error {
	dirRead, err := os.Open(dir)
	if err != nil {
		return err
	}
	dirFiles, err := dirRead.Readdir(0)
	if err != nil {
		return err
	}

	// Loop over the  files.
	for index := range dirFiles {
		entery := dirFiles[index]

		// Get name of file and its full path.
		filename := entery.Name()
		fullPath := path.Join(dir, filename)

		// Remove the file.
		os.Remove(fullPath)

		fmt.Println(fullPath)

	}

	return nil
}
