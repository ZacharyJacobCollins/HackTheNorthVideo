avconv -f video4linux2 -r 25 -i /dev/video0 -f alsa -i plughw:U0x46d0x825,0 -ar 22050 -ab 64k -strict experimental -acodec aac -vcodec mpeg4 -y ../../../webcam1.mp4
