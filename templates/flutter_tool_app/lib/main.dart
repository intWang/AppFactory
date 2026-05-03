import 'package:flutter/material.dart';
import 'package:flutter_tool_app/app/app.dart';
import 'package:flutter_tool_app/app/bootstrap.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final AppServices services = await bootstrapApp();
  runApp(ToolApp(services: services));
}
