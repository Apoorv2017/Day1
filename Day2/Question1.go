package main

import (
	"fmt"
	"strings"
)

func fre(word string, ch chan map[string]int) {
	temp := map[string]int{}
	char := strings.Split(word, "")
	for _, i := range char {
		temp[string(i)]++
	}
	ch <- temp
}
func main() {
	word := []string{"quick", "qbrown", "fox", "lazy", "dog"}
	ch := make(chan map[string]int)
	freq_map := map[string]int{}
	for _, i := range word {
		go fre(i, ch)
		temp := <-ch
		for key, val := range temp {
			freq_map[key] += val
		}
	}
	fmt.Println(freq_map)
}
