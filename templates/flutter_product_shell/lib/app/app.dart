import 'package:flutter/material.dart';

class ProductShellApp extends StatelessWidget {
  const ProductShellApp({
    super.key,
    required this.title,
    required this.home,
  });

  final String title;
  final Widget home;

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: title,
      home: home,
    );
  }
}
