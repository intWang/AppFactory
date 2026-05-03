import 'package:flutter/material.dart';
import 'package:flutter_product_shell/app/app.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('renders product shell app', (WidgetTester tester) async {
    await tester.pumpWidget(
      const ProductShellApp(
        title: 'Template',
        home: Scaffold(body: Text('Template Product Home')),
      ),
    );

    expect(find.byType(MaterialApp), findsOneWidget);
    expect(find.text('Template Product Home'), findsOneWidget);
  });
}
