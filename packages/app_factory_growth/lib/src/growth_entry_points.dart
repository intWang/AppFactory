import 'upgrade_policy.dart';

class GrowthEntryPoints {
  const GrowthEntryPoints({
    this.hasUpgradePrompt = true,
    this.upgradePolicy = const UpgradePolicy(),
  });

  final bool hasUpgradePrompt;
  final UpgradePolicy upgradePolicy;
}
