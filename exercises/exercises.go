package main

import "fmt"

type person struct {
	name     string
	age      int
	job_role string
}

var word string

func main() {

	// 1 - How many of each letter is present in the input word?

	letters := map[rune]int{}

	fmt.Println("Type the word: ")
	fmt.Scanln(&word)

	for _, letter := range word { // Iterate over each letter of the word
		if _, present := letters[letter]; present { // If the letter is present in the map, increment its quantity by 1
			letters[letter] += 1
		} else { // Otherwise create it and make it quantity equals 1
			letters[letter] = 1
		}
	}

	fmt.Println("Printing the letters!")

	for letter, count := range letters { // Getting the key-value pair out of the for loop
		fmt.Printf("Letter: %c, Count: %d\n", letter, count)
	}

	// 2 - Print the names and job roles of the people who are above 25 years old.
	// > need to know how to read the values from structures from keyboard

	// 3 - Remove the people from the array whose names are longer than 26 letters.

	// 4 - Change the name of each person to their first name.

	// 5 - Add a new person through keyboard input.
}
