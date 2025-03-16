package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstNames      string
	lasttName       string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	/*
		fmt.Println(remainingTickets)
		fmt.Println(&remainingTickets)

			fmt.Println("Welcome to", conferenceName, "booking application")
			fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are available")
			fmt.Println("Get you tickets here to attend")
	*/
	greetUser()

	//for {

	firstName, lasttName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lasttName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTickets(firstName, lasttName, email, userTickets)
		wg.Add(1)

		go sendTicket(userTickets, firstName, lasttName, email)

		//bookings[0] = firstName + " " + lasttName (useful when array is used)

		/*
			fmt.Printf("The whole Slice: %v\n", bookings)
			fmt.Printf("The first element: %v\n", bookings[0])
			fmt.Printf("The type of Slice: %T\n", bookings)
			fmt.Printf("The length of Slice: %v\n", len(bookings)) */

		// call firstNames
		firstNames := getFirstNames()
		fmt.Printf("The first names of our booking are %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("All the tickets are booked.Come back next year")
			//break

		}

	} else {
		if !isValidName {
			fmt.Println("The First nae or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("The email you entered does not contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}

	}
	wg.Wait()

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available.\n", conferenceTickets, remainingTickets)

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, data := range bookings {
		firstNames = append(firstNames, data.firstNames)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lasttName string
	var email string
	var userTickets uint
	// ask user for there name
	fmt.Println("Please enter your First Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your Last Name: ")
	fmt.Scan(&lasttName)

	fmt.Println("Please enter your Email: ")
	fmt.Scan(&email)

	fmt.Println("Please enter required no. of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lasttName, email, userTickets

}
func bookTickets(firstName string, lasttName string, email string, userTickets uint) {
	remainingTickets -= userTickets

	var userData = UserData{ // UserData is the name of struc and key is the name declared when intializing the struct but value must be variable declared in the function
		firstNames:      firstName,
		lasttName:       lasttName,
		email:           email,
		numberOfTickets: userTickets,
	}

	/*userData["firstName"] = firstName // storing the data using the key value pair
	userData["email"] = email
	userData["lasttName"] = lasttName
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) */

	bookings = append(bookings, userData)
	fmt.Printf("The list of bookign is %v \n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation mail on %v.\n", firstName, lasttName, userTickets, email)
	fmt.Printf(" %v tickets are remaining \n", remainingTickets)

}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
