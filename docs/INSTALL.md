# Installation Guide

`mls` (MrLeanStorage) is available as pre-built binaries for macOS, Linux, and Windows, or can be compiled directly from source.

---

## 1. macOS (Homebrew Cask) — Recommended

The easiest and most modern way to install `mls` on macOS is through our custom Homebrew Tap. Since `mls` is distributed as a pre-compiled binary for optimal performance, it is packaged as a **Homebrew Cask** in accordance with GoReleaser v2 guidelines.

### Standard Installation

To tap the repository and install `mls`:

```bash
# Add the custom tap
brew tap MohamedLamineAllal/mls

# Install the mls cask
brew install mls
```

*Note: Homebrew will automatically resolve `mls` to the cask distribution since the older source-based formula has been fully deprecated and deleted.*

Alternatively, you can be explicit:

```bash
brew install --cask mls
```

### Automatic Quarantine Bypass

On macOS, binaries downloaded via browsers or third-party installers are automatically placed in quarantine. The `mls` cask includes a post-install hook that automatically strips the macOS quarantine attribute so you can run the CLI immediately without gatekeeper warnings:

```bash
/usr/bin/xattr -dr com.apple.quarantine /path/to/mls
```

### Troubleshooting: Checksum Mismatch / Formula Errors

If you previously had the older formula installed or cached, you might receive a verification checksum error (`Error: Formula reports different checksum: replace_with_sha256_hash`). 

To resolve this and sync your local Homebrew Tap cache with the latest remote tap deletion, run:

```bash
# Update Homebrew and prune deleted formulas
brew update
brew tap --repair

# Re-run installation
brew install mls
```

---

## 2. Linux (Pre-built Packages)

Since Homebrew Casks are a macOS-only feature, Linux users should install `mls` using our pre-built Debian/RPM packages or manual binary extraction.

### Debian / Ubuntu (`.deb`)

```bash
# Download the latest deb package
wget https://github.com/MohamedLamineAllal/MrLeanStorage/releases/latest/download/mls_0.1.3_linux_amd64.deb

# Install package
sudo dpkg -i mls_0.1.3_linux_amd64.deb
```

### Fedora / RedHat / CentOS (`.rpm`)

```bash
# Download the latest rpm package
wget https://github.com/MohamedLamineAllal/MrLeanStorage/releases/latest/download/mls_0.1.3_linux_amd64.rpm

# Install package
sudo rpm -ivh mls_0.1.3_linux_amd64.rpm
```

### Manual Binary Installation (Any Linux Distribution)

```bash
# Download and extract the pre-built tarball to /usr/local/bin
curl -sL https://github.com/MohamedLamineAllal/MrLeanStorage/releases/latest/download/mls_0.1.3_linux_amd64.tar.gz | tar xz -C /usr/local/bin mls
```

---

## 3. Windows (Manual Installation)

For Windows environments, download and extract the ZIP archive containing the pre-compiled binary.

```powershell
# Create a bin directory in your User Profile if it doesn't exist
New-Item -ItemType Directory -Force -Path "$HOME\bin"

# Download the latest release zip
Invoke-WebRequest -Uri "https://github.com/MohamedLamineAllal/MrLeanStorage/releases/latest/download/mls_0.1.3_windows_amd64.zip" -OutFile "$HOME\bin\mls.zip"

# Expand the archive
Expand-Archive -Path "$HOME\bin\mls.zip" -DestinationPath "$HOME\bin" -Force

# Clean up zip
Remove-Item -Path "$HOME\bin\mls.zip"

# Add $HOME\bin to your User PATH environment variable if not already present
[Environment]::SetEnvironmentVariable("Path", [Environment]::GetEnvironmentVariable("Path", "User") + ";$HOME\bin", "User")
```

---

## 4. Building From Source (All Platforms)

If you prefer to compile `mls` yourself from the latest source, ensure you have Go 1.26+ installed.

```bash
# Clone the repository
git clone https://github.com/MohamedLamineAllal/MrLeanStorage.git
cd MrLeanStorage

# Compile the binary
go build -o mls main.go

# Install the binary globally
sudo mv mls /usr/local/bin/mls
```
