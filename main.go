package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	r := NewRequestPdf("")

	//html template path
	templatePath := "templates/sample.html"

	//path for download pdf
	outputPath := "storage/example.pdf"

	type d struct {
		Image   string
		Company string
		Contact string
		Country string
	}

	imgBytes, _ := ioutil.ReadFile("images/icon.png")
	img := base64.StdEncoding.EncodeToString(imgBytes)

	//html template data
	templateData := make([]d, 0, 100)

	for i := 0; i < 100; i++ {
		templateData = append(templateData, d{
			Image:   img,
			Company: "Jhon Lewis",
			Contact: "Maria Anders",
			Country: "Germany",
		})
	}

	if err := r.ParseTemplate(templatePath, map[string]any{
		"data": templateData,
	}); err == nil {
		ok, _ := r.GeneratePDF(outputPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
