package main

import (
	"github.com/TypicalAM/goread/cmd/goread"
)

var version = "v1.6.2"

func main() {
	goread.SetVersion(version)
	goread.Execute()
}
