package domain

type VersionTarget struct {
	ProductSlug       string `json:"product_slug"`
	TargetType        string `json:"target_type"`
	CurrentVersion    string `json:"current_version"`
	CurrentBuild      int    `json:"current_build"`
	LatestVersion     string `json:"latest_version"`
	LatestBuild       int    `json:"latest_build"`
	ForceUpgradeAfter int    `json:"force_upgrade_after"`
}

func (v VersionTarget) UpgradeMode() string {
	return v.UpgradeModeForBuild(v.CurrentBuild)
}

func (v VersionTarget) UpgradeModeForBuild(currentBuild int) string {
	gap := v.LatestBuild - currentBuild
	if gap <= 0 {
		return "none"
	}
	if gap > v.ForceUpgradeAfter {
		return "forced"
	}
	return "optional"
}

type CheckUpgradeRequest struct {
	ProductSlug    string `json:"product_slug"`
	CurrentVersion string `json:"current_version"`
	CurrentBuild   int    `json:"current_build"`
}
