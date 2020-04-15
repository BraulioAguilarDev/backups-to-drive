package main

import (
	"context"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	drive "google.golang.org/api/drive/v2"
)

// App struct
type App struct {
	Cert   []byte
	Token  string
	Config *oauth2.Config
}

// NewCredentials func
func NewCredentials() (*App, error) {
	data, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}

	return &App{
		Cert: data,
	}, nil
}

// GetClient func uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func (app *App) GetClient(ctx context.Context) (*http.Client, error) {
	config, err := google.ConfigFromJSON(app.Cert, drive.DriveScope)
	if err != nil {
		return nil, err
	}

	app.Config = config

	if err := app.TokenCacheFile(); err != nil {
		return nil, err
	}

	tok, err := app.TokenFromFile()
	if err != nil {
		tok = app.GetTokenFromWeb()
		app.SaveToken(tok)
	}

	return app.Config.Client(ctx, tok), nil
}

// Initialize func read google credentials
// generate service instance
func (app *App) Initialize() (*drive.Service, error) {
	ctx := context.Background()

	client, err := app.GetClient(ctx)
	if err != nil {
		return nil, err
	}

	srv, err := drive.New(client)
	if err != nil {
		return nil, err
	}

	return srv, nil
}
