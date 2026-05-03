import 'package:app_factory_foundation/app_factory_foundation.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  test('registers and resolves capabilities by type', () {
    final CapabilityRegistry registry = CapabilityRegistry();

    registry.register<String>('logger');

    expect(registry.resolve<String>(), 'logger');
  });
}
