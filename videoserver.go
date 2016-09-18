package main

import (
    "net/http"
    "time"
    "io/ioutil"
    "fmt"
    "os/exec"
)

func main() {
    scanCams()

}

/**
*   Retrieves available webcams attached to network.
*   @return [array of cam names as strings]
*/
func scanCams() {
  cmd := exec.Command("ls", "/dev/video*")
  out, err := cmd.Output()
  check(err)
  fmt.Print(string(out))
}

/**
*   Retrieves available audio devices attached to network.
*   @return [array of audio devices as strings]
*/
func scanAudio() {

}

//Takes the name of the webcam and starts recording inside of a go routine.
func startWebcam() {

}

func startWebserver() {
  //Start the webserver
	fs := http.FileServer(http.Dir("../../../../../videos"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", nil)
}

//Get the current utc timestamp, write time to text file
func heartBeat() {
  now := time.Now().UTC().Format("20060102150405")
  timestamp := []byte(now)
  err := ioutil.WriteFile("./timestamp.txt", timestamp, 0644)
  check(err)
}

//Print error if err
func check(e error) {
	if(e != nil) {
		fmt.Println(e.Error())
	}
}
