package main

import "github.com/jakoblorz/metrikxd/cmd"

//go:generate go generate ./...

func main() {
	cmd.Execute()
}
