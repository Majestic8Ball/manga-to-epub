package epub

import (
  "archive/zip"
  "fmt"
  "io"
  "os"
  "path/filepath"
)

// EPub being made
type Epub struct {
  file *os.file
  zipWriter *zip.Writer
  title string
  author string
}
