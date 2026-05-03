import 'package:example_tool_app/app/bootstrap.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  test('bootstrap returns app services', () async {
    final AppServices services = await bootstrapApp();

    expect(services.registry, isNotNull);
    expect(services.config.appName, 'Example Tool App');
  });
}
