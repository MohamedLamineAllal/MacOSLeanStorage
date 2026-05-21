//go:build windows

package defaults

func GetDefaultConfig() string {
	return `targets:
  # System/User Caches
  - name: "User Local Caches"
    path: "%LocalAppData%\\Temp\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"

  # Arc Browser (Windows uses standard Chromium structures here)
  - name: "Arc CacheStorage Cache"
    path: "%LocalAppData%\\Arc\\User Data\\**\\CacheStorage"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Arc CacheStorage Cache sub"
    path: "%LocalAppData%\\Arc\\User Data\\**\\CacheStorage\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Arc File System"
    path: "%LocalAppData%\\Arc\\User Data\\**\\File System"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Arc File System sub"
    path: "%LocalAppData%\\Arc\\User Data\\**\\File System\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Arc File IndexedDB"
    path: "%LocalAppData%\\Arc\\User Data\\**\\IndexedDB\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"

  # Google Chrome
  - name: "Chrome Global Cache"
    path: "%LocalAppData%\\Google\\Chrome\\User Data\\Default\\Cache\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Chrome CacheStorage Cache"
    path: "%LocalAppData%\\Google\\Chrome\\User Data\\**\\CacheStorage\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Chrome crx Cache"
    path: "%LocalAppData%\\Google\\Chrome\\User Data\\**\\Web Applications\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Chrome File System"
    path: "%LocalAppData%\\Google\\Chrome\\User Data\\**\\File System\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Chrome File IndexedDB"
    path: "%LocalAppData%\\Google\\Chrome\\User Data\\**\\IndexedDB\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"

  # Communication Tools
  - name: "Discord Cache"
    path: "%AppData%\\discord\\Cache\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Slack CacheStorage"
    path: "%LocalAppData%\\Slack\\**\\CacheStorage\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Microsoft Teams CacheStorage"
    path: "%LocalAppData%\\Packages\\MSTeams_8wekyb3d8bbwe\\LocalCache\\Microsoft\\MSTeams\\**\\CacheStorage\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"

  # Development Tools
  - name: "Cursor Cache"
    path: "%AppData%\\Cursor\\Cache\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Cursor Cache/Cache_Data"
    path: "%AppData%\\Cursor\\Cache\\Cache_Data\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Cursor CachedStorage"
    path: "%AppData%\\Cursor\\Service Worker\\CacheStorage\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Cursor CachedExtensionVSIXs"
    path: "%AppData%\\Cursor\\CachedExtensionVSIXs\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "VSCode CachedData"
    path: "%AppData%\\Code\\CachedData\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "VSCode Cache/Cache_Data"
    path: "%AppData%\\Code\\Cache\\Cache_Data\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "VSCode CachedStorage"
    path: "%AppData%\\Code\\Service Worker\\CacheStorage\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "VSCode WebStorage CacheStorage"
    path: "%AppData%\\Code\\WebStorage\\**\\CacheStorage\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "VSCode CachedExtensionVSIXs"
    path: "%AppData%\\Code\\CachedExtensionVSIXs\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"

  # AI & Other
  - name: "OpenAI Atlas Cache"
    path: "%AppData%\\com.openai.atlas\\Cache\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "OpenAI Atlas CachedStorage"
    path: "%AppData%\\com.openai.atlas\\Service Worker\\CacheStorage\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "OpenAI Atlas File System"
    path: "%AppData%\\com.openai.atlas\\File System\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Telegram Desktop Media Cache"
    path: "%AppData%\\Telegram Desktop\\tdata\\user_data\\media_cache\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Telegram Desktop Cache"
    path: "%AppData%\\Telegram Desktop\\tdata\\user_data\\cache\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Figma Local Storage"
    path: "%AppData%\\Figma\\Local Storage\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Figma Local Cache"
    path: "%LocalAppData%\\Figma\\**\\Cache\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Spotify Cache"
    path: "%LocalAppData%\\Spotify\\Storage\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"

  # System/Build Tools
  - name: "Go Build Cache"
    path: "%LocalAppData%\\go-build\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Scoop Cache" # Replacing Homebrew with Windows equivalent Scoop
    path: "%LocalAppData%\\scoop\\apps\\scoop\\current\\cache\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "npm/node-gyp"
    path: "%AppData%\\npm-cache\\_logs\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "pip Cache"
    path: "%LocalAppData%\\pip\\Cache\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "pnpm Cache"
    path: "%LocalAppData%\\pnpm\\Cache\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"
  - name: "Microsoft OneNote notes backup"
    path: "%LocalAppData%\\Microsoft\\OneNote\\*\\Backup\\**"
    threshold_days: 30
    safety_level: 1
    type: "both"

  # Commands
  - name: "PNPM Store Prune"
    command: "pnpm store prune"
    interval_days: 30
    safety_level: 1
  - name: "npm clean cache"
    command: "npm cache clean --force"
    interval_days: 30
    safety_level: 1

dry_run: true
ignore_patterns:
  - "desktop.ini"
  - "thumbs.db"
  - "$RECYCLE.BIN"
schedule: "0 0 0 * * *"
`
}
