# Go Conference Booking Application

**Overview**

This is a simple Go-based ticket booking system for a conference. It allows users to book tickets, stores their information using a struct, and sends ticket confirmations asynchronously using goroutines.

**Features**

1. Takes user input for booking tickets.

2. Validates user input (name, email, ticket count).

3. Stores bookings using a slice of structs.

4. Sends ticket confirmation asynchronously using goroutines and sync.WaitGroup.

5. Displays the first names of booked users.

**Technologies Used**

Go (Golang)

Concurrency (goroutines, sync.WaitGroup)# booking-app
