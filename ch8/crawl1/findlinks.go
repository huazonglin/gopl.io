// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 240.

// Crawl1 crawls web links starting with the command-line arguments.
//
// This version quickly exhausts available file descriptors
// due to excessive concurrent calls to links.Extract.
//
// Also, it never terminates because the worklist is never closed.
package main

import (
	"fmt"
	"log"
	"os"
	"time"
	//"gopl.io/ch5/links"
)

//!+crawl
func crawl(url string) []string {
	fmt.Println("\n Received URL for crawl = " + url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("\n crawl list == ", list, "\n crawl URL =  ", url+"\n\n\n")
	return list
}

//!-crawl

//!+main
func main() {
	fmt.Println("Start .........")
	worklist := make(chan []string)

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	fmt.Println("Point 1 .........")
	fmt.Println("worklist == ", worklist)
	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		fmt.Println("\n\nPoint 2 ====================================================================================")
		fmt.Println("\n TOP list = ", list)
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

// links.go 是从ch5 copy 过来的，用下面命令运行。
// go run *.go http://gopl.io

//!-main

/*
//!+output
$ go build gopl.io/ch8/crawl1
$ ./crawl1 http://gopl.io/
http://gopl.io/
https://golang.org/help/

https://golang.org/doc/
https://golang.org/blog/
...
2015/07/15 18:22:12 Get ...: dial tcp: lookup blog.golang.org: no such host
2015/07/15 18:22:12 Get ...: dial tcp 23.21.222.120:443: socket:
                                                        too many open files
...
//!-output
*/
