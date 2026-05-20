# Storage Analysis Report - May 17, 2026

This report documents the findings of a system storage audit performed to refine the cleanup strategy for `MrLeanStorage`.

## Audit Results
The following directories were scanned to identify significant storage consumers on the current development machine:

| Path | Size | Description |
| :--- | :--- | :--- |
| `~/Library/Caches/Google/Chrome` | 312 MB | Browser Cache |
| `~/Library/Application Support/Code/User/workspaceStorage` | 276 MB | VSCode Workspace Metadata |
| `~/Library/Caches/com.microsoft.VSCode` | 3.4 MB | VSCode Global Cache |
| `~/Library/Developer/Xcode/DerivedData` | 20 KB | Xcode Build Artifacts |

## Analysis
- **Browser Caches**: Chrome represents the largest single cache consumer, indicating that the initial strategy to target `~/Library/Caches` is sound.
- **VSCode Workspace Storage**: This is significant (276 MB). Given our policy (Safe after 7 days), this is a prime candidate for automated cleanup.
- **Xcode DerivedData**: Currently minimal, but known to grow rapidly over long development cycles.

## Recommendations
1.  **Prioritize VSCode Workspace Storage**: Ensure the scanner and cleaner are configured to handle `~/Library/Application Support/Code/User/workspaceStorage` as a high-value cleanup target.
2.  **Chromium Profile Support**: Confirm that the scanner logic correctly navigates `User Data` for Chrome to capture profile-specific caches, which may be larger than the global cache folder identified here.
3.  **Xcode Cleanup**: Keep `DerivedData` in the list; it is low-risk and high-reward in terms of reclaiming space during long-term projects.
