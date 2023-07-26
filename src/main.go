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
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

const (
	numObjects   = 100000 // Number of objects to allocate and deallocate
	arenaEnabled = true   // Set to true to use arenas, false to use GC
)

func main() {
//    var users []*User
	// Measure the execution time
	startTime := time.Now()

	if arenaEnabled {
        fmt.Println("arenaEnabled case with numObjects = ", numObjects)
		a := arena.NewArena() // Create a new arena
		defer a.Free() // Free all memory allocated by the arena after it goes out of scope.

	    // Slice to store pointers to User objects using arena memory
        users := arena.MakeSlice[*User](a, 0, numObjects) // Make a slice of length 0 and capacity numObjects

		// Allocate memory for all User objects from the arena and add pointers to the slice
		for i := 0; i < numObjects; i++ {
			userObj := arena.New[User](a) // Allocate memory from the arena
            // Set the fields of the object through the pointer.
			*userObj = User{
				FirstName: fmt.Sprintf("User%d", i),
				LastName:  fmt.Sprintf("Lastname%d", i),
				Email:     fmt.Sprintf("user%d@example.com", i),
				Phone:     fmt.Sprintf("123456789%d", i),
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
//    // Marshal each User object to JSON (to simulate some processing)
//    for _, user := range users {
//        _, err := json.Marshal(user)
//        if err != nil {
//            log.Fatal(err)
//        }
//    }

	// Measure the execution time
	elapsedTime := time.Since(startTime)

	fmt.Printf("Time taken: %s\n", elapsedTime)
}

