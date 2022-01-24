// pass_fail reports whether a grade is passing or failing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            // 除去空格换行符
	grade, err := strconv.ParseFloat(input, 64) //字符串转化为浮点数

	if err != nil {
		log.Fatal(err)
	}

	var status string //块作用域
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
