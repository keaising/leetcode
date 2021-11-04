package main

import (
	"log"
	"strings"
)

func main() {
	result := wordsTyping([]string{
		"hello",
		"world",
	}, 2, 8)
	log.Println(result)
}

func wordsTyping(sentence []string, rows int, cols int) int {
    var s = strings.Join(sentence, " ") + " "
    var m = make(map[int]int)
    m[0] = 0
    for i := 1; i < len(s); i++ {
        if s[i] == ' ' {
            m[i] = 1
        } else {
            m[i] = m[i-1] - 1
        }
    }
    var start = 0
    for i := 0; i < rows; i++ {
         start += cols
         start += m[start%len(s)]
    }
    return start / len(s)
}
