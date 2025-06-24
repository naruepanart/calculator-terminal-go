package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	outputFile = "abc.exe"
)

func main() {
	// Read the URL from the update.txt file.
	urlBytes, err := os.ReadFile("update.txt")
	if err != nil {
		fmt.Println("Failed to read URL from update.txt:", err)
		return
	}
	dropboxURL := strings.TrimSpace(string(urlBytes))

	// Attempt to terminate the running abc.exe process to update it.
	terminateCmd := exec.Command("taskkill", "/F", "/IM", outputFile)
	err = terminateCmd.Run()
	if err != nil {
		fmt.Println("Warning: Unable to terminate abc.exe or it was not running:", err)
	}

	// Download the updated abc.exe from Dropbox.
	err = downloadFile(outputFile, dropboxURL)
	if err != nil {
		fmt.Println("Update failed:", err)
		return
	}

	// Launch the new version of abc.exe.
	startCmd := exec.Command("./" + outputFile)
	startCmd.Stdout = os.Stdout
	startCmd.Stderr = os.Stderr
	startCmd.Stdin = os.Stdin

	err = startCmd.Start()
	if err != nil {
		fmt.Println("Failed to start the updated abc.exe:", err)
		return
	}

	os.Exit(0)
}

// downloadFile downloads a file from a given URL and saves it to a given file path.
func downloadFile(filePath string, url string) error {
	// Get the file from the URL with an HTTP GET request.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file that we'll be saving the response to.
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the response body to the file.
	_, err = io.Copy(file, resp.Body)
	return err
}
