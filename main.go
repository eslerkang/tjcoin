package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eslerkang/tjcoin/utils"
)

const port string =":4000"

type URLDescription struct {
	URL string
	Method string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{
			URL: "/",
			Method: "GET",
			Description: "See Documentation",
		},
	}
	b, err := json.Marshal(data)
	utils.HandleError(err)
	fmt.Printf("%s\n", b)
	fmt.Fprintf(rw, "%s\n", b)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))	
}