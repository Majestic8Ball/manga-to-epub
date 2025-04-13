package epub

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/writingtoole/epub"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Will be fed in a directory with a Chapter, needs to sort the pages -> make the EPub -> put it somewhere(?)
func MakeEPub(dir string, title string, author string, outputFolder string) error {
	manga := epub.New()

	// Loop through the dir adding all the images to the EPub with manga.AddImage
	imgs, err := filepath.Glob(filepath.Join(dir, "*.jpg"))
	if err != nil {
		return fmt.Errorf("error reading image directory: %w", err)
	}

	// Maybe we don't need? Extracting software puts 01, 02 etc so lexci sorting is right
	//sort.Slice(imgs, func(i, j int) bool {
	//  return imgs[i] < imgs[j]
	//})
	sort.Strings(imgs)

	fmt.Printf("Found %d images in directory %s\n", len(imgs), dir)

	for i, img := range imgs {
		imgFile := filepath.Base(img)

		imgData, err := os.ReadFile(img)
		if err != nil {
			fmt.Printf("Warning: Error reading image %s: %v\n", imgFile, err)
			continue
		}

		_, err = manga.AddImage(imgFile, imgData)
		if err != nil {
			fmt.Printf("Warning: Error adding image %s to EPUB: %v\n", imgFile, err)
			continue
		}

		xhtmlContent := fmt.Sprintf(`
    <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/1999/xhtml">
    <html xmlns="http://www.w3.org/1999/xhtml">
    <head><title>Page %d</title></head>
    <body>
      <img src="%s" />
    </body>
    </html>`, i+1, imgFile)

		originalBase := strings.TrimSuffix(imgFile, filepath.Ext(imgFile))
		xhtmlFile := fmt.Sprintf("%s.xhtml", originalBase)

		_, err = manga.AddXHTML(xhtmlFile, xhtmlContent)
		if err != nil {
			fmt.Printf("Warning: Error adding XHTML %s to EPUB: %v\n", xhtmlFile, err)
			continue
		}

		pageName := fmt.Sprintf("Page %d", i+1)
		manga.AddNavpoint(pageName, xhtmlFile, i+1)
	}

	outputPath := filepath.Join(outputFolder, fmt.Sprintf("%s.epub", title))

	// Set metadata
	// Going to need more input from flags; author, identifier
	manga.SetTitle(title)
	manga.AddAuthor(author)
	err = manga.SetUUID(uuid.New().String())
	// epub readers need toc, so we need sections but like for this use we dont rlly need it. try just dummy sections
	if err != nil {
		return fmt.Errorf("error generating uuid: %w", err)
	}

	fmt.Printf("Writing EPUB to: %s\n", outputPath)
	err = manga.Write(outputPath)
	if err != nil {
		return fmt.Errorf("error writing EPUB file: %w", err)
	}

	fmt.Printf("EPUB created successfully at: %s\n", outputPath)
	return nil
}
