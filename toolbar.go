package main

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolbar() *widget.Toolbar {
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), app.refreshPriceContent),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)

	return toolBar
}
