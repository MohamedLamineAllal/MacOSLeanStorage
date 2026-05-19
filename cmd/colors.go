package cmd

import "github.com/fatih/color"

var (
	// colorTarget is used for displaying target names.
	colorTarget = color.New(color.FgCyan, color.Bold)
	// colorPath is used for displaying filesystem paths.
	colorPath = color.New(color.FgBlue)
	// colorCommand is used for displaying shell commands.
	colorCommand = color.New(color.FgYellow)
	// colorMatch is used for highlighting matched items.
	colorMatch = color.New(color.FgMagenta)
	// colorSuccess is used for success messages and final summaries.
	colorSuccess = color.New(color.FgGreen, color.Bold)
	// colorWarning is used for warnings.
	colorWarning = color.New(color.FgYellow, color.Bold)
	// colorError is used for error messages.
	colorError = color.New(color.FgRed, color.Bold)
	// colorDryRun is used for dry-run specific output.
	colorDryRun = color.New(color.FgHiBlack, color.Italic)
	// colorInfo is used for general informational messages.
	colorInfo = color.New(color.FgCyan)
)
