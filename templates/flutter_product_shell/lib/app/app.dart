import 'package:app_factory_growth/app_factory_growth.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class ProductShellApp extends StatelessWidget {
  const ProductShellApp({
    super.key,
    required this.title,
    required this.home,
    this.upgradeDecision,
    this.onForceUpgradeExit,
  });

  final String title;
  final Widget home;
  final UpgradeDecision? upgradeDecision;
  final VoidCallback? onForceUpgradeExit;

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: title,
      home: _UpgradeGate(
        upgradeDecision: upgradeDecision,
        onForceUpgradeExit: onForceUpgradeExit,
        child: home,
      ),
    );
  }
}

class _UpgradeGate extends StatefulWidget {
  const _UpgradeGate({
    required this.upgradeDecision,
    required this.child,
    this.onForceUpgradeExit,
  });

  final UpgradeDecision? upgradeDecision;
  final VoidCallback? onForceUpgradeExit;
  final Widget child;

  @override
  State<_UpgradeGate> createState() => _UpgradeGateState();
}

class _UpgradeGateState extends State<_UpgradeGate> {
  bool _dialogShown = false;

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    final UpgradeDecision? upgradeDecision = widget.upgradeDecision;
    if (_dialogShown || upgradeDecision == null || !upgradeDecision.shouldPrompt) {
      return;
    }

    _dialogShown = true;
    WidgetsBinding.instance.addPostFrameCallback((_) {
      if (!mounted) {
        return;
      }

      showDialog<void>(
        context: context,
        barrierDismissible: !upgradeDecision.isForced,
        builder: (BuildContext context) {
          return AlertDialog(
            title: Text(upgradeDecision.isForced ? 'Upgrade Required' : 'Update Available'),
            content: Text(
              upgradeDecision.isForced
                  ? 'Your app version is too old. Please upgrade now to continue using the app.'
                  : 'A newer app version is available. Upgrade now for the latest fixes and features.',
            ),
            actions: <Widget>[
              if (!upgradeDecision.isForced)
                TextButton(
                  onPressed: () => Navigator.of(context).pop(),
                  child: const Text('Later'),
                ),
              FilledButton(
                onPressed: () {
                  Navigator.of(context).pop();
                  if (upgradeDecision.isForced) {
                    if (widget.onForceUpgradeExit != null) {
                      widget.onForceUpgradeExit!.call();
                    } else {
                      SystemNavigator.pop();
                    }
                  }
                },
                child: Text(upgradeDecision.isForced ? 'Exit and Upgrade' : 'Upgrade'),
              ),
            ],
          );
        },
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return widget.child;
  }
}
