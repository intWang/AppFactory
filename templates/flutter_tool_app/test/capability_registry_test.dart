import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_tool_app/core/capabilities/capability_registry.dart';

void main() {
  test('registers and resolves capabilities by type', () {
    final registry = CapabilityRegistry();

    registry.register<String>('logger');

    expect(registry.resolve<String>(), 'logger');
  });
}
