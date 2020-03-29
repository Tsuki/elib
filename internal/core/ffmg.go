package core

import (
	"elib/log"
	"github.com/pkg/errors"
	"github.com/xfrr/goffmpeg/ffmpeg"
	"github.com/xfrr/goffmpeg/transcoder"
	"path"
	"path/filepath"
)

func Transcode(input string) (*transcoder.Transcoder, error) {
	// Create new instance of transcoder
	trans := new(transcoder.Transcoder)
	var cfg ffmpeg.Configuration
	cfg.FfmpegBin = "C:\\Users\\feeli\\Project\\elib\\ffmpeg.exe"
	cfg.FfprobeBin = "C:\\Users\\feeli\\Project\\elib\\ffprobe.exe"
	trans.SetConfiguration(cfg)
	// Initialize transcoder passing the input file path and output file path
	ext := path.Ext(input)
	outputPath := input[0:len(input)-len(ext)] + ".mp4"
	if err := trans.Initialize(input, outputPath); err != nil {
		return nil, errors.Errorf("can't init ffmpeg: %v", err)
	}

	//enable nvenc
	//trans.MediaFile().SetVideoCodec("h264_nvenc")

	// Handle error...

	// Start transcoder process without checking progress
	//done := trans.Run(true)
	// Returns a channel to get the transcoding progress
	//progress := trans.Output()
	//logger := log.NewContextLogger(filepath.Base(input))
	//// Example of printing transcoding progress
	//for msg := range progress {
	//	if int(msg.Progress)%5 == 0 {
	//		logger.Infof("%+v", msg)
	//	}
	//}
	//
	//// This channel is used to wait for the process to end
	//if err := <-done; err != nil {
	//	return nil, errors.Errorf("ffmpeg run err: %s %w", input, err)
	//}
	//logger.Info("finish")
	return trans, nil
}

func Worker(id int, jobs <-chan *transcoder.Transcoder, results chan<- string) {
	for trans := range jobs {
		base := filepath.Base(trans.MediaFile().InputPath())
		logger := log.NewContextLogger(base)
		logger.Info("worker", id, "started  job ", trans.MediaFile().InputPath())
		if id <= 2 {
			trans.MediaFile().SetVideoCodec("h264_nvenc")
		}
		done := trans.Run(true)
		msg := trans.Output()
		Progress := 0
		for i := range msg {
			if p := int(i.Progress); p%5 == 0 && p != Progress {
				Progress = p
				logger.Info("Progress: ", Progress, "%")
			}
		}
		if err := <-done; err != nil {
			logger.Error("err", err)
		}
		results <- trans.MediaFile().InputPath()
	}
}

func RunTransCode(trans *transcoder.Transcoder) error {
	done := trans.Run(true)
	return <-done
}
