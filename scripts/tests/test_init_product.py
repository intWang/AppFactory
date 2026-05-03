import tempfile
import unittest
from pathlib import Path

from scripts.init_product import create_product_scaffold, slug_to_package_name


def write_template(path: Path, content: str) -> None:
  path.parent.mkdir(parents=True, exist_ok=True)
  path.write_text(content, encoding='utf-8')


class InitProductTests(unittest.TestCase):
  def test_creates_expected_product_structure_and_files(self) -> None:
    with tempfile.TemporaryDirectory() as tmpdir:
      repo_root = Path(tmpdir)
      write_template(repo_root / 'templates/docs/intake.template.md', '# Intake\n\n## Project Name\n\n## Project Slug\n')
      write_template(repo_root / 'templates/docs/prd-draft.template.md', '# PRD Draft\n\n## Project Name\n\n## Project Slug\n')
      write_template(repo_root / 'templates/docs/ux-spec.template.md', '# UX Spec\n\n## Project Name\n\n## Project Slug\n')
      write_template(repo_root / 'templates/docs/prd-final.template.md', '# PRD\n\n## Project Name\n\n## Project Slug\n')
      write_template(repo_root / 'templates/docs/add.template.md', '# Architecture Design\n\n## Project Name\n\n## Project Slug\n')
      write_template(repo_root / 'templates/docs/dev-plan.template.md', '# Development Plan\n\n## Project Name\n\n## Project Slug\n')
      write_template(repo_root / 'templates/docs/test-cases.template.md', '# Test Cases\n\n## Project Name\n\n## Project Slug\n')
      write_template(repo_root / 'templates/docs/test-report.template.md', '# Test Report\n\n## Project Name\n\n## Project Slug\n')
      write_template(repo_root / 'templates/docs/release-gate.template.md', '# Release Gate\n\n## Project Name\n\n## Project Slug\n')
      write_template(repo_root / 'templates/docs/retro-input.template.md', '# Retro Input\n\n## Project Name\n\n## Project Slug\n')
      write_template(repo_root / 'templates/docs/figma-link.template.md', '# Figma Link\n\n## Project Name\n\n## Project Slug\n')

      create_product_scaffold(repo_root, 'focus-timer', 'Focus Timer')

      product_root = repo_root / 'products/focus-timer'
      self.assertTrue((product_root / 'client/pubspec.yaml').exists())
      self.assertTrue((product_root / 'client/config/app.yaml').exists())
      self.assertTrue((product_root / 'client/lib/app/app.dart').exists())
      self.assertTrue((product_root / 'client/lib/app/bootstrap.dart').exists())
      self.assertTrue((product_root / 'client/lib/features/home/home_page.dart').exists())
      self.assertTrue((product_root / 'client/test/unit/bootstrap_test.dart').exists())
      self.assertTrue((product_root / 'client/test/widget/app_smoke_test.dart').exists())
      self.assertTrue((product_root / 'docs/00-intake.md').exists())
      self.assertTrue((product_root / 'docs/03-prd-final.md').exists())
      self.assertTrue((product_root / 'design/figma-link.md').exists())
      self.assertTrue((product_root / 'design/exports/.gitkeep').exists())
      self.assertTrue((product_root / 'build/outputs/.gitkeep').exists())
      self.assertTrue((product_root / 'server/README.md').exists())
      self.assertTrue((product_root / 'server/config/server.yaml').exists())
      self.assertTrue((product_root / 'server/src/.gitkeep').exists())
      self.assertTrue((product_root / 'server/tests/.gitkeep').exists())
      self.assertTrue((product_root / 'server/deploy/.gitkeep').exists())

      pubspec = (product_root / 'client/pubspec.yaml').read_text(encoding='utf-8')
      self.assertIn('name: focus_timer', pubspec)

      config = (product_root / 'client/config/app.yaml').read_text(encoding='utf-8')
      self.assertIn('name: focus-timer', config)
      self.assertIn('server_mode: reserved', config)

      prd = (product_root / 'docs/03-prd-final.md').read_text(encoding='utf-8')
      self.assertIn('Focus Timer', prd)
      self.assertIn('focus-timer', prd)

      server_readme = (product_root / 'server/README.md').read_text(encoding='utf-8')
      self.assertIn('shared account', server_readme.lower())
      self.assertIn('pm and am', server_readme.lower())

  def test_refuses_to_overwrite_existing_product(self) -> None:
    with tempfile.TemporaryDirectory() as tmpdir:
      repo_root = Path(tmpdir)
      write_template(repo_root / 'templates/docs/intake.template.md', '# Intake\n')
      write_template(repo_root / 'templates/docs/prd-draft.template.md', '# PRD Draft\n')
      write_template(repo_root / 'templates/docs/ux-spec.template.md', '# UX Spec\n')
      write_template(repo_root / 'templates/docs/prd-final.template.md', '# PRD\n')
      write_template(repo_root / 'templates/docs/add.template.md', '# Architecture Design\n')
      write_template(repo_root / 'templates/docs/dev-plan.template.md', '# Development Plan\n')
      write_template(repo_root / 'templates/docs/test-cases.template.md', '# Test Cases\n')
      write_template(repo_root / 'templates/docs/test-report.template.md', '# Test Report\n')
      write_template(repo_root / 'templates/docs/release-gate.template.md', '# Release Gate\n')
      write_template(repo_root / 'templates/docs/retro-input.template.md', '# Retro Input\n')
      write_template(repo_root / 'templates/docs/figma-link.template.md', '# Figma Link\n')

      (repo_root / 'products/existing-app').mkdir(parents=True)

      with self.assertRaises(FileExistsError):
        create_product_scaffold(repo_root, 'existing-app', 'Existing App')

  def test_slug_to_package_name_converts_hyphens(self) -> None:
    self.assertEqual(slug_to_package_name('focus-timer-pro'), 'focus_timer_pro')


if __name__ == '__main__':
  unittest.main()
