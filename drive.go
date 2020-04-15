package main

import (
	"fmt"
	"log"

	drive "google.golang.org/api/drive/v2"
)

// Drive struct
type Drive struct {
	Service     *drive.Service
	FolderName  string
	Title       string
	Description string
	MimeType    string
	FileName    string
	ParentName  string
}

// GetOrCreateFolder func
func (d *Drive) GetOrCreateFolder() string {
	folderID := ""
	if d.FolderName == "" {
		return ""
	}

	q := fmt.Sprintf("title=\"%s\" and mimeType=\"application/vnd.google-apps.folder\"", d.FolderName)

	r, err := d.Service.Files.List().Q(q).MaxResults(1).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve foldername: %s", err)
	}

	if len(r.Items) > 0 {
		folderID = r.Items[0].Id
	} else {
		f := &drive.File{
			Title:       d.FolderName,
			Description: "Auto Create by gdrive-upload",
			MimeType:    "application/vnd.google-apps.folder",
		}

		r, err := d.Service.Files.Insert(f).Do()
		if err != nil {
			fmt.Printf("An error occurred when create folder: %v\n", err)
		}

		folderID = r.Id
	}

	return folderID
}
