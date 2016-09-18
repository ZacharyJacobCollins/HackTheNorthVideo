package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
	currentTime := time.Now().UTC.Format("2006-05-05 15:00:00")

	fmt.Println(currentTime);

	//Start the webserver
	fs := http.FileServer(http.Dir("../../../../../videos"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", nil)
}
