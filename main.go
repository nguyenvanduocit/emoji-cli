package main

import (
	"github.com/nguyenvanduocit/emoji"
	"flag"
	"fmt"
)
func main() {
	var text string
	flag.StringVar(&text, "text", "I love cat", "Text used to find emoji")
	emojis, err := emoji.FindEmoji(text)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(emojis)
	}
}