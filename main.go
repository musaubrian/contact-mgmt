package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//global variables
var check_existing string = checkExisting()
var cont_name string
var phone_num string
var street_name string
var choose_opt string
var empty string = ""
var line = "◼◼-----------------------------◼◼"

//error handling function
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// the input taking function
//usues bufio "package"
func getInput(ask string, r *bufio.Reader) (string, error) {
	fmt.Print(ask)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

//Title of the txt file
func userName() string {
	reader := bufio.NewReader(os.Stdin)
	username, _ := getInput("Your name: ", reader)
	return username
}

//gets contact name
func contactName() string {
	reader := bufio.NewReader(os.Stdin)
	contName, _ := getInput("Enter contact's name: ", reader)
	return contName
}

//gets contact number
func contactNumber() string {
	reader := bufio.NewReader(os.Stdin)
	contNum, _ := getInput("Enter phone number: ", reader)
	return contNum
}

//street name
func streetName() string {
	reader := bufio.NewReader(os.Stdin)
	strtName, _ := getInput("Street name: ", reader)
	return strtName
}

//checks if there is an existing file
func checkExisting() string {
	reader := bufio.NewReader(os.Stdin)
	check, _ := getInput("Create new contact-list / Override existing list / Add contacts (N / O / A) ", reader)
	return check
}

//option to continue
func option() string {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Would you like to continue? (y/n) ", reader)
	return opt
}

//Shows location of the saved txt file
func savedFile() {
	fmt.Println("File saved in contact-list.txt")
}

//creating the text file
func writeFile() {
	var head = userName()
	cont_name = contactName()
	phone_num = contactNumber()
	street_name = streetName()

	//writing to a file
	file, err := os.Create("./contact-list.txt") //directory to create file
	checkError(err)

	write := bufio.NewWriter(file)

	// actual writing of the formatted strings
	_, err = fmt.Fprint(write, head, "'s contact list ☎", "\n")
	checkError(err)

	_, err = fmt.Fprint(write, line, "\n")
	checkError(err)

	_, err = fmt.Fprint(write, empty, "\n")
	checkError(err)

	_, err = fmt.Fprint(write, "Contact name: ", cont_name, "\n")
	checkError(err)

	_, err = fmt.Fprint(write, "phone number: ", phone_num, "\n")
	checkError(err)

	_, err = fmt.Fprint(write, "Street name: ", street_name, "\n")
	checkError(err)

	write.Flush()

	choose_opt = option()

	if choose_opt == "y" || choose_opt == "yes" {
		addContact()
	} else {
		savedFile()
	}

}

//adding more onto existing
func addContact() {
	var cont_name = contactName()
	var phone_num = contactNumber()
	var street_name = streetName()
	var choose_opt string

	file, err := os.OpenFile("./contact-list.txt", os.O_APPEND|os.O_WRONLY, 0644) //opens the file that was written
	checkError(err)

	write := bufio.NewWriter(file)

	defer file.Close()

	_, err = fmt.Fprint(write, line, "\n")
	checkError(err)

	_, err = fmt.Fprint(write, empty, "\n")
	checkError(err)

	_, err = fmt.Fprint(write, "Contact name: ", cont_name, "\n")
	checkError(err)

	_, err = fmt.Fprint(write, "Phone number: ", phone_num, "\n")
	checkError(err)

	_, err = fmt.Fprint(write, "Street name: ", street_name, "\n")
	checkError(err)

	write.Flush()

	choose_opt = option()

	if choose_opt == "y" || choose_opt == "yes" {
		addContact()
	} else {
		savedFile()
	}

}

func main() {

	if check_existing == "N" || check_existing == "new" || check_existing == "n" {
		userName()
		writeFile()

	} else if check_existing == "O" || check_existing == "o" || check_existing == "override" {
		writeFile()

	} else if check_existing == "A" || check_existing == "a" || check_existing == "add" {
		addContact()
	}

}
