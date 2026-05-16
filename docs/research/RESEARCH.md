# Research: macOS Storage Cleanup Analysis

## 1. High-Storage Consumer Locations

### Browser Caches (Chromium-based: Arc, Chrome, Edge)
Chromium-based browsers are notorious for "Service Worker" caches and standard HTTP caches.

*   **Arc Browser**:
    *   **Main Cache**: `~/Library/Caches/company.thebrowser.Browser/`
    *   **User Data**: `~/Library/Application Support/Arc/User Data/`
        *   Profiles are stored in subdirectories: `Default`, `Profile 1`, `Profile 2`, etc.
        *   Each profile has:
            *   `Cache/`
            *   `Code Cache/`
            *   `Service Worker/CacheStorage/`
*   **Google Chrome**: `~/Library/Caches/Google/Chrome/` and `~/Library/Application Support/Google/Chrome/` (similar profile structure).

### Developer Tools
*   **VSCode**:
    *   **Cached Data**: `~/Library/Application Support/Code/CachedData` (transpiled JS).
    *   **Workspace Storage**: `~/Library/Application Support/Code/User/workspaceStorage` (UI state per project).
    *   **Logs**: `~/Library/Application Support/Code/logs`.
*   **Xcode**:
    *   **DerivedData**: `~/Library/Developer/Xcode/DerivedData` (safe to delete, will rebuild).
    *   **Archives**: `~/Library/Developer/Xcode/Archives` (contains build history, large).
*   **Package Managers**:
    *   **NPM**: `~/.npm`
    *   **Yarn**: `~/Library/Caches/Yarn`
    *   **Bun**: `~/.bun/install/cache`
    *   **Go**: `~/Library/Caches/go-build`

### System Locations
*   **System Caches**: `~/Library/Caches`
*   **Logs**: `~/Library/Logs` and `/var/log`
*   **Temporary Items**: `/private/var/folders` (managed by macOS, but can be cleaned).

## 2. Multi-Profile Handling Logic
To ensure we clean all profiles for apps like Arc or VSCode:
1.  Identify the base directory (e.g., `~/Library/Application Support/Arc/User Data/`).
2.  Iterate through all subdirectories.
3.  Check if the subdirectory contains a `Cache` or `Service Worker` folder.
4.  Apply cleanup rules to each.

## 3. Safe Removal Logic & Metrics

### Metric: Access Time (`atime`) vs. Modification Time (`mtime`)
*   **`atime`** is the best indicator of whether a cache is still "useful". If a cache file hasn't been accessed in 7 days, it's likely safe to remove.
*   **`mtime`** can be misleading for caches that are read frequently but rarely updated.

### Safety Thresholds
*   **100% Safe (Anytime)**:
    *   `~/Library/Application Support/Code/CachedData`
    *   `~/Library/Developer/Xcode/DerivedData`
    *   `~/Library/Caches/company.thebrowser.Browser/Default/Cache` (standard HTTP cache)
*   **Safe after 3 Days**:
    *   `/private/var/folders` (system temp)
    *   `~/Library/Logs`
*   **Safe after 7 Days**:
    *   `Service Worker/CacheStorage` (Browser caches)
    *   `workspaceStorage` (VSCode - will reset UI state like open files/collapsed folders)

## 4. Automation Strategy
*   **Scanning**: Use a concurrent walker to map sizes and last-access times.
*   **Cleaning**: Filter the map based on safety rules and user-defined thresholds.
*   **Scheduling**: A background agent (Go binary) that runs a light check every 24 hours.
