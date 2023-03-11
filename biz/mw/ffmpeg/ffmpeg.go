package ffmpeg

import (
	"bytes"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GetSnapshot(videoPath string) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).
		Run()

	return buf, nil
}
