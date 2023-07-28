package main

import (
	"arena"
	"fmt"
	"testing"
)

// Type User is a struct with JSON struct tags to specify the JSON field names.
type TestUser struct {
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	TestCompany TestCompany `json:"company"`
}

type TestCompany struct {
	Name                string              `json:"name"`
	Address             string              `json:"address"`
	NumberOfEmployees   int                 `json:"number_of_employees"`
	TestFinancialReport TestFinancialReport `json:"financial_report"`
}

type TestFinancialReport struct {
	Revenue int `json:"revenue"`
	Profit  int `json:"profit"`
}

const numObjectsTesting = 1000000 // Number of objects to allocate and deallocate

// BenchmarkArenas tests the performance of allocating and marshaling User objects using arenas.
func BenchmarkArenas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Allocate arena
		a := arena.NewArena()
		// Allocate slice of TestUser objects
		users := arena.MakeSlice[TestUser](a, numObjectsTesting, numObjectsTesting)

		for i := 0; i < numObjectsTesting; i++ {
			users[i] = TestUser{
				FirstName: fmt.Sprintf("TestUser%d", i),
				LastName:  fmt.Sprintf("Lastname%d", i),
				Email:     fmt.Sprintf("user%d@example.com", i),
				Phone:     fmt.Sprintf("123456789%d", i),
				TestCompany: TestCompany{
					Name:    fmt.Sprintf("TestCompany%d", i),
					Address: fmt.Sprintf("Address%d", i),
					TestFinancialReport: TestFinancialReport{
						Revenue: 4 * i,
						Profit:  3 * i,
					},
				},
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
				TestCompany: TestCompany{
					Name:    fmt.Sprintf("TestCompany%d", i),
					Address: fmt.Sprintf("Address%d", i),
					TestFinancialReport: TestFinancialReport{
						Revenue: 4 * i,
						Profit:  3 * i,
					},
				},
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
