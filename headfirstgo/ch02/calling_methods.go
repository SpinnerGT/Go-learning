package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// var now time.Time = time.Now()
	now := time.Now()
	//var year int = now.Year()
	year := now.Year()

	fmt.Println(year)
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
