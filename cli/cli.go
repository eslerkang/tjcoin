package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/eslerkang/tjcoin/explorer"
	"github.com/eslerkang/tjcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to TJCoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:   Set the port of the server(default 4000)\n")
	fmt.Printf("-mode:   Choose between 'html' and 'rest'(default 'rest')\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
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
		explorer.Start(*port - 1000)
	default:
		usage()
	}
}
