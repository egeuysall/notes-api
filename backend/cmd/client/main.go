package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

const apiKey = "Ce8doLbtQq3q8kVjOsGXM0vTd6B9UcDcxEYExB6BqrWXWluEa49MZRsagXL6haPVUGYFP43OOBxdhJxIKWNhfWXqfoyXQFrAVM5fPH5WZqOhYCWLu1XrNjoVc5cC5sOVS8er0EZNXWCUpIeleG7Ya54zBiLXPQE9xhJBWqXyxdaNjuu9qj0rnMIFsK7bL33xNs0hkmLwxWFS0U4Wyc8YR4019BYFbpZ7wz5oFBATIAkecX7w3YDZCsnXNztQSxQR"

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), `Usage:
  -n, --note		string		Add a note
  -i, --id			string		ID of the note to get or delete
  -d, --delete		boolean		Delete the note with given ID
  -v, --version		boolean		Show application version
`)
}

func postNote(note string) {
	quoted := fmt.Sprintf("%q", note)

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/notes", bytes.NewBufferString(quoted))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating POST request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making POST request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(os.Stderr, "Server returned status %d: %s\n", resp.StatusCode, string(bodyBytes))
		os.Exit(1)
	}

	var result struct {
		Notes []struct {
			ID string `json:"id"`
		} `json:"notes"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding POST response: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Your note is posted on http://localhost:3000/notes/%s\n", result.Notes[0].ID)
}

func getNote(id string) {
	url := fmt.Sprintf("http://localhost:8080/v1/notes/%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating GET request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making GET request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(os.Stderr, "Server returned status %d: %s\n", resp.StatusCode, string(bodyBytes))
		os.Exit(1)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading GET response body: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Note content: %s\n", string(bodyBytes))
}

func deleteNote(id string) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/v1/notes/%s", id), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating DELETE request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making DELETE request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(os.Stderr, "Server returned status %d: %s\n", resp.StatusCode, string(bodyBytes))
		os.Exit(1)
	}

	fmt.Printf("Note with ID %s successfully deleted.\n", id)
}

func main() {
	var note string
	var noteID string
	var deleteFlag bool
	var showVersion bool
	var version string = "v1.0.0"

	flag.StringVar(&note, "n", "", "Add a note")
	flag.StringVar(&note, "note", "", "Add a note")

	flag.StringVar(&noteID, "i", "", "ID of the note to get or delete")
	flag.StringVar(&noteID, "id", "", "ID of the note to get or delete")

	flag.BoolVar(&deleteFlag, "d", false, "Delete the note with given ID")
	flag.BoolVar(&deleteFlag, "delete", false, "Delete the note with given ID")

	flag.BoolVar(&showVersion, "v", false, "Show application version")
	flag.BoolVar(&showVersion, "version", false, "Show application version")

	flag.Usage = usage
	flag.Parse()

	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	opsCount := 0
	if note != "" {
		opsCount++
	}
	if noteID != "" {
		opsCount++
	}

	if opsCount == 0 {
		fmt.Fprintln(os.Stderr, "Error: You must specify either -n/--note or -i/--id.")
		flag.Usage()
		os.Exit(1)
	}

	if opsCount > 1 {
		fmt.Fprintln(os.Stderr, "Error: Flags -n/--note and -i/--id are mutually exclusive.")
		flag.Usage()
		os.Exit(1)
	}

	if note != "" {
		postNote(note)
		return
	}

	if noteID != "" {
		if deleteFlag {
			deleteNote(noteID)
		} else {
			getNote(noteID)
		}
		return
	}
}