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

func TestVersionTargetCanEvaluateRequestSpecificBuildGap(t *testing.T) {
	target := VersionTarget{
		LatestBuild:       7,
		ForceUpgradeAfter: 3,
	}

	if got := target.UpgradeModeForBuild(4); got != "optional" {
		t.Fatalf("expected optional upgrade for build gap 3, got %s", got)
	}

	if got := target.UpgradeModeForBuild(3); got != "forced" {
		t.Fatalf("expected forced upgrade for build gap 4, got %s", got)
	}
}
