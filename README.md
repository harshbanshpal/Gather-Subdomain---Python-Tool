## If you are unable to use subfinder tool after installation use this tool instead of it.

# Command

Here's a step-by-step guide to using the script:

Install Go by following the instructions on the Go website: https://golang.org/doc/install

Install subfinder by following the instructions in the subfinder repository: https://github.com/projectdiscovery/subfinder

Save the script to a file (e.g. subdomain_gather.go).

Build and run the script with the following commands:

go build subdomain_gather.go

./subdomain_gather example.com


Replace example.com with the domain for which you want to gather subdomains. The script will use the subfinder tool to gather subdomains and then save the results to a file called subdomains.txt.
