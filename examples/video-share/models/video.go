package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Video struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Comments    int    `json:"comments"`
	Status      string `json:"status"`
	UrlSafeName string `json:"url_safe_name"`
	Ext         string `json:"ext"`
	CreatedAt   int64  `json:"created_at"`
}

const VIDEO_SELECT = "SELECT id, status, ext, url_safe_name as urlsafename, title, comments, UNIX_TIMESTAMP(created_at) as createdat from videos"

func ClearVideoForWorker(db *sqlx.DB, name string) string {
	_, err := db.NamedExec("UPDATE videos set worker=null where worker=:worker",
		map[string]interface{}{"worker": name})
	if err != nil {
		return err.Error()
	}
	return ""
}
func SelectVideoForWorker(db *sqlx.DB, name string) (*Video, string) {
	_, err := db.NamedExec("UPDATE videos set worker=:worker where worker is null and status='uploaded' limit 1",
		map[string]interface{}{"worker": name})
	item := Video{}
	sql := fmt.Sprintf("%s where worker=:name", VIDEO_SELECT)
	rows, err := db.NamedQuery(sql, map[string]interface{}{"name": name})
	if err != nil {
		return &item, err.Error()
	} else {
		if rows.Next() {
			rows.StructScan(&item)
		} else {
			return &item, "not_found"
		}
	}

	return &item, ""
}
func SelectVideo(db *sqlx.DB, name string) (*Video, string) {
	item := Video{}
	sql := fmt.Sprintf("%s where url_safe_name=:name", VIDEO_SELECT)
	rows, err := db.NamedQuery(sql, map[string]interface{}{"name": name})
	if err != nil {
		return &item, err.Error()
	} else {
		if rows.Next() {
			rows.StructScan(&item)
		} else {
			return &item, "not_found"
		}
	}

	return &item, ""
}
func SelectVideos(db *sqlx.DB, userId int) ([]Video, string) {
	videos := []Video{}
	sql := fmt.Sprintf("%s where user_id = %d order by created_at desc limit 1000", VIDEO_SELECT, userId)
	if userId == 0 {
		sql = fmt.Sprintf("%s order by created_at desc limit 1000", VIDEO_SELECT)
	}
	err := db.Select(&videos, sql)
	s := ""
	if err != nil {
		s = err.Error()
	}

	return videos, s
}
func InsertVideo(db *sqlx.DB, title, safeName string, id int) string {
	_, err := db.NamedExec("INSERT INTO videos (title, url_safe_name, user_id, status) values (:title, :safe_name, :id, :status)",
		map[string]interface{}{"title": title, "id": id,
			"safe_name": safeName,
			"status":    "not_uploaded"})
	if err != nil {
		return err.Error()
	}
	return ""
}
func UpdateVideo(db *sqlx.DB, status, ext, name string) string {
	_, err := db.NamedExec("UPDATE videos set ext=:ext,status=:status where url_safe_name=:name",
		map[string]interface{}{"status": status, "name": name, "ext": ext})
	if err != nil {
		return err.Error()
	}
	return ""
}
