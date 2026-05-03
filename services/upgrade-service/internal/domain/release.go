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

type ReleaseVersion struct {
	ID           string `json:"id"`
	ProductSlug  string `json:"product_slug"`
	TargetType   string `json:"target_type"`
	VersionLabel string `json:"version_label"`
	BuildNumber  int    `json:"build_number"`
	UpgradeURL   string `json:"upgrade_url,omitempty"`
}

type CreateReleaseRequest struct {
	ProductSlug  string `json:"product_slug"`
	TargetType   string `json:"target_type"`
	VersionLabel string `json:"version_label"`
	BuildNumber  int    `json:"build_number"`
	UpgradeURL   string `json:"upgrade_url"`
}

type DeploymentRecord struct {
	ID              string `json:"id"`
	TargetVersionID string `json:"target_version_id"`
	Environment     string `json:"environment"`
	Status          string `json:"status"`
}

type CreateDeploymentRequest struct {
	TargetVersionID string `json:"target_version_id"`
	Environment     string `json:"environment"`
	Status          string `json:"status"`
}

type SwitchTargetRequest struct {
	ProductSlug string `json:"product_slug"`
	TargetType  string `json:"target_type"`
	ToVersionID string `json:"to_version_id"`
	Operator    string `json:"operator"`
}

type RollbackTargetRequest struct {
	ProductSlug           string `json:"product_slug"`
	TargetType            string `json:"target_type"`
	RolledBackToVersionID string `json:"rolled_back_to_version_id"`
	Operator              string `json:"operator"`
}

type SwitchEvent struct {
	ID            string `json:"id"`
	ProductSlug   string `json:"product_slug"`
	TargetType    string `json:"target_type"`
	FromVersionID string `json:"from_version_id,omitempty"`
	ToVersionID   string `json:"to_version_id"`
	Operator      string `json:"operator"`
}

type RollbackEvent struct {
	ID                    string `json:"id"`
	ProductSlug           string `json:"product_slug"`
	TargetType            string `json:"target_type"`
	RolledBackFromVersion string `json:"rolled_back_from_version_id"`
	RolledBackToVersion   string `json:"rolled_back_to_version_id"`
	Operator              string `json:"operator"`
}

type TargetBundle struct {
	Client  VersionTarget `json:"client"`
	Service VersionTarget `json:"service"`
}
