param(
  [Parameter(Mandatory=$true)]
  [string]$ApkPath,                     # 例如: .\data\apks\OtakuMaster_1.0.1.apk

  [Parameter(Mandatory=$true)]
  [int64]$LatestVersionCode,            # 例如: 10001

  [Parameter(Mandatory=$true)]
  [string]$LatestVersionName,           # 例如: 1.0.1

  [Parameter(Mandatory=$true)]
  [int64]$MinSupportedVersionCode,      # 例如: 10001

  [Parameter(Mandatory=$true)]
  [string]$MinSupportedVersionName,     # 例如: 1.0.1

  [Parameter(Mandatory=$true)]
  [string]$DownloadUrl,                 # 例如: http://YOUR_SERVER_IP:8080/download/OtakuMaster_1.0.1.apk

  [Parameter(Mandatory=$true)]
  [string]$ReleaseNotes,                # 例如: 修复闪退；新增系列详情页

  [Parameter(Mandatory=$true)]
  [bool]$ForceUpdate,                   # 例如: $true

  [Parameter(Mandatory=$true)]
  [string]$ForceUpdateMessage,          # 例如: 该版本已停止服务，请立即更新

  [string]$StableJsonPath = ".\data\version\android\stable.json"
)

if (!(Test-Path $ApkPath)) {
  Write-Host "APK not found: $ApkPath" -ForegroundColor Red
  exit 1
}

# 计算 sha256
$hash = (Get-FileHash -Algorithm SHA256 $ApkPath).Hash.ToLower()

# 文件大小 bytes
$size = (Get-Item $ApkPath).Length

# 生成 JSON
$obj = [ordered]@{
  latestVersionCode       = $LatestVersionCode
  latestVersionName       = $LatestVersionName
  minSupportedVersionCode = $MinSupportedVersionCode
  minSupportedVersionName = $MinSupportedVersionName
  downloadUrl             = $DownloadUrl
  releaseNotes            = $ReleaseNotes
  forceUpdate             = $ForceUpdate
  forceUpdateMessage      = $ForceUpdateMessage
  checksumSha256          = $hash
  apkSizeBytes            = $size
}

# 确保目录存在
$dir = Split-Path -Parent $StableJsonPath
if (!(Test-Path $dir)) { New-Item -ItemType Directory -Path $dir | Out-Null }

# 写入 pretty JSON
$json = ($obj | ConvertTo-Json -Depth 10)
Set-Content -Path $StableJsonPath -Value $json -Encoding UTF8

Write-Host "Updated: $StableJsonPath" -ForegroundColor Green
Write-Host "sha256: $hash"
Write-Host "size:   $size bytes"
