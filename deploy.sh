#!/bin/bash

# ==========================================
# OtakuMaster-server 部署/更新脚本
# ==========================================

# 1. 停止当前运行的容器（如果是第一次部署，此步骤可能会报错，已忽略）
echo ">>> 正在停止旧服务..."
docker-compose stop api
docker-compose rm -f api

# 2. 重新构建镜像 (强制不使用缓存，确保包含最新代码)
echo ">>> 正在重新构建镜像..."
docker-compose build --no-cache api

# 3. 启动服务 (后台运行)
echo ">>> 正在启动新服务..."
docker-compose up -d api

# 4. 清理未使用的镜像 (可选，防止磁盘占用过多)
echo ">>> 清理悬空镜像..."
docker image prune -f

echo "✅ 部署完成！"
echo "你可以使用 'docker logs -f otaku-api' 查看日志。"
