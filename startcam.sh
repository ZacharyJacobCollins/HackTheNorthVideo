#!/bin/bash

Timeout=30 # 30 minutes

function timeout_monitor() {
   sleep "$Timeout"
   kill "$1"
}

# start the timeout monitor in 
# background and pass the PID:
timeout_monitor "$$" &
Timeout_monitor_pid=$!

avconv -f video4linux2 -r 25 -i /dev/video0 -f alsa -i plughw:U0x46d0x825,0 -ar 22050 -ab 64k -strict experimental -acodec aac -vcodec mpeg4 -y ../../../../../videos/webcam1.mp4

kill "$Timeout_monitor_pid"
