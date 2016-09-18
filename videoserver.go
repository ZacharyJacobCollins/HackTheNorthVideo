package main

import (
    "net/http"
    "time"
    "io/ioutil"
    "fmt"
    "os/exec"
)

func main() {
    //cams := scanCams()
    //audioDevices := scanAudio()
    //fmt.Println(cams)
    //fmt.Println(audioDevices) 
    //startWebcam() 
    startWebserver()      

}

/**
*   TODO: make this database driven
*   Retrieves available webcams attached to network.
*   @return [array of cam names as strings]
*/
func scanCams() []string { 
	//Declare active cameras here 
	cams := []string{"/dev/video0"}
	return cams
}

func scanAudio() []string {
	//Declare 
	audioDevices := []string{"U0x46d0x825"}
 	return audioDevices	
}


//Takes the name of the webcam and starts recording inside of a go routine.
func startWebcam() {
	cmd := exec.Command("avconv", "-f", "video4linux2", "-r", "25", "-i", "/dev/video0", "-f", "alsa", "-i", "plughw:U0x46d0x825,0", "-ar", "22050", "-ab", "64k", "-strict", "experimental", "-acodec", "aac", "-vcodec", "mpeg4", "-y", "webcam1.mp4", "/dev/video0")
	out, err := cmd.Output()
	check(err)
	fmt.Print(string(out))
	
}

func startWebserver() {
  //Start the webserver
	fs := http.FileServer(http.Dir("../../../../../videos"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", nil)
}

//Get the current utc timestamp, write time to text file
func writeTimestamp() {
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
