package main

import (
	"fmt"

	"net/http"

	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {

	//Variables
	var gate = "Topasvegen"
	fmt.Println(gate)

	var husNummer = "5"
	fmt.Println(husNummer)

	// Make a GET request to the API endpoint
	resp, err := http.Get("https://api2.haugnett.com/hkraft/?key=7ddc6abc54f51e584b783f713a797f60&gatenavn=" + gate + "%20" + husNummer)
	if err != nil {
		// Handle error
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle error
		return
	}

	// Print the response body
	fmt.Println(string(body))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
