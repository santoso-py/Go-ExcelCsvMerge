package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

func main() {
	a := app.New()
	w := a.NewWindow("Excel & CSV Merger")
	w.Resize(fyne.NewSize(500, 200))

	inputPath := widget.NewEntry()
	outputPath := widget.NewEntry()

	btnSelectInput := widget.NewButton("Pilih Folder Input", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri != nil {
				inputPath.SetText(uri.Path())
			}
		}, w)
	})

	btnSelectOutput := widget.NewButton("Pilih Folder Output", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri != nil {
				outputPath.SetText(uri.Path())
			}
		}, w)
	})

	btnMerge := widget.NewButton("Merge Files", func() {
		inputFolder := inputPath.Text
		outputFolder := outputPath.Text

		if inputFolder == "" || outputFolder == "" {
			dialog.ShowError(fmt.Errorf("Pilih folder input dan output terlebih dahulu"), w)
			return
		}

		files, err := os.ReadDir(inputFolder)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Gagal membaca folder input: %v", err), w)
			return
		}

		mergedData := [][]string{}
		headerAdded := false

		for _, file := range files {
			filePath := filepath.Join(inputFolder, file.Name())

			if strings.HasSuffix(file.Name(), ".csv") {
				f, err := os.Open(filePath)
				if err != nil {
					dialog.ShowError(fmt.Errorf("Gagal membuka file CSV: %v", err), w)
					return
				}
				defer f.Close()

				reader := csv.NewReader(f)
				rows, err := reader.ReadAll()
				if err != nil {
					dialog.ShowError(fmt.Errorf("Gagal membaca file CSV: %v", err), w)
					return
				}

				if !headerAdded {
					mergedData = append(mergedData, rows[0])
					headerAdded = true
				}
				mergedData = append(mergedData, rows[1:]...)
			} else if strings.HasSuffix(file.Name(), ".xlsx") {
				f, err := excelize.OpenFile(filePath)
				if err != nil {
					dialog.ShowError(fmt.Errorf("Gagal membuka file Excel: %v", err), w)
					return
				}

				sheetName := f.GetSheetName(0)
				rows, err := f.GetRows(sheetName)
				if err != nil {
					dialog.ShowError(fmt.Errorf("Gagal membaca file Excel: %v", err), w)
					return
				}

				if !headerAdded {
					mergedData = append(mergedData, rows[0])
					headerAdded = true
				}
				mergedData = append(mergedData, rows[1:]...)
			}
		}

		outputFilePath := filepath.Join(outputFolder, "combined.csv")
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Gagal membuat file output: %v", err), w)
			return
		}
		defer outputFile.Close()

		writer := csv.NewWriter(outputFile)
		err = writer.WriteAll(mergedData)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Gagal menulis file CSV: %v", err), w)
			return
		}
		writer.Flush()

		dialog.ShowInformation("Sukses", "File berhasil digabung dan disimpan di "+outputFilePath, w)
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Pilih Folder Input:"),
		container.NewHBox(inputPath, btnSelectInput),
		widget.NewLabel("Pilih Folder Output:"),
		container.NewHBox(outputPath, btnSelectOutput),
		btnMerge,
	))

	w.ShowAndRun()
}
