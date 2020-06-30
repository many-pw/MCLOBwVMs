package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"

	"github.com/tjarratt/babble"
	"jjaa.me/models"
	"jjaa.me/persist"
)

var babbler = babble.NewBabbler()

func main() {
	fmt.Println("worker")
	rand.Seed(time.Now().UnixNano())
	Db := persist.Connection()
	babbler.Count = 4
	name := babbler.Babble()
	for {
		video, _ := models.SelectVideoForWorker(Db, name)
		fmt.Println("got", video)
		filename := video.UrlSafeName

		DownloadFromBucket(filename + "." + video.Ext)

		exec.Command("ffmpeg", "-ss", "00:00:03", "-i",
			filename+"."+video.Ext,
			"-vframes", "1", "-q:v", "2",
			filename+".jpg").Run()
		models.UpdateVideo(Db, "jpg_ready", video.Ext, filename)

		models.ClearVideoForWorker(Db, name)
	}
}
