# Comprehensive Safe Cleanup Inventory

This document identifies major macOS storage consumers that are safe for automated cleanup. Safety is categorized by the nature of the data (Cache vs. Derived) and age requirements.

## Safety Categories
- **Level 1 (Safe Anytime)**: Rebuildable data (caches, compiled outputs, GPU caches).
- **Level 2 (Safe after 7+ Days)**: Session-specific data, logs, workspace metadata.
- **Level 3 (Caution)**: SQLite databases, history, local configuration.

## Inventory Table

| Application | Target Path | Safety Level | Notes |
| :--- | :--- | :--- | :--- |
| **Arc** | `~/Library/Application Support/Arc/User Data/.../Cache` | L1 | Browser cache |
| **Chrome** | `~/Library/Caches/Google/Chrome` | L1 | Global browser cache |
| **Discord** | `~/Library/Application Support/discord/Cache` | L1 | Electron cache |
| **Cursor** | `~/Library/Application Support/Cursor/CachedData` | L1 | Editor cache |
| **VSCode** | `~/Library/Application Support/Code/CachedData` | L1 | Editor cache |
| **VSCode** | `~/Library/Application Support/Code/User/workspaceStorage`| L2 | Metadata (safe > 7d) |
| **OpenAI Atlas** | `~/Library/Caches/com.openai.atlas` | L1 | App cache |
| **Telegram** | `~/Library/Caches/ru.keepcoder.Telegram` | L1 | Media/msg cache |
| **Figma** | `~/Library/Application Support/Figma/Local Storage` | L1 | Electron cache |
| **Spotify** | `~/Library/Caches/com.spotify.client` | L1 | Music cache |
| **Go** | `~/Library/Caches/go-build` | L1 | Build artifacts |
| **Homebrew** | `~/Library/Caches/Homebrew` | L1 | Downloaded formulas |
| **npm/node** | `~/Library/Caches/node-gyp` | L1 | Build artifacts |
| **pip/pnpm** | `~/Library/Caches/pip` / `pnpm` | L1 | Downloaded packages |

## Implementation Strategy
1. **Automated Scanner**: Use this list to populate `internal/config/config.go`.
2. **Dynamic Traversal**: For apps like Arc/Chrome, the scanner must iterate through `User Data` directories to target caches in *all* profiles, not just the default one.
3. **Threshold Enforcement**:
    - **L1**: Apply 3-day staleness threshold to allow for frequent access.
    - **L2**: Enforce 7-day threshold to maintain session continuity.
4. **Safety Check**: Always dry-run by default.
