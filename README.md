# keccak256_brute
A brute force attempt on Keccak256 precompile to get a defined output

This project contains a Go script that performs a brute-force attack using the Keccak256 hashing algorithm. The script attempts to find a hash value that matches a predefined target by altering input data. It utilizes goroutines for parallel processing to optimize the search process.

### Overview

The script combines fixed strings with a variable segment (c), hashes the concatenated result using the Keccak256 algorithm (as used in Ethereum), and checks if the output matches the desired hash output. If a match is found, the script prints the successful input and terminates.

### Prerequisites / Dependencies

- Go (version 1.11 or higher) must be installed on your machine.
- github.com/ethereum/go-ethereum package is required for the Keccak256 hashing.


### Installation

Step 1: Install Go
If Go is not already installed on your system, download and install it from Go's <a href="https://go.dev/dl/" target="_blanck" > official website. </a>

Step 2: Clone the Repository
Clone this repository to your local machine using

Step 3: Install Dependencies
Run the following command to install the necessary Go dependencies:
```go get github.com/ethereum/go-ethereum```

Step 4: Build the Project
Compile the project to ensure everything is set up correctly:
```go build```

Step 5: Usage
Run the compiled program using:
```./bruteforce```

This will initiate the brute-force process across multiple goroutines. The script will output the value of c if a matching hash is found, otherwise, it will report that no solution was found.

### Configuration

You can adjust the number of goroutines and the target hash directly in the code. These settings are found at the beginning of the main function and within the bruteForce function definition.

### Contributing

Contributions to this project are welcome. You can contribute by improving the efficiency of the brute-force algorithm, adding configuration options, or optimizing memory management. Please create a pull request with your proposed changes.