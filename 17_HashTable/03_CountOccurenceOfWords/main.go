package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//get the book
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10.txt")
	if err != nil {
		log.Fatal(err)
	}
	//scan the response
	scan := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	//Set scanner for split operation
	scan.Split(bufio.ScanWords)
	//Slice to hold counts
	bucket := make([]int, 200)
	for scan.Scan() {
		n := HashBucket(scan.Text(), 12)
		bucket[n]++
	}
	fmt.Println(bucket[:12])
}
func HashBucket(word string, bucket int) int {
	letter := int(word[0])
	buckets := letter % bucket
	return buckets
}
