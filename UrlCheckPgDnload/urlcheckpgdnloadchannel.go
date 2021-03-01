package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// Check a list of web sites url, prints up / down depending on server is up or down
func checkAndSaveBodyChannel(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		// Send err info to the channel
		ch <- fmt.Sprintf("Error: %v, Server at URL: %s is down\n", err, url)

	} else {
		defer resp.Body.Close()

		data := fmt.Sprintf("URL %s Status code: %d \n", url, resp.StatusCode)
		if (resp.StatusCode == 200) {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := "output/"
			file += strings.Split(url, "//")[1]  	// http://www.somesite.com
			file += ".txt"    							// www.somesite.com.txt
			data += fmt.Sprintf("Writing response body to %s \n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				data += fmt.Sprintf(", error: %v", err)
				log.Fatal(err)
			}

		} else {
			data += fmt.Sprintf("Server error: %d at URL: %s is down\n", resp.StatusCode, url)
		}

		ch <- data
	}
}

func main() {
	urls := []string{
		"http://golang.org",
		"http://www.google.com",
		"http://medium.com",
	}

	err := os.RemoveAll("./output")
	if err != nil {
		fmt.Println("Error deleting output/*: ", err)
	}

	err = os.Mkdir("./output", 0664)
	if err != nil {
		fmt.Println("Error creating output dir: ", err)
	}

	ch := make(chan string)
	defer close(ch)

	for _, url := range urls {
		// Start each instance as a goroutine
		go checkAndSaveBodyChannel(url, ch)
		fmt.Println(strings.Repeat("#", 20));
	}

	// Print number of go routines (main + one per url
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine())

	// Read and print all the expected results
	for i:=0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
}