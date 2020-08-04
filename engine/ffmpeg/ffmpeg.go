package ffmpeg

import (
	"fmt"
	"github.com/xfrr/goffmpeg/ffmpeg"
	"github.com/xfrr/goffmpeg/models"
	"github.com/xfrr/goffmpeg/transcoder"
	"log"
	"path"
	"sync"
)

var mm = make(map[string]<-chan models.Progress)
var mutex sync.RWMutex

func Tomp4(inputPath, outputPath string) {
	log.Println("input-output:", inputPath, outputPath)
	mutex.RLock()
	defer mutex.RUnlock()
	if t := mm[inputPath]; t != nil {
		log.Println("is doing ", inputPath)
		return
	}

	outputPath = outputPath + ".mp4"
	trans := new(transcoder.Transcoder)

	fpath := ffmpeg.Configuration{FfprobeBin: "E:/worktool/ffmpeg/bin/ffprobe.exe ", FfmpegBin: "E:/worktool/ffmpeg/bin/ffmpeg.exe "}
	trans.SetConfiguration(fpath)
	err := trans.Initialize(inputPath, outputPath)
	log.Println("err:", err)
	trans.MediaFile().SetResolution("320x240")
	trans.MediaFile().SetVideoBitRate("400k")
	trans.MediaFile().SetFrameRate(25)

	go func() {
		done := trans.Run(true)
		fmt.Print(done)
		progress := trans.Output()
		mm[inputPath] = progress
		for msg := range progress {
			fmt.Println(msg)
		}
		delete(mm, inputPath)
	}()

}

func ListProgress() map[string]float64 {
	res := make(map[string]float64)
	for transs := range mm {
		out := mm[transs]
		num := <-out
		filenameWithSuffix := path.Base(transs) + ".mp4"
		fmt.Println("get file", filenameWithSuffix)
		res[filenameWithSuffix] = num.Progress
	}
	return res
}
