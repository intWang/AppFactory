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
}
