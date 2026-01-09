package http

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
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

	// 处理 latest 别名
	if filename == "latest" {
		latest, err := h.findLatestApk()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No latest version found"})
			return
		}
		filename = latest
	}

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

func (h *DownloadHandler) findLatestApk() (string, error) {
	entries, err := os.ReadDir(h.ApkDir)
	if err != nil {
		return "", err
	}

	var apks []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".apk") {
			apks = append(apks, e.Name())
		}
	}

	if len(apks) == 0 {
		return "", fmt.Errorf("no apks found")
	}

	// 按版本号降序排序 (v1 > v2)
	sort.Slice(apks, func(i, j int) bool {
		v1 := extractVersion(apks[i])
		v2 := extractVersion(apks[j])
		return compareVersions(v1, v2)
	})

	return apks[0], nil
}

func extractVersion(name string) []int {
	// 匹配 OtakuMaster_1.0.1.apk 中的 1.0.1
	// 兼容其他命名，只要包含数字版本即可
	re := regexp.MustCompile(`_([\d.]+)\.apk$`)
	matches := re.FindStringSubmatch(name)
	if len(matches) < 2 {
		return []int{0}
	}
	parts := strings.Split(matches[1], ".")
	var v []int
	for _, p := range parts {
		i, _ := strconv.Atoi(p)
		v = append(v, i)
	}
	return v
}

// compareVersions returns true if v1 > v2
func compareVersions(v1, v2 []int) bool {
	l := len(v1)
	if len(v2) < l {
		l = len(v2)
	}
	for i := 0; i < l; i++ {
		if v1[i] != v2[i] {
			return v1[i] > v2[i]
		}
	}
	// 前缀相同，长度长的版本号更大 (e.g. 1.0.1 > 1.0)
	return len(v1) > len(v2)
}
