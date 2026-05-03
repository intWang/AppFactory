enum UpgradePromptMode {
  none,
  optional,
  forced,
}

class UpgradeDecision {
  const UpgradeDecision({
    required this.mode,
    required this.currentVersion,
    required this.currentBuildNumber,
    required this.latestVersion,
    required this.latestBuildNumber,
    this.upgradeUrl,
  });

  final UpgradePromptMode mode;
  final String currentVersion;
  final int currentBuildNumber;
  final String latestVersion;
  final int latestBuildNumber;
  final String? upgradeUrl;

  bool get shouldPrompt => mode != UpgradePromptMode.none;
  bool get isForced => mode == UpgradePromptMode.forced;
}
