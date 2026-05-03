import 'package:app_factory_growth/app_factory_growth.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  test('returns no prompt when latest build is not newer', () {
    const UpgradePolicy policy = UpgradePolicy();

    final UpgradeDecision decision = policy.evaluate(
      currentVersion: '1.0.0',
      currentBuildNumber: 4,
      latestVersion: '1.0.0',
      latestBuildNumber: 4,
    );

    expect(decision.shouldPrompt, isFalse);
    expect(decision.isForced, isFalse);
  });

  test('returns optional prompt when newer build is within threshold', () {
    const UpgradePolicy policy = UpgradePolicy();

    final UpgradeDecision decision = policy.evaluate(
      currentVersion: '1.0.0',
      currentBuildNumber: 4,
      latestVersion: '1.0.2',
      latestBuildNumber: 6,
    );

    expect(decision.shouldPrompt, isTrue);
    expect(decision.mode, UpgradePromptMode.optional);
  });

  test('returns forced prompt when build gap exceeds threshold', () {
    const UpgradePolicy policy = UpgradePolicy();

    final UpgradeDecision decision = policy.evaluate(
      currentVersion: '1.0.0',
      currentBuildNumber: 2,
      latestVersion: '1.0.5',
      latestBuildNumber: 6,
    );

    expect(decision.shouldPrompt, isTrue);
    expect(decision.isForced, isTrue);
  });
}
