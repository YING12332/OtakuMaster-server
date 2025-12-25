package main

import (
	"fmt"

	"github.com/YING12332/OtakuMaster-server/internal/config"
	httpx "github.com/YING12332/OtakuMaster-server/internal/http"
	"github.com/YING12332/OtakuMaster-server/internal/modules/version"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 组装版本模块
	versionRepo := &version.FileRepo{BaseDir: config.Cfg.VersionDataDir}
	versionSvc := version.NewService(versionRepo)
	versionHandler := version.NewHandler(versionSvc)

	downloadHandler := httpx.NewDownloadHandler(config.Cfg.ApkDir)

	httpx.RegisterRoutes(r, httpx.RouterDeps{
		VersionHandler:  versionHandler,
		DownloadHandler: downloadHandler,
	})

	addr := fmt.Sprintf(":%s", config.Cfg.HTTPPort)
	_ = r.Run(addr)
}
