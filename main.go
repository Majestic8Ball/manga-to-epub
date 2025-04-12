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
// argument should be the dir of a overall manga, then it needs to recurse through all the chapters
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
	// Recurse through dir to get the chapters, it is safe to do so since we already validated
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if is dir and is direct child of the dir of overall manga
		if info.IsDir() && filepath.Dir(path) == dir && path != dir {
			fmt.Printf("Found chapter: %s\n", path)
			// Process chapter
		}

		// Continue walking
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1) // Exit with error
	}
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
