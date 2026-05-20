# Release Process & Best Practices

To maintain high quality and reliability for `mls`, follow these release best practices.

## 1. Semantic Versioning (SemVer)
We follow [Semantic Versioning 2.0.0](https://semver.org/).
- **MAJOR**: Incompatible API changes or major architectural refactors.
- **MINOR**: Backward-compatible functionality additions.
- **PATCH**: Backward-compatible bug fixes or performance improvements.

## 4. Automated Release & Homebrew Distribution
GoReleaser acts as an automated release bot, handling the entire distribution lifecycle when a new tag is pushed. Here is exactly what happens:

### 1. The Build Phase (in your main repo)
- **Builds Binaries**: It builds the `mls` binary for macOS (ARM/Intel), Linux, and Windows.
- **Archives**: It creates the `.tar.gz` and `.zip` files.
- **Checksums**: It calculates the SHA256 checksums for every single file it just created.

### 2. The Release Phase (in your main repo)
- It creates a new GitHub Release in `MrLeanStorage` using the configured `header`/`footer` templates.
- It uploads the binaries and packages to that release page.

### 3. The Homebrew Phase (the "Magic" part)
GoReleaser performs a Git operation to update your `homebrew-mls` repository:

1. **Authentication**: It uses the `HOMEBREW_TAP_GITHUB_TOKEN` provided in your GitHub Secrets to authenticate with the `homebrew-mls` repository.
2. **Clone**: It performs a shallow clone of your `homebrew-mls` repository.
3. **Templating**: It uses a Formula template (generated via `.goreleaser.yaml`) and plugs in the version and the new SHA256 hashes it just calculated.
4. **Commit & Push**: It writes the updated `mls.rb` file, commits it, and pushes it back to the `homebrew-mls` repository automatically.

### What you see in `homebrew-mls`
After the release finishes, you will see a new commit in the `homebrew-mls` repo:
- **Commit Message**: `chore(mls): update formula to vX.Y.Z`
- **File Changes**: The `mls.rb` file is updated with the new `url` and `sha256` value.

### Why this is brilliant
You never have to open the `homebrew-mls` repo manually.
- You don't have to calculate hashes (which is prone to error).
- You don't have to update version strings manually.
- You don't have to push to the tap.

## 3. Best Practices
- **Never mutate release tags**: If a mistake is found, release a new version (e.g., `v0.1.1`).
- **Clean Main**: Keep the `main` branch always in a releasable state.
- **Documentation**: Ensure `README.md` and `USER_GUIDE.md` are updated before creating a tag.
- **Automated Builds**: Always rely on the CI/CD pipeline for binary distribution to ensure consistency across platforms (macOS, Linux, Windows).
