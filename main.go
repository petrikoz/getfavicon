package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	outdir         = "/tmp"
	googleFavicons = "https://www.google.com/s2/favicons?domain="
)

func main() {

	// Get URL from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter URL: ")
	urlUser, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Parse URL for filename
	urlCleaned, err := url.Parse(urlUser)
	if err != nil {
		log.Fatal(err)
	}
	filepath := outdir + "/" + urlCleaned.Host + ".png"

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Get the data
	urlRequest := googleFavicons + urlUser
	resp, err := http.Get(urlRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Saved as '%s'\n", filepath)

}
