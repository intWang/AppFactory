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
