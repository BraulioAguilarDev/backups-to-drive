package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/oauth2"
)

// GetTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func (app *App) GetTokenFromWeb() *oauth2.Token {
	authURL := app.Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)
	fmt.Printf("Enter Verfication Code:\n")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		fmt.Printf("Unable to read authorization code %v", err)
	}

	tok, err := app.Config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Unable to retrieve token from web %v", err)
	}

	return tok
}

// TokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func (app *App) TokenCacheFile() error {
	usr, err := user.Current()
	if err != nil {
		app.Token = usr.HomeDir

		return err
	}

	tokenCacheDir := ".credentials"
	os.MkdirAll(tokenCacheDir, 0700)

	app.Token = filepath.Join(tokenCacheDir, url.QueryEscape("drive-api-cert.json"))

	return err
}

// TokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func (app *App) TokenFromFile() (*oauth2.Token, error) {
	f, err := os.Open(app.Token)
	if err != nil {
		return nil, err
	}

	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()

	return t, err
}

// SaveToken uses a file path to create a file and store the
// token in it.
func (app *App) SaveToken(token *oauth2.Token) {
	f, err := os.Create(app.Token)
	if err != nil {
		fmt.Printf("Unable to cache oauth token: %v", err)
	}

	defer f.Close()

	json.NewEncoder(f).Encode(token)
}
