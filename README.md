### Race Tracks with Hoppers

The solution is using breadth-first search (BFS) to look for the shortest path in unweighted graphs.

Time Complexity: O(X * Y) where X and Y are the grid dimensions.

Space Complexity: O(X * Y).

### Run the solution

1. Use `go run main.go` to run the program
2. Type the following input:
```   
2
5 5
4 0 4 4
1
1 4 2 3
3 3
0 0 2 2
2
1 1 0 2
0 2 1 1
```

3. Use `go test -v ./...` to run all tests
