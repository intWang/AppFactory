import 'package:example_tool_app/app/app.dart';
import 'package:example_tool_app/app/bootstrap.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('renders product app shell', (WidgetTester tester) async {
    final AppServices services = await bootstrapApp();

    await tester.pumpWidget(ExampleToolApp(services: services));

    expect(find.byType(MaterialApp), findsOneWidget);
    expect(find.text('Example Tool App Home'), findsOneWidget);
  });
}
