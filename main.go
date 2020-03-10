package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

type Url struct {
	visited bool
}

var tags []string

func returnTags(url string) []string  {
	resp, _ := http.Get(url)
	tokenizer := html.NewTokenizer(resp.Body) // returns a new HTML Tokenizer(stream of html tokens) for the given Reader
	for {
		tt := tokenizer.Next() // scans the next token and returns it's type

		if tt == html.ErrorToken {
			return nil
		}

		if tt == html.StartTagToken {
			t := tokenizer.Token() // returns the current token(data and other attributes)
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href"  {
						fmt.Println("Found href:", a.Val)
						tags = append(tags, a.Val)
						break
					}
				}
			}
		}

		if tt == html.EndTagToken {
			t := tokenizer.Token()

			if t.Data == "html" {
				return tags
			}
		}

	}

}

func traceMap(tags []string){
	str := fmt.Sprintf("https://schier.co/blog/indie-to-acquisition")
	returnTags(str)
	//for _, link := range tags {
	//
	//}
}

//func siteMap(url string){
//	resp, _ := http.Get(url)
//
//	tags := returnTags(resp.Body)
//
//	fmt.Println(tags)
//
//	resp.Body.Close()
//}

func main(){
	returnTags("https://schier.co")
	traceMap(tags)
	fmt.Println(tags)
}
