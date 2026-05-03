import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_tool_app/app/bootstrap.dart';

void main() {
  test('bootstrap returns app services', () async {
    final services = await bootstrapApp();

    expect(services.registry, isNotNull);
    expect(services.config.appName, isNotEmpty);
  });
}
