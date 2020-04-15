package main

import (
	"flag"
	"fmt"
	"mime"
	"path/filepath"
)

var (
	inputPath  = flag.String("i", "", "Input file path")
	outputFile = flag.String("o", "", "Output filename")
	folderName = flag.String("f", "", "Folder name")
)

// Reference
// https://developers.google.com/drive/api/v3/quickstart/go
func main() {
	flag.Parse()
	app, err := NewCredentials()
	if err != nil {
		fmt.Printf("Unable to read client secret file. Error: %s", err)
	}

	srv, err := app.Initialize()
	if err != nil {
		fmt.Printf("Unable to get path to cached credential file. Error: %s", err)
	}

	outputTitle := *outputFile
	if len(outputTitle) == 0 {
		outputTitle = filepath.Base(*inputPath)
	}

	ext := filepath.Ext(*inputPath)

	mimeType := "application/octet-stream"
	if len(ext) > 0 {
		mimeType = mime.TypeByExtension(ext)
	}

	if len(mimeType) == 0 {
		mimeType = "application/octet-stream"
	}

	drive := &Drive{
		Service:     srv,
		FolderName:  *folderName,
		Title:       outputTitle,
		Description: "",
		MimeType:    mimeType,
		FileName:    *inputPath,
	}

	_, error := drive.UploadFile()
	if error != nil {
		fmt.Printf("Upload error: %s", error.Error())
	}
}
