package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt") //txt file which contains all wordls
	if err != nil {
		log.Fatal(err)
	}
	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Println(str)
}
