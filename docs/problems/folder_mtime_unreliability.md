# Problem: Unreliability of Directory Modification Time (mtime) for Staleness Checks

## Description
The current implementation relies on directory modification time (`mtime`) to determine if a directory is "stale" and eligible for removal when using the `folder` cleanup mode.

## The Issue
On macOS (and many POSIX filesystems), the `mtime` of a directory is updated whenever an entry inside the directory is created, deleted, or renamed. More importantly, simply browsing the folder in Finder often triggers metadata updates or access-time updates that are reflected in the parent's directory attributes. 

As a result, a directory that contains truly stale data (files not touched for months) may have a very recent `mtime` simply because the user opened the folder in Finder to look at its contents, rendering the staleness check ineffective.

## Impact
- **False Negatives**: Stale directories are not cleaned up because they appear "active" due to recent access or metadata updates.
- **Inconsistent Behavior**: Cleanup depends on user interaction (e.g., browsing) rather than the actual content's age, making the cleanup unpredictable.

## Proposed Solution
Instead of relying on the directory's `mtime`, the cleaner should evaluate staleness based on the **oldest file within the directory structure** (or the newest file, depending on the policy).

1.  **Deep Staleness Check**: For folders, the scanner should perform a recursive check of the contents.
2.  **Aggregation**: A directory is considered stale only if **none** of its contents have been modified within the `threshold_days`.
3.  **Metadata Strategy**: Optionally, consider using the `ctime` (change time) or a combination of `mtime` across all files to better determine if a directory tree has been meaningfully updated.

This approach ensures that cleanup is driven by the actual contents, not by how or when a user viewed the directory.
