package main

import (
	"fmt"
	"strings"
	"time"

	//"io/ioutil"
	//"log"
	"net/http"
	"os"
	"runtime"
)

func checkUrlChannel(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		// Send err info to the channel
		fmt.Printf("Error: %v, Server at URL: %s is down\n", err, url)
		ch <- url

	} else {
		defer resp.Body.Close()

		data := fmt.Sprintf("URL %s Status code: %d \n", url, resp.StatusCode)
		data += fmt.Sprintf("%s is UP\n", url)

		ch <- data
	}
}

func main() {

	// Infinely run in the background and check the server
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
		go checkUrlChannel(url, ch)
	}

	// Print number of go routines (main + one per url
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine())

	fmt.Println("Press CTRL-C to stop the process");

	// Goroutine call version 1
	/*
	for {
		go checkUrlChannel(<-ch, ch)
		fmt.Printf(strings.Repeat("#", 20))

		// .... This would be better with a scheduled event like:   time.After(), time.Ticker() etc...
		time.Sleep(time.Second * 2)
	}
	 */

	// Goroutine call version 2
	/*
	for url := range ch {
		// .... This would be better with a scheduled event like:   time.After(), time.Ticker() etc...
		time.Sleep(time.Second * 2)

		go checkUrlChannel(url, ch)
		fmt.Printf(strings.Repeat("#", 20))
	}
	 */

	// Goroutine call version 3
	for url := range ch {
		// Passing the string here avoid race condition that may occur should be pass a pointer
		// of the string data instead
		go func (u string) {
			// .... This would be better with a scheduled event like:   time.After(), time.Ticker() etc...
			time.Sleep(time.Second * 2)

			go checkUrlChannel(u, ch)
			fmt.Printf(strings.Repeat("#", 20))
		} (url)
	}
}