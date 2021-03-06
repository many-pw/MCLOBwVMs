package controllers

import (
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"jjaa.me/models"
)

func VideosNew(c *gin.Context) {
	if !BeforeAll("user", c) {
		return
	}
	c.HTML(http.StatusOK, "videos__new.tmpl", gin.H{
		"flash": flash,
		"user":  user,
	})
}
func VideosAllIndex(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	videos, _ := models.SelectVideos(Db, 0)
	c.HTML(http.StatusOK, "videos__all_index.tmpl", gin.H{
		"videos": videos,
		"user":   user,
		"flash":  flash,
	})
}
func VideosIndex(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	videos, _ := models.SelectVideos(Db, user.Id)
	c.HTML(http.StatusOK, "videos__index.tmpl", gin.H{
		"videos": videos,
		"user":   user,
		"flash":  flash,
	})
}
func VideosShow(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	video, _ := models.SelectVideo(Db, c.Param("name"))
	c.HTML(http.StatusOK, "videos__show.tmpl", gin.H{
		"video": video,
		"user":  user,
		"flash": flash,
	})
}
func VideosUpload(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	video, _ := models.SelectVideo(Db, c.Param("name"))
	c.HTML(http.StatusOK, "videos__upload.tmpl", gin.H{
		"video": video,
		"flash": "",
		"user":  user,
	})

}
func VideosCreate(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	title := strings.TrimSpace(c.PostForm("title"))
	if title == "" {
		SetFlash("title needed", c)
		c.Redirect(http.StatusFound, "/videos/new")
		c.Abort()
		return
	}

	babbler.Count = 4
	words := babbler.Babble()
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	safeName := reg.ReplaceAllString(strings.ToLower(words), "-")
	models.InsertVideo(Db, title, safeName, user.Id)
	models.IncrementUserCount(Db, "videos", user.Id)
	c.Redirect(http.StatusFound, "/videos/upload/"+safeName)
	c.Abort()
}
func VideosDestroy(c *gin.Context) {

	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
func VideosFile(c *gin.Context) {
	BeforeAll("", c)
	fileHeader, _ := c.FormFile("file")
	go func(fileHeader *multipart.FileHeader, name string) {
		video, _ := models.SelectVideo(Db, name)
		tokens := strings.Split(fileHeader.Filename, ".")
		ext := tokens[len(tokens)-1]
		fileWithExt := video.UrlSafeName + "." + ext
		f, _ := fileHeader.Open()
		UploadToBucket(fileWithExt, f)
		f.Close()
		models.UpdateVideo(Db, "uploaded", ext, video.UrlSafeName)
	}(fileHeader, c.Param("name"))
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
