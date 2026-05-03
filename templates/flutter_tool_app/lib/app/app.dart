import 'package:flutter/material.dart';
import 'package:flutter_tool_app/app/bootstrap.dart';
import 'package:flutter_tool_app/core/ui/app_theme.dart';
import 'package:flutter_tool_app/features/home/home_page.dart';

class ToolApp extends StatelessWidget {
  const ToolApp({
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
