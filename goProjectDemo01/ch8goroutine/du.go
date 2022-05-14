package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var v = flag.Bool("v", false, "whether show the detail")
var aboard chan struct{} = make(chan struct{})

/* var tokens chan struct{} = make(chan struct{}, 20) */
var wg sync.WaitGroup

func canceled() bool {
	select {
	case <-aboard:
		return true
	default:
		return false
	}
}

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	var nfiles, nbytes int64
	var tick <-chan time.Time
	fileSizes := make(chan int64)
	if *v {
		tick = time.Tick(1 * time.Second)
	}
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(aboard)
	}()

	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()
	/* 	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}() */

	// Print the results.
LOOP:
	for {
		select {
		case <-aboard:
			for range fileSizes {
			}
			break LOOP
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		case size, ok := <-fileSizes:
			if !ok {
				break LOOP
			}
			nfiles++
			nbytes += size
		}
	}
	printDiskUsage(nfiles, nbytes)

}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

//!-main

//!+walkDir

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	defer wg.Done()
	if canceled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	/* 	tokens <- struct{}{} */
	entries, err := ioutil.ReadDir(dir)
	/* 	<-tokens */
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

//!-walkDir

// The du1 variant uses two goroutines and
// prints the total after every file is found.
