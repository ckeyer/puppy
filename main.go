package main

import (
	_ "github.com/ckeyer/commons/debug"
	_ "github.com/ckeyer/commons/version"
	"github.com/ckeyer/puppy/cmd"
)

func main() {
	cmd.Execute()
}
