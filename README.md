# MrLeanStorage (mls)

`mls` is a high-performance, concurrent storage cleanup tool for macOS (with cross-platform support for Linux and Windows), designed to safely and efficiently reclaim disk space. Written in Go, it features a small memory footprint, a seamless user experience, and a built-in background daemon.

---

## ⚡ Key Features

- **🚀 Concurrency Engine**: Uses high-performance goroutine worker pools in both the scanner (`os.ReadDir` based) and cleaner to scan and delete files simultaneously, maximizing macOS SSD throughput.
- **🔄 Zero-Downtime Hot Reloads**: Supports instant configuration reloading via `SIGHUP` signal without terminating the running background daemon.
- **🛌 Missed-Task Recovery Ticker**: Background daemon runs an automatic missed-task recovery ticker every 30 minutes to catch up on cleanup runs that were missed while your Mac was asleep.
- **📦 Application Cache Migration**: State files (e.g., last-run logs) are safely persisted under the persistent local cache (`~/Library/Caches/mls`) rather than the volatile `/tmp` directory.
- **🛡️ Dry-Run Safety**: Defaults to a strict dry-run mode so you can preview exactly which files will be deleted before taking any destructive action.
- **🤖 launchd Background Integration**: Complete agent management CLI to install, start, stop, restart, and inspect background services seamlessly on macOS.

---

## 📦 Installation

For full multi-platform instructions, see the detailed [Installation Guide](./docs/INSTALL.md).

### macOS (Homebrew Cask) — Recommended

`mls` is distributed as a Homebrew Cask via a custom Tap for seamless macOS installation and automatic quarantine bypass:

```bash
# Add our custom Tap
brew tap MohamedLamineAllal/mls

# Install mls
brew install mls
```

*Note: Homebrew will automatically map this to our Cask distribution. If you encounter any checksum errors from outdated Formula caches, resolve them by running `brew update && brew tap --repair` first.*

### Linux & Windows

For Debian/Ubuntu (`.deb`), RedHat/Fedora (`.rpm`), or Windows manual installations, please refer to the [Installation Guide](./docs/INSTALL.md#2-linux-pre-built-packages).

---

## 🚀 Usage

### 1. Initialize & Open Configuration

On first run, `mls` automatically creates a default configuration file at `~/.MrLeanStorage.yaml`.

```bash
# Open configuration in your default editor
mls config open

# Reveal configuration location in Finder
mls config reveal
```

### 2. Scan for Old Files

Analyze the configured targets and view matched files and sizes:

```bash
mls scan
```

Use the verbose flag to output all matches beyond the default summary:

```bash
mls scan -v
```

### 3. Clean Files (Dry Run & Confirmation)

Dry run to preview deletions:

```bash
mls clean
```

Execute the real cleanup by disabling dry-run:

```bash
mls clean --dry-run=false
```

---

## ⏰ Background Automation (macOS)

Manage the background daemon seamlessly using standard `launchd` controls built right into the CLI:

```bash
# Install the launchd background agent
mls agent install

# Start the background service
mls agent start

# Check background daemon status
mls agent status

# Restart / Hot reload configuration
mls agent restart

# Stop the background service
mls agent stop

# Uninstall the background agent completely
mls agent uninstall
```

---

## 🛠️ Configuration Example

The `~/.MrLeanStorage.yaml` configuration uses simple and flexible YAML format:

```yaml
# Global safety switch. If true, no files are ever deleted.
dry_run: true

# Patterns to globally ignore during scanning and deep staleness checks
ignore_patterns:
  - ".DS_Store"
  - "._*"
  - ".Spotlight-V100"
  - ".Trashes"
  - ".fseventsd"

# Cron scheduling expression (supports 6-field standard with seconds)
# Format: Second Minute Hour DayOfMonth Month DayOfWeek
schedule: "0 0 0 * * *"

# Target directories to monitor and system commands to execute
targets:
  - name: "VSCode Caches"
    path: "~/Library/Caches/com.microsoft.VSCode"
    threshold_days: 7
    type: "file" # "file", "folder", or "both"
    safety_level: 1

  - name: "Chrome Caches"
    path: "~/Library/Caches/Google/Chrome/Default/Cache"
    threshold_days: 14
    type: "file"
    safety_level: 1

  - name: "PNPM Global Pruning"
    command: "pnpm store prune"
    interval_days: 7 # Run this command target once every 7 days
```

Check [docs/configuration/Examples/](./docs/configuration/Examples/) for extensive config references.

---

## 🧪 Testing

Run the test suite with race condition detection:

```bash
go test -race ./...
```

---

## 📄 License

MIT License. See [LICENSE](LICENSE) for details.
