package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Person struct {
	name     string
	age      int
	job_role string
}

var word string

func main() {

	// 1 - How many of each letter is present in the input word?

	letters := map[rune]int{}

	// Reading the word from keyboard
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

	fmt.Println("------------------------------------------------------")

	// 2 - Print the names and job roles of the people who are above 25 years old.

	//Creating a Slice of person to store all the people

	people := []Person{}

	p1 := Person{}
	p1.name = "Bender Gap"
	p1.age = 15
	p1.job_role = "Student"

	p2 := Person{}
	p2.name = "Bender Diff"
	p2.age = 28
	p2.job_role = "Engineer"

	p3 := Person{}
	p3.name = "Bender Supremacy"
	p3.age = 25
	p3.job_role = "Programmer"

	p4 := Person{}
	p4.name = "Bender Turbo omega griefer"
	p4.age = 34
	p4.job_role = "Loser"

	p5 := Person{}
	p5.name = "Bender Abyss"
	p5.age = 27
	p5.job_role = "Millionaire"

	p6 := Person{}
	p6.name = "Bender Disgusting supreme inter"
	p6.age = 50
	p6.job_role = "UltraLoser"

	// Adding all the people to the slice

	people = append(people, p1)
	people = append(people, p2)
	people = append(people, p3)
	people = append(people, p4)
	people = append(people, p5)
	people = append(people, p6)

	// Iterating through slice and checking each people's age

	for _, value := range people {
		if value.age > 25 {
			fmt.Println("Name: " + value.name + ". Job: " + value.job_role + ".")
		}

	}

	fmt.Println("------------------------------------------------------")

	// 3 - Remove the people from the array whose names are longer than 20 letters.

	var newPerson []Person

	// Iterating through slice and checking each people's name length

	for _, person := range people {
		if length := utf8.RuneCountInString(person.name); length <= 20 {
			newPerson = append(newPerson, person)
		}
	}

	// Printing out the new structure

	for _, value := range newPerson {
		fmt.Println("Name: " + value.name)
		fmt.Println("Age: ", value.age)
		fmt.Println("Job: " + value.job_role)
	}

	fmt.Println("------------------------------------------------------")

	// 4 - Change the name of each person to their first name.

	var firstNameOnly []Person

	// Iterating through slice and getting each people's first name only

	for _, value := range people {
		firstName := strings.Split(value.name, " ")
		value.name = firstName[0]
		firstNameOnly = append(firstNameOnly, value)
	}

	// Printing out the new structure

	for _, value := range firstNameOnly {
		fmt.Println("Name: " + value.name)
		fmt.Println("Age: ", value.age)
		fmt.Println("Job: " + value.job_role)
	}

	fmt.Println("------------------------------------------------------")

	// 5 - Add a new person through keyboard input.

	keyboardInput := []Person{}

	// Reading all the values from keyboard

	p7 := Person{}
	fmt.Print("Type the name: ")
	fmt.Scan(&p7.name)
	fmt.Print("Type the age: ")
	fmt.Scan(&p7.age)
	fmt.Print("Type the job: ")
	fmt.Scan(&p7.job_role)
	keyboardInput = append(keyboardInput, p7)

	// Printing out the final structure

	for _, value := range keyboardInput {
		fmt.Println("Name: " + value.name)
		fmt.Println("Age: ", value.age)
		fmt.Println("Job: " + value.job_role)

	}

}
