package main

import (
	"fmt"
	"time"
)

func main() {
	go printLink("https://www.jetbrains.com/help/go/")
	go printLink("https://youtrack.jetbrains.com/issues/GO")
	go printLink("https://blog.jetbrains.com/")
	go printLink("https://intellij-support.jetbrains.com/hc/en-us/categories/200793669-General")
	go printLink("https://intellij-support.jetbrains.com/hc/en-us/community/topics")
	time.Sleep(2 * time.Second)
}

func printLink(link string) {
	fmt.Println("Printing:", link)
}
