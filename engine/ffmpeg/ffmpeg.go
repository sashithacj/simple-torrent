package ffmpeg

import (
	"fmt"
	"log"
	"sync"

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
	var mutex sync.RWMutex
	// Start transcoder process without checking progress
	go func() {
		mutex.RLock()
		done := trans.Run(true)
		fmt.Print(done)
		progress := trans.Output()
		for msg := range progress {
			fmt.Println(msg)
		}
		mutex.RUnlock()
	}()

}
