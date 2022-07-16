package main

import "github.com/jakoblorz/autofone/cmd"

//go:generate go generate ./...

func main() {
	cmd.Execute()
}
