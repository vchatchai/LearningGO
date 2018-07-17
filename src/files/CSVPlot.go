package files

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Arafatk/glot"
)

func CSVPlot(file string) (err error) {
	_, err = os.Stat(file)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()
	if err != nil {
		return
	}
	xP := []float64{}
	yP := []float64{}
	for _, rec := range allRecords {
		x, _ := strconv.ParseFloat(rec[0], 64)
		y, _ := strconv.ParseFloat(rec[1], 64)
		xP = append(xP, x)
		yP = append(yP, y)
	}
	points := [][]float64{}
	points = append(points, xP)
	points = append(points, yP)
	fmt.Println(points)

	dimensios := 2
	persist := true
	debug := false

	plot, _ := glot.NewPlot(dimensios, persist, debug)

	plot.SetTitle("Using Glot with CSV data")
	plot.SetXLabel("X-Axis")
	plot.SetYLabel("Y-Axis")
	style := "circle"
	plot.AddPointGroup("Circle:", style, points)
	plot.SavePlot("output.png")
	return
}
