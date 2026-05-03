import 'package:example_tool_app/app/app.dart';
import 'package:example_tool_app/app/bootstrap.dart';
import 'package:flutter/material.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final AppServices services = await bootstrapApp();
  runApp(ExampleToolApp(services: services));
}
