import 'package:app_factory_foundation/app_factory_foundation.dart';
import 'package:app_factory_growth/app_factory_growth.dart';

class AppServices {
  const AppServices({
    required this.config,
    required this.registry,
  });

  final AppConfig config;
  final CapabilityRegistry registry;
}

Future<AppServices> bootstrapApp() async {
  final CapabilityRegistry registry = CapabilityRegistry();
  registry.register<AppLogger>(AppLogger());
  registry.register<ErrorReporter>(ErrorReporter());
  registry.register<GrowthEntryPoints>(const GrowthEntryPoints());

  const AppConfig config = AppConfig(appName: 'Example Tool App');
  return AppServices(
    config: config,
    registry: registry,
  );
}
