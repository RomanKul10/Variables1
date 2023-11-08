package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don`t type your number in, just press ENTER when ready"

func main() {

	/* // one way - declare, then assign (two steps)
	   // 1 варіант -  оголосити, потім призначити
	   var firstNumber int
	   firstNumber = 2

	   // another way, declare type and name and assign value
	   //інший варіант, оголосити тип і назву та призначити значення
	   var secondNumber = 5

	   // one step variable: declare name, assign value, and let Go figure out the type
	   // скорочений варіант: оголосити назву, призначити значення та дозволити Go визначити тип
	   subtraction := 7
	*/

	// seed the random number generator (передати рамдомне число)
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	// Intn(8) - ціле число від 0 (по замовчуванню) до 8

	var secondNumber = rand.Intn(8) + 2

	var subtraction = rand.Intn(8) + 2

	var answer int

	reader := bufio.NewReader(os.Stdin)

	// display a welcome/ instruction
	fmt.Println("Guess the Number Game")
	fmt.Println("-----------------------")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the game
	// проведіть їх через гру
	fmt.Println("Multiply your numberby", firstNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')
	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')
	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')
	// give them the answer
	answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is", answer)

}
