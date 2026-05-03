package domain

import "testing"

func TestUpgradeModeUsesSharedThreeBuildRule(t *testing.T) {
	target := VersionTarget{
		CurrentBuild:      1,
		LatestBuild:       4,
		ForceUpgradeAfter: 3,
	}
	if got := target.UpgradeMode(); got != "optional" {
		t.Fatalf("expected optional upgrade for three-build gap, got %s", got)
	}

	target.LatestBuild = 5
	if got := target.UpgradeMode(); got != "forced" {
		t.Fatalf("expected forced upgrade for greater-than-three-build gap, got %s", got)
	}
}
