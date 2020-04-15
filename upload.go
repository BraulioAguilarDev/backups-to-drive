package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
	drive "google.golang.org/api/drive/v2"
)

// UploadFile func
func (d *Drive) UploadFile() (*drive.File, error) {
	input, err := os.Open(d.FileName)
	if err != nil {
		return nil, err
	}

	// Grab file info
	inputInfo, err := input.Stat()
	if err != nil {
		return nil, err
	}

	parentID := d.GetOrCreateFolder()

	fmt.Println("Start upload...")
	f := &drive.File{
		Title:       d.Title,
		Description: d.Description,
		MimeType:    d.MimeType,
	}

	if len(parentID) > 0 {
		p := &drive.ParentReference{
			Id: parentID,
		}

		f.Parents = []*drive.ParentReference{p}
	}

	getRate := MeasureTransferRate()

	// progress call back
	showProgress := func(current, total int64) {
		fmt.Printf("Uploaded at %s, %s/%s\r", getRate(current), Comma(current), Comma(total))
	}

	r, err := d.Service.Files.Insert(f).ResumableMedia(context.Background(), input, inputInfo.Size(), d.MimeType).ProgressUpdater(showProgress).Do()
	if err != nil {
		return nil, err
	}

	// Total bytes transferred
	bytes := r.FileSize
	fmt.Printf("Uploaded '%s' at %s, total %s\n", r.Title, getRate(bytes), FileSizeFormat(bytes, false))
	fmt.Printf("Upload Done. ID : %s\n", r.Id)
	return r, nil
}
