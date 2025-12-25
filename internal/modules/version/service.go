package version

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetVersion(platform, channel string, currentVersionCode int64) (*VersionInfo, error) {
	v, err := s.repo.GetLatest(platform, channel)
	if err != nil {
		return nil, err
	}

	// 可选：后端根据 currentVersionCode 自动强更（预留）
	if currentVersionCode > 0 && currentVersionCode < v.MinSupportedVersionCode {
		v.ForceUpdate = true
		if v.ForceUpdateMessage == "" {
			v.ForceUpdateMessage = "该版本已停止服务，请立即更新"
		}
	}
	return v, nil
}
