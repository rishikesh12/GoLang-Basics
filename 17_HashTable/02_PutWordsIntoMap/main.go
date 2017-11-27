package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt") //txt file which contains all the words
	if err != nil {
		log.Fatal(err)
	}
	words := map[string]string{}

	scan := bufio.NewScanner(res.Body)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		words[scan.Text()] = ""
	}
	if err = scan.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	for k, _ := range words {
		fmt.Println(k)
	}
}
