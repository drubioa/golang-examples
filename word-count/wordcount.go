package main

import (
	"golang.org/x/tour/wc" //Debe lanzarse en golang.org
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words:= strings.Fields(s)
	for a := range words  {
		var count int
		for b := range words  {
			if(words[a] == words[b]) {
				count = count + 1
			}
		}
		m[words[a]] = count
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
