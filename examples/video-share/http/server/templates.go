package server

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/justincampbell/timeago"
)

func AddTemplates(r *gin.Engine, prefix string) {
	fm := template.FuncMap{
		"mod": func(i, j int) bool { return i%j == 0 },
		"ago": func(i int64) string {
			d, _ := time.ParseDuration(fmt.Sprintf("%ds", time.Now().Unix()-i))
			return timeago.FromDuration(d)
		},
		"tokens": func(s string, i int) string { return strings.Split(s, ".")[i] },
		"add":    func(i, j int) int { return i + j },
	}
	r.SetFuncMap(fm)
	r.LoadHTMLGlob(prefix + "templates/*.tmpl")
}
