package main


import (
	"bufio"
	"fmt"
	"os"
	"library-management-system/controllers"
)
func welcome() int {
	var choice int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Here are the operations you can perform:")
	fmt.Println("1. Add a new book")
	fmt.Println("2. Remove an existing book")
	fmt.Println("3. Borrow a book")
	fmt.Println("4. Return a book")
	fmt.Println("5. List all available books")
	fmt.Println("6. List all borrowed books by a member")
	fmt.Print("Please select the number of the operation you want to perform:")

	input, _ := reader.ReadString('\n')
	n, err := fmt.Sscanf(input, "%d", &choice)

	if choice < 1 || choice > 6 || n != 1 || err != nil {
		fmt.Println("Invalid input. Please Try Again")
		fmt.Println()
		return welcome()
	} else {
		fmt.Println("You selected:", choice)
		return choice
	}

}

func main() {
	fmt.Println("Welcome to the Library Management System!")
	choice := welcome()
	fmt.Println(controllers.Controller(choice))

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Print("Let's do another operation? Enter 1 for ok, any other key to exit: ")
		input, _ := reader.ReadString('\n')

		var again int
		n, err := fmt.Sscanf(input, "%d", &again)
		if n != 1 || err != nil || again != 1 {
			fmt.Println("--Thank you--")
			break
		}
		choice := welcome()
		fmt.Println(controllers.Controller(choice))
	}
}