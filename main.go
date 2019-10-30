package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	svg "github.com/ajstarks/svgo"
)

type Status struct {
	Resource string
	Status   string
	Details  string
}

func main() {
	Radiate()
}

func Radiate() {
	data := readData()
	resWidth := 2560
	resHeight := 1600
	totalBoxes := len(data)
	width := resWidth/totalBoxes + 50
	height := resHeight/totalBoxes + 50

	for _, status := range data {
		RenderResourceStatus(status, width, height)
	}
}

func readData() []Status {
	plan, _ := ioutil.ReadFile("status.json")
	var data []Status
	err := json.Unmarshal(plan, &data)
	if err != nil {
		log.Fatalf("failed to read data: %v", err)
	}
	return data
}

func RenderResourceStatus(status Status, width int, height int) {
	canvas := svg.New(os.Stdout)
	style := "font-size:12pt;text-anchor:middle"

	canvas.Start(width, height)
	if status.Status == "PASS" {
		canvas.Rect(0, 0, width, height, "fill:green")
	} else {
		canvas.Rect(0, 0, width, height, "fill:red")
	}
	resourceName := fmt.Sprintf("%s", status.Resource)
	resourceDetails := fmt.Sprintf("(%s)", status.Details)
	canvas.Text(width/2, height*3/5, resourceName, style)
	canvas.Text(width/2, height*3/4, resourceDetails, style)
	canvas.End()
}
