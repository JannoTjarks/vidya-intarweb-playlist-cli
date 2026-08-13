// Package vidyaintarwebplaylist provides basic functions and types for working with the Aersia Community's Vidya Intarweb Playlist
package vidyaintarwebplaylist

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

type Roster struct {
	Changelog string `json:"changelog"`
	URL       string `json:"url"`
	Ext       string `json:"ext"`
	NewID     int    `json:"new_id"`
	Tracks    []track
}

type track struct {
	ID          int    `json:"id"`
	Game        string `json:"game"`
	Title       string `json:"title"`
	Comp        string `json:"comp"`
	Arr         string `json:"arr"`
	File        string `json:"file"`
	SourceID    int    `json:"s_id,omitempty"`
	SourceTitle string `json:"s_title,omitempty"`
	SourceFile  string `json:"s_file,omitempty"`
}

func GetRoster(rosterName string) Roster {
	url, rosterError := getRosterURL(rosterName)
	if rosterError != nil {
		log.Fatalln("The rosterName must be one of the following entries\n\t1. vip\n\t2. source\n\t3. mellow\n\t4. exciled")
	}
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}
	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		log.Fatal(reqErr)
	}
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	roster := Roster{}
	jsonErr := json.Unmarshal(body, &roster)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return roster
}

func getRosterURL(rosterName string) (string, error) {
	switch rosterName {
	case "vip":
		return "https://www.vipvgm.net/roster.min.json", nil
	case "source":
		return "https://www.vipvgm.net/roster.min.json", nil
	case "mellow":
		return "https://www.vipvgm.net/roster-mellow.min.json", nil
	case "exiled":
		return "https://www.vipvgm.net/roster-exiled.min.json", nil
	default:
		return "", errors.New("the rosterName must be one of the following entries\n\t1. vip\n\t2. source\n\t3. mellow\n\t4. exiled")
	}
}
