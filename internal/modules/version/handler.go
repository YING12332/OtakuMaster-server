package version

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetVersion(c *gin.Context) {
	platform := c.Query("platform")
	channel := c.Query("channel")
	curStr := c.Query("currentVersionCode")

	var cur int64
	if curStr != "" {
		if n, err := strconv.ParseInt(curStr, 10, 64); err == nil {
			cur = n
		}
	}

	v, err := h.svc.GetVersion(platform, channel, cur)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load version info"})
		return
	}

	// 版本接口不缓存（确保立刻生效）
	c.Header("Cache-Control", "no-cache")
	c.JSON(http.StatusOK, v)
}
