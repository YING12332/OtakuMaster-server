package version

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Repo interface {
	GetLatest(platform, channel string) (*VersionInfo, error)
}

type FileRepo struct {
	BaseDir string // e.g. /data/version
}

func (r *FileRepo) GetLatest(platform, channel string) (*VersionInfo, error) {
	if platform == "" {
		platform = "android"
	}
	if channel == "" {
		channel = "stable"
	}
	p := filepath.Join(r.BaseDir, platform, fmt.Sprintf("%s.json", channel))
	b, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	var v VersionInfo
	if err := json.Unmarshal(b, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
