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
	gap := v.LatestBuild - v.CurrentBuild
	if gap <= 0 {
		return "none"
	}
	if gap > v.ForceUpgradeAfter {
		return "forced"
	}
	return "optional"
}
