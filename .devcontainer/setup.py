#!/usr/bin/env python3
"""Provision the PMM devcontainer: RPM packages, Go toolchain, make init, PG setup.

See CONTRIBUTING.md. Idempotent: a marker file in the user's home records a
completed setup so re-runs exit immediately.

NOTE (multi-agent, bead cosmos-main-jmg3): shell=True removed root-cause —
every command runs as an explicit argv list (S602), and the pg_hba.conf lines
are appended in Python instead of `echo >>`. S404/S603 (subprocess use by a
dev-env provisioning tool) are covered by the repo-wide
`**/.devcontainer/**` per-file-ignores scope, same class as `**/*scripts*`.
"""

import os
import pathlib
import subprocess
import sys
import time

GO_VERSION = os.getenv("GO_VERSION")
if GO_VERSION is None:
    msg = "GO_VERSION is not set"
    raise RuntimeError(msg)

PG_CONF = pathlib.Path("/srv/postgres14/postgresql.conf")
PG_HBA = pathlib.Path("/srv/postgres14/pg_hba.conf")


def run_commands(commands: list[list[str]]) -> None:
    """Runs given commands (argv lists) and checks exit codes."""
    for cmd in commands:
        subprocess.check_call(cmd)


def install_packages() -> None:
    """Installs required and useful RPM packages."""
    run_commands([
        [
            "dnf",
            "install",
            "-y",
            "--enablerepo=ol9_codeready_builder",
            "gcc",
            "git",
            "make",
            "pkgconfig",
            "vim",
            "psmisc",
            "procps",
            "lsof",
            "diffutils",
            "man",
            "man-pages",
            "ansible-lint",
            "glibc-static",
            "openssl-devel",
            "krb5-devel",
        ],
    ])


def install_go() -> None:
    """Installs Go toolchain."""
    run_commands([
        [
            "curl",
            "-sS",
            "https://raw.githubusercontent.com/travis-ci/gimme/v1.5.6/gimme",
            "-o",
            "/usr/local/bin/gimme",
        ],
        ["chmod", "+x", "/usr/local/bin/gimme"],
    ])

    go_version = subprocess.check_output(
        ["/usr/local/bin/gimme", "-r", GO_VERSION],
        text=True,
    ).strip()

    gimme_go_dir = f"go{go_version}.linux.amd64"

    run_commands([
        ["/usr/local/bin/gimme", go_version],
        ["rm", "-fr", "/usr/local/go"],
        ["mv", "-f", f"/root/.gimme/versions/{gimme_go_dir}", "/usr/local/go"],
        [
            "update-alternatives",
            "--install",
            "/usr/bin/go",
            "go",
            "/usr/local/go/bin/go",
            "0",
        ],
        ["update-alternatives", "--set", "go", "/usr/local/go/bin/go"],
        [
            "update-alternatives",
            "--install",
            "/usr/bin/gofmt",
            "gofmt",
            "/usr/local/go/bin/gofmt",
            "0",
        ],
        ["update-alternatives", "--set", "gofmt", "/usr/local/go/bin/gofmt"],
        ["mkdir", "-p", "/root/go/bin"],
        ["go", "version"],
        ["go", "env"],
    ])


def make_init() -> None:
    """Runs make init."""
    run_commands([
        ["make", "init"],
    ])


def _sed_uncomment(pattern: str, replacement: str) -> None:
    """Replaces a commented PG config directive with the active form."""
    subprocess.check_call([
        "/usr/bin/sed",
        "-i",
        "-e",
        f"s/#{pattern}/{replacement}/",
        str(PG_CONF),
    ])


def setup() -> None:
    """Runs various setup commands."""
    # allow connecting from any host, needed to connect from host to PG running in docker
    _sed_uncomment("listen_addresses = 'localhost'", "listen_addresses = '*'")
    # Turns fsync off. Create database operations with fsync on are very slow on Ubuntu.
    # Having fsync off in dev environment is fine.
    _sed_uncomment("fsync = on", "fsync = off")
    # Configure pg_hba.conf for password authentication from all hosts (dev environment only)
    # Note: In dev, we allow both trust and scram-sha-256 for convenience
    with PG_HBA.open("a", encoding="utf-8") as hba:
        hba.write("host    all         all     0.0.0.0/0     trust\n")
        hba.write("host    all         all     0.0.0.0/0     scram-sha-256\n")


def main() -> None:
    """Provisions the devcontainer (packages, Go, make init, PG setup)."""
    install_packages()
    install_go()
    make_init()

    # do basic setup
    setup()


MARKER = pathlib.Path.home() / ".devcontainer-setup-done"
if MARKER.exists():
    sys.exit(0)

start = time.time()
main()

MARKER.touch()
