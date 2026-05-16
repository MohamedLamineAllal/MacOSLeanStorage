# Comprehensive Application Storage Analysis

This document provides a detailed breakdown of storage consumption patterns for various applications on macOS and defines the safe cleanup strategy.

## 1. Chromium-Based Browsers (Arc, Chrome, Brave, etc.)

### Storage Pattern
Chromium uses a profile-based architecture. All user data, including caches, is partitioned by profile.

- **Main Cache (Global)**: `~/Library/Caches/<bundle_id>/`
- **User Data (Profile-Specific)**: `~/Library/Application Support/<app_name>/User Data/`
    - Folders like `Default`, `Profile 1`, `Profile 2` contain:
        - `Cache/`: Standard HTTP cache.
        - `Code Cache/`: Compiled JavaScript.
        - `Service Worker/CacheStorage/`: Often the largest "hidden" consumer.
        - `GPUCache/`: Browser-specific GPU state.

### Multi-Profile Logic
To handle multiple profiles, the tool must:
1.  Locate the `User Data` directory.
2.  Iterate through subdirectories (excluding system ones like `SwiftShader`, `GrShaderCache`).
3.  Detect the presence of a `Cache` or `Service Worker` directory.
4.  Calculate size and apply staleness filters (`atime`).

### Specific Apps
- **Arc**: `~/Library/Application Support/Arc/User Data/`
- **Google Chrome**: `~/Library/Application Support/Google/Chrome/`

---

## 2. Developer Tools (VSCode, Cursor, Xcode)

### VSCode & Cursor
Both are Electron-based and share similar structures.
- **Cached Data**: `~/Library/Application Support/<App>/CachedData` (Safe anytime).
- **Workspace Storage**: `~/Library/Application Support/<App>/User/workspaceStorage` (Safe after 7 days - resets UI state).
- **Global Cache**: `~/Library/Caches/com.microsoft.VSCode` or `~/Library/Caches/com.todesktop.cursor`.

### Xcode
- **DerivedData**: `~/Library/Developer/Xcode/DerivedData` (Safe anytime - will recompile).
- **Archives**: `~/Library/Developer/Xcode/Archives` (Safe after 30 days or based on project activity).
- **iOS DeviceSupport**: `~/Library/Developer/Xcode/iOS DeviceSupport` (Safe to remove old OS versions).

---

## 3. Communication & Electron Apps (Discord, Slack)

### Discord
- **Path**: `~/Library/Application Support/discord/`
- **Clean Targets**: `Cache`, `Code Cache`, `GPUCache`.
- **Logic**: Safe to remove anytime if Discord is closed.

---

## 4. AI & Specialized Tools (Antigravity, OpenAI Atlas)

### Antigravity
- **Application Support**: `~/Library/Application Support/Antigravity/`
- **Transient Cache**: `~/Library/Caches/Antigravity/`
- **Internal Memory**: `~/.gemini/` (Use caution - contains conversation logs).

### OpenAI Atlas
- **Path**: `~/Library/Application Support/OpenAI/Atlas/`
- **Cache**: `~/Library/Caches/com.openai.atlas/`
- **Logic**: Cleaning `Caches` is safe. `Application Support` contains SQLite databases for memory/history.

---

## 5. Decision Matrix: Logic & Time Constraints

| Application Type | Target Folder | Safety | Time Constraint (`atime`) |
| :--- | :--- | :--- | :--- |
| **Browsers** | `Cache` | 100% | > 7 days |
| **Browsers** | `Service Worker` | 100% | > 14 days |
| **Editors** | `CachedData` | 100% | > 1 day |
| **Editors** | `workspaceStorage` | Safe | > 7 days |
| **Xcode** | `DerivedData` | 100% | > 3 days |
| **General** | `GPUCache` | 100% | Anytime |
| **System** | `/private/var/folders` | Safe | > 3 days |

## Multi-Profile Discovery Implementation
The Go implementation should use `os.ReadDir` on the base `User Data` path and filter for directories containing the target cleanup sub-folders. This ensures that even "guest" or "temporary" profiles are cleaned.
