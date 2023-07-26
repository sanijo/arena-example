package main

import (
	"arena"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

const numObjectsTesting = 100000

// BenchmarkArenas tests the performance of allocating and marshaling User objects using arenas.
func BenchmarkArenas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := arena.NewArena()
		users := arena.MakeSlice[*User](a, 0, numObjectsTesting)

		for i := 0; i < numObjectsTesting; i++ {
			userObj := arena.New[User](a)
			*userObj = User{
				FirstName: fmt.Sprintf("User%d", i),
				LastName:  fmt.Sprintf("Lastname%d", i),
				Email:     fmt.Sprintf("user%d@example.com", i),
				Phone:     fmt.Sprintf("123456789%d", i),
			}
			users = append(users, userObj)
		}

		for _, user := range users {
			_, err := json.MarshalIndent(user, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
		}

		a.Free()
	}
}

// BenchmarkGarbageCollector tests the performance of allocating and marshaling User objects using the garbage collector.
func BenchmarkGarbageCollector(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var users []*User

		for i := 0; i < numObjectsTesting; i++ {
			userObj := &User{
				FirstName: fmt.Sprintf("User%d", i),
				LastName:  fmt.Sprintf("Lastname%d", i),
				Email:     fmt.Sprintf("user%d@example.com", i),
				Phone:     fmt.Sprintf("123456789%d", i),
			}
			users = append(users, userObj)
		}

		for _, user := range users {
			_, err := json.MarshalIndent(user, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// run_benchmarks runs the benchmarks.
func TestMain(m *testing.M) {
    fmt.Println("Running benchmarks...")
    fmt.Printf("Number of Objects: %d\n", numObjectsTesting)

    // Run benchmarks
    arenaTime := testing.Benchmark(BenchmarkArenas)
    gcTime := testing.Benchmark(BenchmarkGarbageCollector)

    // Convert and Print results in milliseconds
    fmt.Printf("Arenas Time: %.2f ms (%d iterations)\n", arenaTime.T.Seconds()*1000, arenaTime.N)
    fmt.Printf("Garbage Collector Time: %.2f ms (%d iterations)\n", gcTime.T.Seconds()*1000, gcTime.N)
}

