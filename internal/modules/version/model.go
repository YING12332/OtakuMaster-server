package version

type VersionInfo struct {
	LatestVersionCode       int64  `json:"latestVersionCode"`
	LatestVersionName       string `json:"latestVersionName"`
	MinSupportedVersionCode int64  `json:"minSupportedVersionCode"`
	MinSupportedVersionName string `json:"minSupportedVersionName"`
	DownloadURL             string `json:"downloadUrl"`
	ReleaseNotes            string `json:"releaseNotes"`
	ForceUpdate             bool   `json:"forceUpdate"`
	ForceUpdateMessage      string `json:"forceUpdateMessage"`
	ChecksumSha256          string `json:"checksumSha256"`
	ApkSizeBytes            int64  `json:"apkSizeBytes"`
}
