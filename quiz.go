package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to my PC components quiz")

	fmt.Printf("What is your name? ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, this is a PC components quiz. \nAre you ready? ", name)

	var ready string
	fmt.Scan(&ready)

	if ready == "yes" {
		fmt.Println("Awesome lets get started!!")
	}

	if ready == "no" {
		fmt.Println("No problem, come back when you are ready :)")
		time.Sleep(10 * time.Second)
		os.Exit(0)
	}

	if ready != "yes" {
		fmt.Println("That is not a option.")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}

	score := 0
	num_questions := 6

	fmt.Println("What is the name of the component that is known as the brain of the computer?")
	var answer string
	fmt.Scan(&answer)

	if answer == "cpu" {
		fmt.Println("Correct!")
		score++
	}

	if answer != "cpu" {
		fmt.Println("Incorrect")
	}

	fmt.Println("How do we measure Processor Speed?")
	var answer2 string
	fmt.Scan(&answer2)

	if answer2 == "gigahertz" {
		fmt.Println("Correct!")
		score++
	}

	if answer2 != "gigahertz" {
		fmt.Println("Incorrect")
	}

	fmt.Println("True or False: Solid State Hard Drives are faster than Hard Disk Drives?")
	var answer3 string
	fmt.Scan(&answer3)

	if answer3 == "true" {
		fmt.Println("Correct!")
		score++
	}

	if answer3 != "true" {
		fmt.Println("Incorrect")
	}

	fmt.Println("True or False: All computer components need to connect to the Motherboard?")
	var answer4 string
	fmt.Scan(&answer4)

	if answer4 == "true" {
		fmt.Println("Correct!")
		score++
	}

	if answer4 != "true" {
		fmt.Println("Incorrect")
	}

	fmt.Println("What gives a computer the power needed to turn on?")
	var answer5 string
	fmt.Scan(&answer5)

	if answer5 == "psu" {
		fmt.Println("Correct!")
		score++
	}

	if answer5 != "psu" {
		fmt.Println("Incorrect")
	}

	fmt.Println("What is a physical component of a computer called?")
	var answer6 string
	fmt.Scan(&answer6)

	if answer6 == "hardware" {
		fmt.Println("Correct!")
		score++
	}

	if answer6 != "hardware" {
		fmt.Println("Incorrect")
	}

	fmt.Printf("Final score was %v out of %v\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You Scored: %v%%.", percent)
}
