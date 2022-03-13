package utilities

import "github.com/pterm/pterm"

func StylePrint(a, b string) {
	err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle(a, pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle(b, pterm.NewStyle(pterm.FgLightMagenta))).
		Render()
	Error(err, "Failed to render Pterm")
}
