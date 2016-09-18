package main

import (
    "net/http"
    "time"
    "io/ioutil"
)

func main() {
	//Get the current utc timestamp
	now := time.Now().UTC().Format("20060102150405")

	//write time to text file 
	timestamp := []byte(now)
	err := ioutil.WriteFile("./timestamp.txt", timestamp, 0644)
	if (err != nil) {
		panic(err)
	}


	
	//Start the webserver
	fs := http.FileServer(http.Dir("../../../../../videos"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", nil)
}

