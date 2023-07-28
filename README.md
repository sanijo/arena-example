# Go Arenas Exploration

### Info
This is a simple program for exploring Go arenas, based on the resources found [here](https://docs.go101.org/std/pkg/arena.html) and [here](https://docs.go101.org/std/src/arena/arena.go.html). 

The benchmark runs were configured with `benchtime=5m` and the number of objects tested (`numObjectsTesting`) was set to 1,000,000.

## Benchmark Results

The benchmark results show the performance of the two different memory management techniques: using Go's garbage collector and using arenas.

### BenchmarkGarbageCollector

This benchmark measures the performance of allocating and deallocating `TestUser` objects using Go's built-in garbage collector. Here are the results:

| Iterations | Time per operation (ms) | Memory per operation (MB) | Allocations per operation |
|------------|-------------------------|---------------------------|---------------------------|
| 405        | 966.71                  | 313.75                    | 12998515                  |

### BenchmarkArenas

This benchmark measures the performance of allocating and deallocating `TestUser` objects using the manual memory management technique with arenas. Here are the results:

| Iterations | Time per operation (ms) | Memory per operation (MB) | Allocations per operation |
|------------|-------------------------|---------------------------|---------------------------|
| 444        | 835.40                  | 266.23                    | 11998472                  |

### Relative Difference
Relative differences represent how much more (or less) the Arena values are compared to the GC values, expressed as a percentage. A negative value indicates that the Arena value is less than the GC value, while a positive value indicates it's more.
Relative difference is calculated as:
`Relative Difference (%) = ((Arena Value - GC Value) / GC Value) * 100`

|                            | GC          | Arena       | Relative Difference (%) |
|----------------------------|-------------|-------------|-------------------------|
| Iterations                 | 405         | 444         | +9.63                   |
| Time per operation (ms)    | 966.71      | 835.40      | -13.6                   |
| Memory per operation (MB)  | 313.75      | 266.23      | -15.07                  |
| Allocations per operation  | 12998515    | 11998472    | -7.69                   |


## Conclusion

The benchmarks suggest that, for this specific task, the arena approach was faster and used less memory per operation compared to Go's garbage collector. This indicates that the arena approach was more efficient in this specific scenario, being able to manage memory in a way that results in less memory usage and faster execution times.

The number of memory allocations in the arena case is slightly less than in the GC case. This could be due to the arena's strategy of pre-allocating a large block of memory and parceling it out as needed, which reduces the number of separate allocations it has to request from the operating system.

These results are specific to this particular benchmark and the type of operations it is performing. The relative performance of garbage collection vs arenas can vary depending on the specific use case, such as the size and lifecycle of the objects being allocated, and the patterns of memory usage in the program.

