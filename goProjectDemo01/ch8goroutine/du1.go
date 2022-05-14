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

var wg sync.WaitGroup
var count map[string]int64 = make(map[string]int64)
var rwlock sync.Mutex
var aboard chan struct{} = make(chan struct{}{})
func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var tick *time.Ticker
	if *v {
		tick = time.NewTicker(1 * time.Second)
	}
	go func(){
		os.Stdin.Read(make([]byte, 1))
		close(aboard)
	}()
	go func() {
		for {
			select {
			case _, ok := <-tick.C:
				if !ok {
					return
				}
				for key, value := range count {
					fmt.Printf("%s:%.1f GB\t", key, float64(value)/1e9)
				}
				fmt.Printf("\n")
			}
		}
	}()
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, root)
	}
	wg.Wait()
	tick.Stop()
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

//!-main

//!+walkDir

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, name string) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			walkDir(subdir, name)
		} else {
			count[name] += entry.Size()
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
