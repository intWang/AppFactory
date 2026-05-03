class AppLogger {
  void info(String message) {
    // Keep the default logger lightweight for template projects.
    // ignore: avoid_print
    print('[INFO] $message');
  }
}
