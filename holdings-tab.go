package main

import (
	"fmt"
	"goldwatcher/repository"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) holdingsTab() *fyne.Container {
	app.HoldingsTable = app.getHoldingsTable()

	holdingsContainer := container.NewVBox(app.HoldingsTable)

	return holdingsContainer
}

func (app *Config) getHoldingsTable() *widget.Table {
	data := app.getHoldingSlice()
	app.Holdings = data
	t := widget.NewTable(
		func() (rows int, cols int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == (len(data[0])-1) && i.Row != 0 {
				// last cell - put in a button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "", func(deleted bool) {
						id, _ := strconv.Atoi(data[i.Row][0].(string)) // Cast to string
						err := app.DB.DeleteHolding(int64(id))
						if err != nil {
							app.ErrorLog.Println(err)
						}

						// Refresh holding table
						app.refreshHoldingsTable()
					}, app.MainWindow)
				})

				w.Importance = widget.HighImportance

				o.(*fyne.Container).Objects = []fyne.CanvasObject{w}
			} else {
				// we're just putting in textual information
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(data[i.Row][i.Col].(string)),
				} // Casting
			}
		},
	)

	colWidths := []float32{50, 200, 200, 200, 100}

	for i := 0; i < len(data[0]); i++ {
		t.SetColumnWidth(i, colWidths[i])
	}

	return t
}

func (app *Config) getHoldingSlice() [][]interface{} {
	var slice [][]interface{}

	holdings, err := app.currentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	slice = append(slice, []interface{}{"ID", "Amount", "Purchase", "Price", "Delete?"})

	for _, x := range holdings {
		var currentRow []interface{}

		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%d toz", x.Amount))
		currentRow = append(currentRow, fmt.Sprintf("%2f", float32(x.PurchasePrice)/100))
		currentRow = append(currentRow, x.PurchaseDate.Format("2006-01-02"))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

	return slice
}

func (app *Config) currentHoldings() ([]repository.Holdings, error) {
	holdings, err := app.DB.AllHoldings()
	if err != nil {
		app.ErrorLog.Println()
		return nil, err
	}

	return holdings, nil
}

func (app *Config) refreshHoldingsTable() {
	app.Holdings = app.getHoldingSlice()
	app.HoldingsTable.Refresh()
}
