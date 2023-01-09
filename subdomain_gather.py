import requests
import re

def gather_subdomains(url):
  subdomains = []
  # Make a request to the URL
  try:
    response = requests.get(url)
  except requests.ConnectionError:
    print("Failed to connect to the website")
    return []
  # Find all subdomains in the response
  pattern = r"(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]"
  subdomains = re.findall(pattern, response.text)
  return subdomains

# Prompt the user for a URL
url = input("Enter a URL to gather subdomains from: ")

# Gather the subdomains
subdomains = gather_subdomains(url)

# Print the results
print("Subdomains:", subdomains)
