package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//一.自定义String/int类型的排序器
type Stringslice []string

func (s Stringslice) Len() int {
	return len(s)
}
func (s Stringslice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Stringslice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type SortIntSlice []int

func (s SortIntSlice) Len() int           { return len(s) }
func (s SortIntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s SortIntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for i < j {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

//二.实现对Track数组的自定义排序器
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type CustomSortTracks struct {
	tracks []*Track
	less   func(i, j *Track) bool
}

func (c CustomSortTracks) Len() int           { return len(c.tracks) }
func (c CustomSortTracks) Less(i, j int) bool { return c.less(c.tracks[i], c.tracks[j]) }
func (c CustomSortTracks) Swap(i, j int)      { c.tracks[i], c.tracks[j] = c.tracks[j], c.tracks[i] }

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

// IsPalindrome(s sort.Interface) bool
func main() {
	var nums []int = []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	fmt.Printf("%v\n", IsPalindrome(SortIntSlice(nums)))
}
