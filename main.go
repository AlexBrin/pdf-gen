package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	r := NewRequestPdf("")

	//html template path
	templatePath := "template/sample.html"

	//path for download pdf
	outputPath := "out/example.pdf"

	imageFile, _ := os.Open("template/image/icon.png")
	imageFileStat, _ := imageFile.Stat()

	data := make([]byte, imageFileStat.Size())
	_, _ = imageFile.Read(data)

	img := base64.StdEncoding.EncodeToString(data)

	type d struct {
		Image       string
		Title       string
		Description string
		Company     string
		Contact     string
		Country     string
	}

	//html template data
	templateData := []d{}
	for i := 0; i < 100; i++ {
		templateData = append(templateData, d{
			Image:       img,
			Title:       "HTML to PDF generator",
			Description: "This is the simple HTML to PDF file.",
			Company:     "Jhon Lewis",
			Contact:     "Maria Anders",
			Country:     "Germany",
		})
	}

	if err := r.ParseTemplate(templatePath, map[string]any{"data": templateData}); err == nil {
		ok, _ := r.GeneratePDF(outputPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
