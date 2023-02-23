package ffmpeg

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
)

func GetSnapshot(videoPath string, imageName string) (imagePath string, err error) {
	err = ffmpeg.Input(videoPath).
		Output(imageName, ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		Run()
	if err != nil {
		return "", err
	}
	pwd, err := os.Getwd()
	return pwd + "/" + imageName, nil
}
