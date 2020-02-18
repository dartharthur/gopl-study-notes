// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	resultsFilePath      = "exercises/1.10/results.txt"
	resultsFilePathPerms = 0644
)

func main() {
	start := time.Now()
	ch := make(chan string)
	f, err := os.OpenFile(resultsFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, resultsFilePathPerms)
	if err != nil {
		log.Fatal(err)
	}
	for _, url := range os.Args[1:] {
		go fetch(url, ch, f) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	timeSince := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Print(timeSince)
	if _, err := f.Write([]byte(timeSince + "\n")); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func fetch(url string, ch chan<- string, f *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	urlRead := fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
	ch <- urlRead
	if _, err := f.Write([]byte(urlRead + "\n")); err != nil {
		f.Close()
		log.Fatal(err)
	}
}
