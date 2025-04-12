package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	//"github.com/Majestic8Ball/manga-to-epub/epub"
)

// Should handle CLI argument parsing, directory scanning to find manga chapters, coordinating conversion process, and error handling

func main() {
	// Maybe we read directories through String, idk
	dirPtr := flag.String("dir", ".", "used to specify the folder directory")
	titlePtr := flag.String("title", "N/A", "designate title of epub")

	// Parses Command-Line flag
	flag.Parse()

	// use the flags value
	dir := *dirPtr
	title := *titlePtr

	fmt.Printf("Directory: %s, Title: %s\n", dir, title)
	// Now we need to use that directory to find and validate the directory
	// Once validated feed it into the epub maker file in epub/epub.go
	err := directoryValidator(dir)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1) // Exit with error
	}
	// Process images to put to epub
}

func directoryValidator(dir string) error {
	// os.Stat returns an interface and an err
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return errors.New("directory does not exist")
	}
	if err != nil {
		return fmt.Errorf("error accessing directory: %w", err)
	}
	if !info.IsDir() {
		return errors.New("path is not a directory")
	}

	// If here, directory exist
	fmt.Printf("Processing directory: %s\n", dir)

	// Success
	return nil
}
