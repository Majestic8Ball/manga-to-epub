package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Majestic8Ball/manga-to-epub/epub"
	"os"
	"path/filepath"
)

const Version = "0.1.0"

// Should handle CLI argument parsing, directory scanning to find manga chapters, coordinating conversion process, and error handling
// argument should be the dir of a overall manga, then it needs to recurse through all the chapters
func main() {
	// Maybe we read directories through String, idk
	dirPtr := flag.String("dir", ".", "used to specify the folder directory")
	titlePtr := flag.String("title", "N/A", "designate title of epub")
	authorPtr := flag.String("author", "N/A", "set author name")
	versionFlag := flag.Bool("version", false, "print version information")

	// Parses Command-Line flag
	flag.Parse()

	if *versionFlag {
		fmt.Printf("manga-to-epub version %s\n", Version)
		os.Exit(0)
	}

	// use the flags value
	dir := *dirPtr
	title := *titlePtr
	author := *authorPtr

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
		// Having errors with the dir -> chapter pipeline, going to have to normalize paths before comparison
		dirAbs, _ := filepath.Abs(dir)
		pathDirAbs, _ := filepath.Abs(filepath.Dir(path))
		// Create output folder
		outputFolder := filepath.Join(dir, "epubs")
		err = os.MkdirAll(outputFolder, 0755)
		if err != nil {
			return fmt.Errorf("error creating output folder: %w", err)
		}

		if info.IsDir() && pathDirAbs == dirAbs && path != dir {
			fmt.Printf("Found chapter: %s\n", path)
			// Process chapter
			// This is going to be from the epub.go func
			chap := filepath.Base(path)
			fullTitle := title + " " + chap

			fmt.Printf("Making Title: %v\n", fullTitle)
			fmt.Printf("Making EPub: %v\n", path)
			err := epub.MakeEPub(path, fullTitle, author, outputFolder)
			if err != nil {
				fmt.Printf("Error creating EPUB: %v\n", err)
				// Continue anyway
				return nil
			}
			fmt.Printf("Successfully created EPUB for %s\n", chap)
		}

		// Continue walking
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1) // Exit with error
	}

	fmt.Println("All chapters processed")
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
