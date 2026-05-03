CREATE TABLE IF NOT EXISTS service_manager_operations (
  id TEXT PRIMARY KEY,
  operation TEXT NOT NULL,
  product_slug TEXT NOT NULL,
  target_type TEXT NOT NULL,
  target_version_id TEXT,
  environment TEXT,
  operator TEXT,
  status TEXT NOT NULL,
  started_at TIMESTAMPTZ NOT NULL,
  completed_at TIMESTAMPTZ,
  error TEXT
);
