package main

import "github.com/imdevin567/jargon/cmd"

// Version ...
var Version string = "1.0.0"

func main() {
	if err := cmd.Execute(Version); err != nil {
		panic(err)
	}
}
