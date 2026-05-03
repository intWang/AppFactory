import 'package:app_factory_ui/app_factory_ui.dart';
import 'package:example_tool_app/app/bootstrap.dart';
import 'package:example_tool_app/features/home/home_page.dart';
import 'package:flutter/material.dart';

class ExampleToolApp extends StatelessWidget {
  const ExampleToolApp({
    super.key,
    required this.services,
  });

  final AppServices services;

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: services.config.appName,
      theme: buildAppTheme(),
      home: const HomePage(),
    );
  }
}
