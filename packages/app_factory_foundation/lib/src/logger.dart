class AppLogger {
  void info(String message) {
    // Keep the default logger lightweight for shared usage.
    // ignore: avoid_print
    print('[INFO] $message');
  }
}
