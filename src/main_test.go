package main

import (
	"arena"
	"fmt"
	"testing"
)

// Type User is a struct with JSON struct tags to specify the JSON field names.
type TestUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

const numObjectsTesting = 100000 // Number of objects to allocate and deallocate

// BenchmarkArenas tests the performance of allocating and marshaling User objects using arenas.
func BenchmarkArenas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := arena.NewArena()
		//                users := arena.MakeSlice[*TestUser](a, numObjectsTesting, 2*numObjectsTesting)
		//
		//                for i := 0; i < numObjectsTesting; i++ {
		//                        userObj := arena.New[TestUser](a)
		//                        *userObj = TestUser{
		//                                FirstName: fmt.Sprintf("TestUser%d", i),
		//                                LastName:  fmt.Sprintf("Lastname%d", i),
		//                                Email:     fmt.Sprintf("user%d@example.com", i),
		//                                Phone:     fmt.Sprintf("123456789%d", i),
		//                        }
		//                        users = append(users, userObj)
		//                }
		users := arena.MakeSlice[TestUser](a, numObjectsTesting, numObjectsTesting)

		for i := 0; i < numObjectsTesting; i++ {
			users[i] = TestUser{
				FirstName: fmt.Sprintf("TestUser%d", i),
				LastName:  fmt.Sprintf("Lastname%d", i),
				Email:     fmt.Sprintf("user%d@example.com", i),
				Phone:     fmt.Sprintf("123456789%d", i),
			}
		}

		// Perform operations on objects
		for _, user := range users {
			user.FirstName = "NewName"
			user.Email = "new@example.com"
		}

		a.Free() // free the arena after each iteration
	}
}

// BenchmarkGarbageCollector tests the performance of allocating and marshaling User objects using the garbage collector.
func BenchmarkGarbageCollector(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var users []*TestUser

		for i := 0; i < numObjectsTesting; i++ {
			userObj := &TestUser{
				FirstName: fmt.Sprintf("User%d", i),
				LastName:  fmt.Sprintf("Lastname%d", i),
				Email:     fmt.Sprintf("user%d@example.com", i),
				Phone:     fmt.Sprintf("123456789%d", i),
			}
			users = append(users, userObj)
		}

		// Perform operations on objects
		for _, user := range users {
			user.FirstName = "NewName"
			user.Email = "new@example.com"
		}

	}
}
