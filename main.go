package main

import (
	"os"

	"github.com/wtchangdm/is-mercury-in-retrograde/cmd/api"
	"github.com/wtchangdm/is-mercury-in-retrograde/cmd/client"
)

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "--server" || arg == "-S" {
			api.Serve()
		}
	}

	client.Do()
}
