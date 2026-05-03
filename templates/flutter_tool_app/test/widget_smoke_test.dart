import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_tool_app/app/app.dart';
import 'package:flutter_tool_app/app/bootstrap.dart';

void main() {
  testWidgets('renders tool app shell', (tester) async {
    final services = await bootstrapApp();

    await tester.pumpWidget(ToolApp(services: services));

    expect(find.byType(MaterialApp), findsOneWidget);
    expect(find.text('Tool App Home'), findsOneWidget);
  });
}
