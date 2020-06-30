package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tjarratt/babble"
	"jjaa.me/models"
	"jjaa.me/persist"
)

var babbler = babble.NewBabbler()
var Db *sqlx.DB

func main() {
	fmt.Println("worker")
	rand.Seed(time.Now().UnixNano())
	Db = persist.Connection()
	babbler.Count = 4
	name := babbler.Babble()
	for {
		video, _ := models.SelectVideoForWorker(Db, name)
		fmt.Println("got", video)

		DownloadFromBucket(video.UrlSafeName + "." + video.Ext)

		ExtractJpg(video)

		models.ClearVideoForWorker(Db, name)
	}
}

func ExtractJpg(video *models.Video) {
	exec.Command("ffmpeg", "-ss", "00:00:03", "-i",
		video.UrlSafeName+"."+video.Ext,
		"-vframes", "1", "-q:v", "2",
		video.UrlSafeName+".jpg").Run()
	models.UpdateVideo(Db, "jpg_ready", video.Ext, video.UrlSafeName)
}
