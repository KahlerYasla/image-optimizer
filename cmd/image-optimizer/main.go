package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/chai2010/webp"
	"golang.org/x/image/draw"
)

func main() {
	// Get the directory from user input
	var dir string
	fmt.Print("Enter the directory containing images: ")
	fmt.Scan(&dir)

	// Read the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Process each file
	for i, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			if ext == ".jpeg" || ext == ".jpg" || ext == ".png" || ext == ".webp" {
				inputPath := filepath.Join(dir, file.Name())
				outputPath := filepath.Join(dir, file.Name()[:len(file.Name())-len(ext)]+".webp")

				if err := convertToWebP(inputPath, outputPath); err != nil {
					fmt.Println("Error converting file:", file.Name(), "-", err)
				} else {
					fmt.Printf("Processed %d of %d: %s -> %s\n", i+1, len(files), file.Name(), outputPath)
				}
			}
		}
	}
}

// convertToWebP converts an image to the WebP format and saves it to the specified output path
func convertToWebP(inputPath, outputPath string) error {
	// Open the input image file
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the image based on its format
	var img image.Image
	switch filepath.Ext(inputPath) {
	case ".jpeg", ".jpg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	case ".webp":
		img, err = webp.Decode(file)
	default:
		return fmt.Errorf("unsupported file format")
	}
	if err != nil {
		return err
	}

	// Resize the image to reduce file size (optional step)
	dst := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	draw.BiLinear.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Encode the image in WebP format with default options
	return webp.Encode(outFile, dst, &webp.Options{Quality: 75})
}
