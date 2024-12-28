package server

import (
	"log"
	"time"

	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
)

func SendVideoFrames(videoTrack *webrtc.TrackLocalStaticSample) {
	for {
		frame := []byte{}

		err := videoTrack.WriteSample(media.Sample{
			Data:     frame,
			Duration: time.Second / 30,
		})

		if err != nil {
			log.Println("Error writing video frame:", err)
			break
		}

		time.Sleep(time.Second / 30)
	}
}
