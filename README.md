[![Build Status](https://github.com/harshadnawathe/leetcode/actions/workflows/go-build.yml/badge.svg)](https://github.com/harshadnawathe/leetcode/actions/workflows/go-build.yml)

# leetcode 
Repository to practice `leetcode` problems

# Project Structure
The entire project contains multiple packages. Each package is created based on the problem title. E.g., a problem titled as `Two Sum`
is placed in a package called `two-sum`.

# Package Structure
Each package contains -
+ README file
+ Golang source code
+ Golang test code

# Generating Package
Default solution package for a leetcode problem can be generated using the [leetcode.go](./leetcode.go) utility.

Run the utility with `go run` and provide url of the leetcode problem as the arguments.

```sh
go run leetcode.go https://leetcode.com/problems/two-sum/
```

Upon successful execution it will generate a solution package with following files.
+ `README.md` file with problem title and link to the actual statement.
+ `solution.go` file with default code from leetcode problem statement.
+ `solution_test.go` file with basic test structure for the function in the `solution.go` but includes no testcases. 

# Running tests
+ Running all tests `go test ./...`
+ Running a specific package tests `go test ./two-sum
