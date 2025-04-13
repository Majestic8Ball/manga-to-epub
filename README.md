# Manga to EPUB Converter

A command-line tool for converting manga image folders into EPUB files.
Intended for use with HakuNeko program: https://hakuneko.download/

Can use created EPUB for programs like Yomu and other EPUB readers

NOTE: Probably only works on UNIX Based systems

## Features

- Converts folders of JPG images into properly formatted EPUB files
- Automatically sorts images in numerical order
- Preserves image quality
- Creates EPUB files with proper metadata (title, author, UUID)

# Installation

## Option 1, via Go (easiest)
```
go install github.com/Majestic8Ball/manga-to-epub@latest
```

## Option 2, manually
1. Download the most recent release
2. Make the file executable
```
chmod +x manga-to-epub-*
```
3. Move it to your /bin/
```
sudo mv manga-to-epub /usr/local/bin/
```
# Usage

Requires file structure to be a directory with all the chapters inside of it separated by their own directories, inside these directories we are looking for *.JPG

## Example Usage

```
manga-to-epub -dir /path/to/manga/directory -title Title -author Author
```
This will then automatically go to the directory and create the epubs and place them in a folder it creates in the directory provided called ``/epubs``

# Future Plans
>Probably easy stuff, im just lazy

- Image flag
- Other metadata stuffs

