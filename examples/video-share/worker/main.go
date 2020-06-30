package main

import (
	"fmt"
	"math/rand"
	"os"
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
		ConvertToMp4(video)
		ConvertToWebm(video)
		ConvertToM4a(video)
		ConvertToOga(video)

		os.Remove("orig_" + video.UrlSafeName + "." + video.Ext)
		DeleteFileFromBucket(video.UrlSafeName + "." + video.Ext)
		models.ClearVideoForWorker(Db, name)
	}
}

func ExtractJpg(video *models.Video) {
	exec.Command("ffmpeg", "-ss", "00:00:03", "-i",
		"orig_"+video.UrlSafeName+"."+video.Ext,
		"-vframes", "1", "-q:v", "2",
		video.UrlSafeName+".jpg").Run()

	go func() {
		UploadToPublicBucket(video.UrlSafeName + ".jpg")
		models.UpdateVideo(Db, "jpg_ready", video.Ext, video.UrlSafeName)
	}()
}
func ConvertToMp4(video *models.Video) {
	exec.Command("ffmpeg", "-i",
		"orig_"+video.UrlSafeName+"."+video.Ext,
		"-vcodec", "h264", "-acodec", "aac",
		video.UrlSafeName+".mp4").Run()
	go func() {
		UploadToPublicBucket(video.UrlSafeName + ".mp4")
		models.UpdateVideo(Db, "mp4_ready", video.Ext, video.UrlSafeName)
	}()
}
func ConvertToWebm(video *models.Video) {
	exec.Command("ffmpeg", "-i",
		"orig_"+video.UrlSafeName+"."+video.Ext,
		video.UrlSafeName+".webm").Run()
	go func() {
		UploadToPublicBucket(video.UrlSafeName + ".webm")
		models.UpdateVideo(Db, "webm_ready", video.Ext, video.UrlSafeName)
	}()
}
func ConvertToM4a(video *models.Video) {
	exec.Command("ffmpeg", "-i",
		"orig_"+video.UrlSafeName+"."+video.Ext,
		video.UrlSafeName+".m4a").Run()
	go func() {
		UploadToPublicBucket(video.UrlSafeName + ".m4a")
		models.UpdateVideo(Db, "m4a_ready", video.Ext, video.UrlSafeName)
	}()
}
func ConvertToOga(video *models.Video) {
	exec.Command("ffmpeg", "-i",
		"orig_"+video.UrlSafeName+"."+video.Ext,
		video.UrlSafeName+".oga").Run()
	go func() {
		UploadToPublicBucket(video.UrlSafeName + ".oga")
		models.UpdateVideo(Db, "live", video.Ext, video.UrlSafeName)
	}()
}
