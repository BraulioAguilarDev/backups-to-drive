package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
)

var (
	inputPath  = flag.String("i", "", "input file path")
	outputFile = flag.String("o", "", "output filename")
	folderName = flag.String("f", "", "folder name")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}

	fmt.Printf("Read file: %s\n", *inputPath)
	outputTitle := *outputFile
	if outputTitle == "" {
		outputTitle = filepath.Base(*inputPath)
	}
	fmt.Printf("Output name: %s\n", outputTitle)

	ext := filepath.Ext(*inputPath)
	mimeType := "application/octet-stream"
	if ext != "" {
		mimeType = mime.TypeByExtension(ext)
	}
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	fmt.Printf("Mime : %s\n", mimeType)

	uploadFile(srv, outputTitle, "", *folderName, mimeType, *inputPath)

}
