package server

import (
	"log"
	"time"

	"github.com/BeauRussell/GoVideoPlayer/videoEncoding"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
)

func SendVideoFrames(videoTrack *webrtc.TrackLocalStaticSample) {
	frames := videoEncoding.SplitVideoToFrames()
	for _, frame := range frames {
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
