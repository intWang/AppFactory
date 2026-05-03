class CapabilityRegistry {
  final Map<Type, Object> _capabilities = <Type, Object>{};

  void register<T extends Object>(T instance) {
    _capabilities[T] = instance;
  }

  T resolve<T extends Object>() {
    final Object? value = _capabilities[T];
    if (value == null) {
      throw StateError('Capability of type $T has not been registered.');
    }
    return value as T;
  }
}
