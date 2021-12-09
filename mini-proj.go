package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	name := scanner.Text()

	if name == "" {
		fmt.Println("Name cannot be empty")
		os.Exit(1)
	} else {
		fmt.Println("Your name is:", name)
	}

	fmt.Println("We have gotten your name")

	fmt.Println("Enter your age")
	scanner_2 := bufio.NewScanner(os.Stdin)
	scanner_2.Scan()

	age, _ := strconv.ParseInt(scanner_2.Text(), 10, 64)

	if 18 > age {
		fmt.Println("Sorry you are not up to age")
		fmt.Println("Thank you for trying out our program")

		os.Exit(1)
	} else {
		fmt.Println("You passed the age test!!")
	}

	fmt.Println("Thank you for trying out our program you have successfully completed all steps!!")
	fmt.Println("Your name is: ", name, " and you are ", age, " years old")

}
