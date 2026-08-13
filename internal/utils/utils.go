// Package utils is used for utilies needed by the vidya-intarweb-playlist-cli
package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go.senan.xyz/taglib"

	vip "github.com/JannoTjarks/vidya-intarweb-playlist-cli/pkg/vidya-intarweb-playlist"
)

func fileExists(fullpath string) (bool, error) {
	_, err := os.Stat(fullpath)
	if err != nil {
		return false, errors.New("file does not exist or the access is blocked by given permissions")
	}

	return true, nil
}

func fileTagged(fullpath string) (bool, error) {
	tags, err := taglib.ReadTags(fullpath)
	if err != nil {
		return false, errors.New("it was not possible to check the Audio Tags")
	}

	if tags["TITLE"] != nil {
		return true, nil
	}

	return false, nil
}

func downloadTrack(fullpath string, downloadURL string) {
	out, fileErr := os.Create(fullpath)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer out.Close()

	httpClient := http.Client{
		Timeout: time.Second * 60,
	}
	req, reqErr := http.NewRequest(http.MethodGet, downloadURL, nil)
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

	_, writeErr := io.Copy(out, res.Body)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}

func writeAudioTags(fullFilename, game, title, arranger, composer string) {
	err := taglib.WriteTags(fullFilename, map[string][]string{
		taglib.Album:    {game},
		taglib.Title:    {title},
		taglib.Arranger: {arranger},
		taglib.Composer: {composer},
		taglib.Comment:  {"Vidya Intarweb Playlist | Video Game Music - Video game music playlist for VIP quality VGM! Listen to the best songs from original soundtracks to remixes/arrangements to high quality rips! - https://www.vipvgm.net"},
	}, 0)

	if err != nil {
		log.Fatal(err)
	}
}

func DownloadTracksByRoster(rosterName string, destinationPath string) {
	var roster = vip.GetRoster(rosterName)
	for _, track := range roster.Tracks {
		var fullpath string
		var downloadURL string
		var fileDescription string

		switch rosterName {
		case "source":
			if track.SourceFile == "" {
				continue
			}
			fullpath = fmt.Sprintf("%s/%s/%s.%s", filepath.Clean(destinationPath),
				rosterName, track.SourceFile, roster.Ext)
			downloadURL = fmt.Sprintf("%s%s%s.%s", roster.URL, "source/", track.SourceFile, roster.Ext)
			fileDescription = fmt.Sprintf("\"%s\" from \"%s\" as source version",
				track.Title, track.Game)
		default:
			fullpath = fmt.Sprintf("%s/%s/%s.%s", filepath.Clean(destinationPath),
				rosterName, track.File, roster.Ext)
			downloadURL = fmt.Sprintf("%s%s.%s", roster.URL, track.File, roster.Ext)
			fileDescription = fmt.Sprintf("\"%s\" from \"%s\" by %s",
				track.Title, track.Game, track.Comp)
		}

		fmt.Printf("Checking local file status: %s\n", fullpath)
		fileExists, _ := fileExists(fullpath)
		if !fileExists {
			fmt.Printf("File does not exist: %s\n", fullpath)
		}
		if fileExists {
			fmt.Printf("File does already exist: %s\n", fullpath)
			fileTagged, _ := fileTagged(fullpath)
			if fileTagged {
				fmt.Printf("File is alrady tagged: %s\n", fullpath)
				fmt.Printf("Download is not needed!\n")
				fmt.Println("")
				continue
			}

			fmt.Printf("File is missing tags: %s\n", fullpath)
		}

		fmt.Printf("Download started: %s\n", fileDescription)
		downloadTrack(fullpath, downloadURL)
		fmt.Printf("Download completed: %s\n", fileDescription)

		fmt.Printf("Writing Audio Tags: %s\n", fileDescription)
		if rosterName != "source" {
			writeAudioTags(fullpath, track.Game, track.Title, track.Arr, track.Comp)
		} else {
			writeAudioTags(fullpath, track.Game, track.SourceTitle, track.Arr, track.Comp)
		}
		fmt.Printf("Writing Audio Tags completed: %s\n", fileDescription)
		fmt.Println("")
	}
}
