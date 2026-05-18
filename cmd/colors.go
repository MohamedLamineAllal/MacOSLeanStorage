package cmd

import "github.com/fatih/color"

var (
	colorTarget  = color.New(color.FgCyan, color.Bold)
	colorPath    = color.New(color.FgBlue)
	colorCommand = color.New(color.FgYellow)
	colorMatch   = color.New(color.FgMagenta)
	colorSuccess = color.New(color.FgGreen, color.Bold)
	colorWarning = color.New(color.FgYellow, color.Bold)
	colorError   = color.New(color.FgRed, color.Bold)
	colorDryRun  = color.New(color.FgHiBlack, color.Italic)
	colorInfo    = color.New(color.FgCyan)
)
