package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthDate   string
	createdDate time.Time
}

func NewUser(fName string, lName string, bDate string) *User {
	created := time.Now()

	user := User{
		firstName:   fName,
		lastName:    lName,
		birthDate:   bDate,
		createdDate: created,
	}

	return &user
}

func (user *User) outputDetails() {
	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthDate)
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	var newUser *User

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	newUser = NewUser(firstName, lastName, birthdate)

	newUser.outputDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userIntput, _ := reader.ReadString('\n')

	cleanedInput := strings.Replace(userIntput, "\n", "", -1)

	return cleanedInput
}
