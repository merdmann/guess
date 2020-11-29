package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	valueToGuess := int(rand.Intn(10))

	MyGuess := utils.promptInt("My guess :" + strconv.FormatInt(int64(valueToGuess), 10))
	for MyGuess != valueToGuess {

		if MyGuess < valueToGuess {
			fmt.Println("to low")
		} else {
			if MyGuess > valueToGuess {
				fmt.Println("to high")
			}
		}
		MyGuess = utils.promptInt("My guess:")
	}
}
