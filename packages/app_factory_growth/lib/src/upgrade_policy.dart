import 'upgrade_decision.dart';

class UpgradePolicy {
  const UpgradePolicy({
    this.forcedUpgradeBuildGap = 3,
  });

  final int forcedUpgradeBuildGap;

  UpgradeDecision evaluate({
    required String currentVersion,
    required int currentBuildNumber,
    required String latestVersion,
    required int latestBuildNumber,
    String? upgradeUrl,
  }) {
    final int buildGap = latestBuildNumber - currentBuildNumber;
    if (buildGap <= 0) {
      return UpgradeDecision(
        mode: UpgradePromptMode.none,
        currentVersion: currentVersion,
        currentBuildNumber: currentBuildNumber,
        latestVersion: latestVersion,
        latestBuildNumber: latestBuildNumber,
        upgradeUrl: upgradeUrl,
      );
    }

    final UpgradePromptMode mode = buildGap > forcedUpgradeBuildGap
        ? UpgradePromptMode.forced
        : UpgradePromptMode.optional;

    return UpgradeDecision(
      mode: mode,
      currentVersion: currentVersion,
      currentBuildNumber: currentBuildNumber,
      latestVersion: latestVersion,
      latestBuildNumber: latestBuildNumber,
      upgradeUrl: upgradeUrl,
    );
  }
}
