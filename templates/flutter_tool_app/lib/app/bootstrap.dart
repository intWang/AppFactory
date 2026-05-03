import 'package:flutter_tool_app/core/capabilities/capability_registry.dart';
import 'package:flutter_tool_app/core/foundation/config/app_config.dart';
import 'package:flutter_tool_app/core/foundation/error/error_reporter.dart';
import 'package:flutter_tool_app/core/foundation/logging/logger.dart';
import 'package:flutter_tool_app/core/growth/growth_entry_points.dart';

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

  const AppConfig config = AppConfig(appName: 'App Factory Tool App');
  return const AppServices(
    config: config,
    registry: registry,
  );
}
