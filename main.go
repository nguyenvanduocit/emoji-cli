package main

import (
	"github.com/nguyenvanduocit/emoji"
	"flag"
	"fmt"
)
func main() {
	var text string
	flag.StringVar(&text, "text", "I love cat", "Text used to find emoji")
	flag.Parse()
	emojis, err := emoji.FindEmoji(text)
	if err != nil {
		fmt.Println(err)
	}else{
		for _,emoji := range emojis{
			fmt.Println(emoji.Text, ":", emoji.Score)
		}
	}
}
