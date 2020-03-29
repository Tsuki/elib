package main

import (
	"elib/internal/core"
	"elib/log"
	"elib/utils"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/xfrr/goffmpeg/transcoder"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "run jellyfin database",
		Long:  "run jellyfin database",
		RunE:  run,
	}
	quit = make(chan os.Signal, 1)
)

func run(_ *cobra.Command, _ []string) error {
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Info("running")
	var list []string
	tmp, err := core.WalkDirectory(utils.Option.InputPath, ".avi")
	log.CheckFatal(err)
	list = append(list, tmp...)

	tmp, err = core.WalkDirectory(utils.Option.InputPath, ".wmv")
	log.CheckFatal(err)
	list = append(list, tmp...)

	log.Debug(spew.Sdump(list))
	numJobs := len(list)
	log.Info("numJobs:", numJobs)
	jobs := make(chan *transcoder.Transcoder, numJobs)
	results := make(chan string, numJobs)
	for w := uint(1); w <= utils.Option.Worker; w++ {
		go core.Worker(int(w), jobs, results)
	}
	for _, i2 := range list {
		trans, err := core.Transcode(i2)
		if err != nil {
			return err
		}
		jobs <- trans
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		if path := <-results; path != "" {
			log.Info(path)
			if err := os.Rename(path, fmt.Sprintf("%s/%s", utils.Option.BackupPath, filepath.Base(path))); err != nil {
				log.Error(path, err)
			}
		}
	}
	return nil
}
