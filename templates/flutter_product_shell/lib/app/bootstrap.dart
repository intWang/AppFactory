import 'package:app_factory_foundation/app_factory_foundation.dart';
import 'package:app_factory_growth/app_factory_growth.dart';

class ProductShellPaths {
  const ProductShellPaths({
    required this.docsDirectory,
    required this.outputDirectory,
  });

  final String docsDirectory;
  final String outputDirectory;
}

class ProductShellServices {
  const ProductShellServices({
    required this.config,
    required this.registry,
    required this.upgradeDecision,
  });

  final AppConfig config;
  final CapabilityRegistry registry;
  final UpgradeDecision upgradeDecision;
}

Future<ProductShellServices> bootstrapProductShell() async {
  final CapabilityRegistry registry = CapabilityRegistry();
  const GrowthEntryPoints growthEntryPoints = GrowthEntryPoints();
  registry.register<AppLogger>(AppLogger());
  registry.register<ErrorReporter>(ErrorReporter());
  registry.register<GrowthEntryPoints>(growthEntryPoints);

  const AppConfig config = AppConfig(
    appName: 'Template Product Shell',
    currentVersion: '0.1.0',
    currentBuildNumber: 1,
    latestVersion: '0.1.0',
    latestBuildNumber: 1,
  );

  final UpgradeDecision upgradeDecision = growthEntryPoints.upgradePolicy.evaluate(
    currentVersion: config.currentVersion,
    currentBuildNumber: config.currentBuildNumber,
    latestVersion: config.latestVersion,
    latestBuildNumber: config.latestBuildNumber,
    upgradeUrl: config.upgradeUrl,
  );
  registry.register<UpgradeDecision>(upgradeDecision);

  return ProductShellServices(
    config: config,
    registry: registry,
    upgradeDecision: upgradeDecision,
  );
}
