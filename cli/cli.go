package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/eslerkang/tjcoin/explorer"
	"github.com/eslerkang/tjcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to TJCoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:   Set the port of the server(default 4000)\n")
	fmt.Printf("-mode:   Choose between 'html' and 'rest'(default 'rest')\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 3000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		//start html explorer
		explorer.Start(*port)
	case "both":
		go rest.Start(*port)
		explorer.Start(*port+1000)
	default:
		usage()
	}
}