package main

import "fmt"

func main() {
	var dept string
	var year int

	fmt.Print("Enter your department name and year: ")
	fmt.Scanf("%s %d", &dept, &year)

	fmt.Println("Dept:", dept, "Year:", year)
	fmt.Printf("Dept: %s, Year: %d", dept, year)

}
