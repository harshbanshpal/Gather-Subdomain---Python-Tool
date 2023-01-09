package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func gatherSubdomains(domain string) []string {
	// Run the subfinder tool
	cmd := exec.Command("subfinder", "-d", domain)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running subfinder:", err)
		return []string{}
	}

	// Split the output by line and return the results
	subdomains := strings.Split(string(output), "\n")
	return subdomains[:len(subdomains)-1]
}

func main() {
	// Check if a domain was provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Error: No domain provided")
		return
	}
	domain := os.Args[1]

	// Gather the subdomains
	subdomains := gatherSubdomains(domain)

	// Save the subdomains to a file
	file, err := os.Create("subdomains.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, subdomain := range subdomains {
		fmt.Fprintln(writer, subdomain)
	}
	writer.Flush()

	fmt.Println("Subdomains saved to file: subdomains.txt")
}
