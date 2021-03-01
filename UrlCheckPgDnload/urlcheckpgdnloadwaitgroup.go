package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

// Check a list of web sites url, prints up / down depending on server is up or down

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Server at URL: %s is down\n", url);

	} else {
		defer resp.Body.Close()

		fmt.Printf("URL %s Status code: %d \n", url, resp.StatusCode)
		if (resp.StatusCode == 200) {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := "output/"
			file += strings.Split(url, "//")[1]  	// http://www.somesite.com
			file += ".txt"    							// www.somesite.com.txt
			fmt.Printf("Writing response body to %s \n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}

		} else {
			fmt.Printf("Server error: %d at URL: %s is down\n", resp.StatusCode, url);
		}
	}

	// Indicate we are done with the task
	wg.Done()
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

	// Create a waitgroup to wait for all runing goroutines
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		// Start each instance as a goroutine
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20));
	}

	// Print number of go routines (main + one per url
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine())

	wg.Wait()
}