package http

import (
	"github.com/YING12332/OtakuMaster-server/internal/modules/version"
	"github.com/gin-gonic/gin"
)

type RouterDeps struct {
	VersionHandler  *version.Handler
	DownloadHandler *DownloadHandler
}

func RegisterRoutes(r *gin.Engine, deps RouterDeps) {
	api := r.Group("/api/v1")
	{
		api.GET("/version", deps.VersionHandler.GetVersion)
	}

	// 下载接口保持顶层（便于 CDN/302/反代）
	r.GET("/download/:file", deps.DownloadHandler.DownloadApk)
	r.HEAD("/download/:file", deps.DownloadHandler.DownloadApk)
}
