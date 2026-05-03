import 'package:app_factory_growth/app_factory_growth.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_product_shell/app/bootstrap.dart';

void main() {
  test('stores docs and output directories', () {
    const ProductShellPaths paths = ProductShellPaths(
      docsDirectory: 'docs',
      outputDirectory: 'build/outputs',
    );

    expect(paths.docsDirectory, 'docs');
    expect(paths.outputDirectory, 'build/outputs');
  });

  test('bootstrap returns upgrade decision for app shell launch checks', () async {
    final ProductShellServices services = await bootstrapProductShell();

    expect(services.upgradeDecision.mode, UpgradePromptMode.none);
    expect(services.registry.resolve<UpgradeDecision>(), same(services.upgradeDecision));
  });
}
