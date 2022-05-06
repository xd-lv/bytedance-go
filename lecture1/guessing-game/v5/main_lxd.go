package main

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
	maxNumber := 100
	rand.Seed(time.Now().UnixNano())
	secretAns := rand.Intn(maxNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		guess, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		if guess < secretAns {
			fmt.Println("smaller than ans")
		} else if guess > secretAns {
			fmt.Println("bigger than ans")
		} else {
			fmt.Println("bingo!")
			break
		}
	}

}
