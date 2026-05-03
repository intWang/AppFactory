CREATE TABLE release_channels (
  id TEXT PRIMARY KEY,
  channel_key TEXT NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE release_versions (
  id TEXT PRIMARY KEY,
  product_slug TEXT NOT NULL,
  target_type TEXT NOT NULL,
  version_label TEXT NOT NULL,
  build_number INTEGER NOT NULL,
  upgrade_url TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE deployment_records (
  id TEXT PRIMARY KEY,
  target_version_id TEXT NOT NULL REFERENCES release_versions(id),
  environment TEXT NOT NULL,
  deployed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  status TEXT NOT NULL
);

CREATE TABLE active_targets (
  id TEXT PRIMARY KEY,
  product_slug TEXT NOT NULL,
  target_type TEXT NOT NULL,
  active_version_id TEXT NOT NULL REFERENCES release_versions(id),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE (product_slug, target_type)
);

CREATE TABLE switch_events (
  id TEXT PRIMARY KEY,
  product_slug TEXT NOT NULL,
  target_type TEXT NOT NULL,
  from_version_id TEXT,
  to_version_id TEXT NOT NULL REFERENCES release_versions(id),
  switched_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  operator TEXT NOT NULL
);

CREATE TABLE rollback_events (
  id TEXT PRIMARY KEY,
  product_slug TEXT NOT NULL,
  target_type TEXT NOT NULL,
  rolled_back_from_version_id TEXT NOT NULL REFERENCES release_versions(id),
  rolled_back_to_version_id TEXT NOT NULL REFERENCES release_versions(id),
  rolled_back_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  operator TEXT NOT NULL
);
