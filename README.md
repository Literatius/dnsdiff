# DNSDiff

A lightweight command-line DNS change tracker written in Go.

DNSDiff resolves domain IPv4 addresses, compares them with previously saved records, and reports changes.

## Features

- DNS lookup for IPv4 addresses
- Detection of added, removed, and unchanged IP addresses
- Local JSON storage
- Support for tracking multiple domains
- Error handling
- Unit tests
- No external dependencies

## Installation

**Requirements:** Go installed on your system.

Clone the repository:

`git clone https://github.com/Literatius/dnsdiff.git`

`cd dnsdiff`

Build the application:

`go build -o dnsdiff .`

## Usage

Run the compiled executable:

`./dnsdiff example.com`

## Testing

Run unit tests:

`go test -v ./...`

Check code coverage:

`go test -cover ./...`
