package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/api/drive/v2"
)

func uploadFile(d *drive.Service, title string, description string,
	parentName string, mimeType string, filename string) (*drive.File, error) {
	input, err := os.Open(filename)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}
	// Grab file info
	inputInfo, err := input.Stat()
	if err != nil {
		return nil, err
	}

	parentId := getOrCreateFolder(d, parentName)

	fmt.Println("Start upload")
	f := &drive.File{Title: title, Description: description, MimeType: mimeType}
	if parentId != "" {
		p := &drive.ParentReference{Id: parentId}
		f.Parents = []*drive.ParentReference{p}
	}
	getRate := MeasureTransferRate()

	// progress call back
	showProgress := func(current, total int64) {
		fmt.Printf("Uploaded at %s, %s/%s\r", getRate(current), Comma(current), Comma(total))
	}

	r, err := d.Files.Insert(f).ResumableMedia(context.Background(), input, inputInfo.Size(), mimeType).ProgressUpdater(showProgress).Do()
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}

	// Total bytes transferred
	bytes := r.FileSize
	// Print information about uploaded file
	fmt.Printf("Uploaded '%s' at %s, total %s\n", r.Title, getRate(bytes), FileSizeFormat(bytes, false))
	fmt.Printf("Upload Done. ID : %s\n", r.Id)
	return r, nil
}
