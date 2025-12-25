package http

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type DownloadHandler struct {
	ApkDir string
}

func NewDownloadHandler(apkDir string) *DownloadHandler {
	return &DownloadHandler{ApkDir: apkDir}
}

func (h *DownloadHandler) DownloadApk(c *gin.Context) {
	filename := c.Param("file")
	filename = strings.TrimPrefix(filename, "/")

	// 简单安全：禁止目录穿越
	if strings.Contains(filename, "..") || strings.ContainsAny(filename, `\/:`) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	full := filepath.Join(h.ApkDir, filename)
	f, err := os.Open(full)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	defer f.Close()

	st, err := f.Stat()
	if err != nil || st.IsDir() {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// 头部：apk 下载
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	c.Header("Cache-Control", "public, max-age=31536000, immutable")
	// Content-Type
	ct := mime.TypeByExtension(filepath.Ext(filename))
	if ct == "" {
		ct = "application/vnd.android.package-archive"
	}
	c.Header("Content-Type", ct)

	// 原生支持 Range
	http.ServeContent(c.Writer, c.Request, filename, st.ModTime().Truncate(time.Second), f)
}
