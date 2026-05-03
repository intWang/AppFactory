import 'package:app_factory_growth/app_factory_growth.dart';
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

  testWidgets('shows optional upgrade prompt when a newer version exists', (WidgetTester tester) async {
    const UpgradeDecision decision = UpgradeDecision(
      mode: UpgradePromptMode.optional,
      currentVersion: '1.0.0',
      currentBuildNumber: 1,
      latestVersion: '1.0.1',
      latestBuildNumber: 2,
    );

    await tester.pumpWidget(
      const ProductShellApp(
        title: 'Template',
        home: Scaffold(body: Text('Template Product Home')),
        upgradeDecision: decision,
      ),
    );
    await tester.pump();

    expect(find.text('Update Available'), findsOneWidget);
    expect(find.text('Later'), findsOneWidget);
    expect(find.text('Upgrade'), findsOneWidget);
  });

  testWidgets('forces upgrade flow when the app is too far behind', (WidgetTester tester) async {
    bool forcedExitTriggered = false;
    const UpgradeDecision decision = UpgradeDecision(
      mode: UpgradePromptMode.forced,
      currentVersion: '1.0.0',
      currentBuildNumber: 1,
      latestVersion: '1.0.5',
      latestBuildNumber: 5,
    );

    await tester.pumpWidget(
      ProductShellApp(
        title: 'Template',
        home: const Scaffold(body: Text('Template Product Home')),
        upgradeDecision: decision,
        onForceUpgradeExit: () {
          forcedExitTriggered = true;
        },
      ),
    );
    await tester.pump();

    expect(find.text('Upgrade Required'), findsOneWidget);
    expect(find.text('Exit and Upgrade'), findsOneWidget);

    await tester.tap(find.text('Exit and Upgrade'));
    await tester.pumpAndSettle();

    expect(forcedExitTriggered, isTrue);
  });
}
