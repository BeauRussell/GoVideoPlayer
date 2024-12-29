package videoEncoding

import (
	"bytes"
	"io"
	"log"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func SplitVideoToFrames() [][]byte {
	testFilePath := "./testVideos/1920x1080_world_spin.mp4"

	var buf bytes.Buffer
	err := ffmpeg.
		Input(testFilePath).
		Output("pipe:", ffmpeg.KwArgs{
			"format":  "rawvideo",
			"pix_fmt": "rgb24",
		}).
		WithOutput(&buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatalf("Error extracting frames: %v", err)
	}

	frameWidth := 1920
	frameHeight := 1080
	frameSize := frameWidth * frameHeight * 3

	frames := [][]byte{}
	for {
		frame := make([]byte, frameSize)
		_, err := buf.Read(frame)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading frame: %v", err)
		}
		frames = append(frames, frame)
	}

	return frames
}
