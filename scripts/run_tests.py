#!/usr/bin/env python3
"""
SQLLogicTest runner for the SQL Vibe Coding Challenge.

This script runs SQLLogicTest files against your database implementation
and reports pass/fail statistics.
"""

import os
import subprocess
import sys
from pathlib import Path
from typing import Tuple

# Path to the SQLLogicTest test files
SCRIPT_DIR = Path(__file__).parent
PROJECT_ROOT = SCRIPT_DIR.parent
TEST_DIR = PROJECT_ROOT / "third_party" / "sqllogictest" / "test"
DB_BINARY = PROJECT_ROOT / "sql-challenge"


def run_test_file(test_file: Path) -> Tuple[bool, str]:
    """Run a single test file and return (passed, error_message)."""
    try:
        result = subprocess.run(
            [str(DB_BINARY), str(test_file)],
            capture_output=True,
            text=True,
            timeout=60,
        )
        # For now, consider any non-error exit as a pass
        # You'll want to implement proper result checking
        if result.returncode == 0:
            return True, ""
        else:
            return False, result.stderr or result.stdout
    except subprocess.TimeoutExpired:
        return False, "Timeout"
    except Exception as e:
        return False, str(e)


def find_test_files() -> list[Path]:
    """Find all .test files in the test directory."""
    if not TEST_DIR.exists():
        print(f"Error: Test directory not found: {TEST_DIR}")
        print("Make sure you cloned with --recurse-submodules")
        sys.exit(1)

    test_files = []
    for root, _, files in os.walk(TEST_DIR):
        for file in files:
            if file.endswith(".test"):
                test_files.append(Path(root) / file)

    return sorted(test_files)


def main():
    # Check if binary exists
    if not DB_BINARY.exists():
        print(f"Error: Database binary not found: {DB_BINARY}")
        print("Run 'make build' first")
        sys.exit(1)

    # Find all test files
    test_files = find_test_files()
    print(f"Found {len(test_files)} test files")
    print()

    # Run tests
    passed = 0
    failed = 0
    summary_only = "--summary" in sys.argv

    for test_file in test_files:
        relative_path = test_file.relative_to(TEST_DIR)
        success, error = run_test_file(test_file)

        if success:
            passed += 1
            if not summary_only:
                print(f"PASS: {relative_path}")
        else:
            failed += 1
            if not summary_only:
                print(f"FAIL: {relative_path}")
                if error:
                    print(f"      {error[:100]}")

    # Print summary
    total = passed + failed
    percentage = (passed / total * 100) if total > 0 else 0

    print()
    print("=" * 50)
    print(f"Results: {passed}/{total} files passed ({percentage:.1f}%)")
    print("=" * 50)

    # Exit with error if not all tests passed
    sys.exit(0 if failed == 0 else 1)


if __name__ == "__main__":
    main()
