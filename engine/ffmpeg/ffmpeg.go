package ffmpeg

import (
	_ "errors"
	"fmt"
	"log"

	"github.com/xfrr/goffmpeg/transcoder"
)

type Ff struct {
	input  string
	output string
}

func Tomp4(inputPath, outputPath string) {
	outputPath = outputPath + ".mp4"
	log.Println("input-output:", inputPath, outputPath)
	// Create new instance of transcoder
	trans := new(transcoder.Transcoder)
	//f := ffmpeg.Configuration{FfprobeBin: "E:/worktool/ffmpeg/bin/ffprobe.exe ", FfmpegBin: "E:/worktool/ffmpeg/bin/ffmpeg.exe "}
	//trans.SetConfiguration(f)
	err := trans.Initialize(inputPath, outputPath)
	log.Println("err:", err)
	//trans.MediaFile().SetVideoCodec("xvid")
	trans.MediaFile().SetResolution("480x320")
	trans.MediaFile().SetVideoBitRate("400k")
	trans.MediaFile().SetFrameRate(25)

	// Handle error...

	// Start transcoder process without checking progress
	done := trans.Run(true)
	fmt.Print(done)
	// This channel is used to wait for the process to end
	progress := trans.Output()

	// Example of printing transcoding progress
	for msg := range progress {
		fmt.Println(msg)
	}

	err = <-done

}
