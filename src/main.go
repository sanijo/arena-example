package main

import (
	"arena"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Type User is a struct with JSON struct tags to specify the JSON field names.
type User struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Company   Company `json:"company"`
}

type Company struct {
	Name              string          `json:"name"`
	Address           string          `json:"address"`
	NumberOfEmployees int             `json:"number_of_employees"`
	FinancialReport   FinancialReport `json:"financial_report"`
}

type FinancialReport struct {
	Revenue int `json:"revenue"`
	Profit  int `json:"profit"`
}

const (
	numObjects   = 100000 // Number of objects to allocate and deallocate
	arenaEnabled = true   // Set to true to use arenas, false to use GC
)

func main() {
	// Measure the execution time
	startTime := time.Now()

	if arenaEnabled {
		fmt.Println("arenaEnabled case with numObjects = ", numObjects)
		a := arena.NewArena() // Create a new arena

		// Slice to store User objects using arena memory
		users := arena.MakeSlice[User](a, numObjects, numObjects)

		// Add objects to the slice
		for i := 0; i < numObjects; i++ {
			users[i] = User{
				FirstName: fmt.Sprintf("TestUser%d", i),
				LastName:  fmt.Sprintf("Lastname%d", i),
				Email:     fmt.Sprintf("user%d@example.com", i),
				Phone:     fmt.Sprintf("123456789%d", i),
				Company: Company{
					Name:              fmt.Sprintf("Company%d", i),
					Address:           fmt.Sprintf("Address%d", i),
					NumberOfEmployees: i,
					FinancialReport: FinancialReport{
						Revenue: 4 * i,
						Profit:  3 * i,
					},
				},
			}
		}

		// Marshal each User object to JSON (to simulate some processing)
		for _, user := range users {
			_, err := json.Marshal(user)
			if err != nil {
				log.Fatal(err)
			}
		}

		a.Free() // free the arena
	} else {
		fmt.Println("arenaDisabled case with numObjects = ", numObjects)
		// Slice to store pointers to User objects
		var users []*User
		// Allocate memory using the garbage collector and add pointers to the slice
		for i := 0; i < numObjects; i++ {
			userObj := &User{
				FirstName: fmt.Sprintf("User%d", i),
				LastName:  fmt.Sprintf("Lastname%d", i),
				Email:     fmt.Sprintf("user%d@example.com", i),
				Phone:     fmt.Sprintf("123456789%d", i),
				Company: Company{
					Name:              fmt.Sprintf("Company%d", i),
					Address:           fmt.Sprintf("Address%d", i),
					NumberOfEmployees: i,
					FinancialReport: FinancialReport{
						Revenue: 4 * i,
						Profit:  3 * i,
					},
				},
			}
			users = append(users, userObj)
		}

		// Marshal each User object to JSON (to simulate some processing)
		for _, user := range users {
			_, err := json.Marshal(user)
			if err != nil {
				log.Fatal(err)
			}
		}

	}

	// Measure the execution time
	elapsedTime := time.Since(startTime)

	fmt.Printf("Time taken: %s\n", elapsedTime)
}
