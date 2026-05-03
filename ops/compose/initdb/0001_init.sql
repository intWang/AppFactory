CREATE TABLE IF NOT EXISTS users (
  id TEXT PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  nickname TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS local_credentials (
  user_id TEXT PRIMARY KEY REFERENCES users(id),
  password_hash TEXT NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS auth_identities (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL REFERENCES users(id),
  provider_key TEXT NOT NULL,
  provider_subject TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE (provider_key, provider_subject)
);

CREATE TABLE IF NOT EXISTS sessions (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL REFERENCES users(id),
  session_token TEXT NOT NULL UNIQUE,
  expires_at TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS provider_links (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL REFERENCES users(id),
  provider_key TEXT NOT NULL,
  linked_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS release_channels (
  id TEXT PRIMARY KEY,
  channel_key TEXT NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS release_versions (
  id TEXT PRIMARY KEY,
  product_slug TEXT NOT NULL,
  target_type TEXT NOT NULL,
  version_label TEXT NOT NULL,
  build_number INTEGER NOT NULL,
  upgrade_url TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS deployment_records (
  id TEXT PRIMARY KEY,
  target_version_id TEXT NOT NULL REFERENCES release_versions(id),
  environment TEXT NOT NULL,
  deployed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  status TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS active_targets (
  id TEXT PRIMARY KEY,
  product_slug TEXT NOT NULL,
  target_type TEXT NOT NULL,
  active_version_id TEXT NOT NULL REFERENCES release_versions(id),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE (product_slug, target_type)
);

CREATE TABLE IF NOT EXISTS switch_events (
  id TEXT PRIMARY KEY,
  product_slug TEXT NOT NULL,
  target_type TEXT NOT NULL,
  from_version_id TEXT,
  to_version_id TEXT NOT NULL REFERENCES release_versions(id),
  switched_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  operator TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS rollback_events (
  id TEXT PRIMARY KEY,
  product_slug TEXT NOT NULL,
  target_type TEXT NOT NULL,
  rolled_back_from_version_id TEXT NOT NULL REFERENCES release_versions(id),
  rolled_back_to_version_id TEXT NOT NULL REFERENCES release_versions(id),
  rolled_back_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  operator TEXT NOT NULL
);

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

INSERT INTO release_channels (id, channel_key)
VALUES ('channel-local', 'local')
ON CONFLICT (channel_key) DO NOTHING;

INSERT INTO release_versions (id, product_slug, target_type, version_label, build_number, upgrade_url)
VALUES
  ('rv-client-2622004', 'shared-client', 'client', '26.2.20.04', 4, 'https://example.com/client-upgrade'),
  ('rv-service-2622003', 'shared-service', 'service', '26.2.20.03', 3, 'https://example.com/service-upgrade')
ON CONFLICT (id) DO NOTHING;

INSERT INTO active_targets (id, product_slug, target_type, active_version_id)
VALUES
  ('at-client-local', 'shared-client', 'client', 'rv-client-2622004'),
  ('at-service-local', 'shared-service', 'service', 'rv-service-2622003')
ON CONFLICT (product_slug, target_type) DO UPDATE
SET active_version_id = EXCLUDED.active_version_id,
    updated_at = NOW();
