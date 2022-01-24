// guess challenges players to guess a random number.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() //从1970年1月1日到现在的秒数(整数)
	rand.Seed(seconds)           //不同的随机数种子
	target := rand.Intn(100) + 1 //整数随机数
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)         //标准输入
	success := false                            //初始化
	for guesses := 0; guesses < 10; guesses++ { //允许10次重复
		if guesses > 0 { //第一次不提示,第二次及以后提示
			fmt.Println("You have", 10-guesses, "guesses left.")
		}
		// fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input) //除去换行符

		guess, err := strconv.Atoi(input) //str转化为int
		if err != nil {
			log.Fatal(err)
		}
		//判断大小
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
