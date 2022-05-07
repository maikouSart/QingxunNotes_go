package main

//ques 1
//猜谜游戏
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	//fmt.Println(secretNumber)

	for {
		fmt.Println("input your guess")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured")
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("error occured")
			continue
		}
		fmt.Println("your guess is", guess)

		if guess > secretNumber {
			fmt.Println("your guess is bigger")
		} else if guess < secretNumber {
			fmt.Println("your guess is smaller")
		} else {
			fmt.Println("bingo!")
			break
		}
	}
}
