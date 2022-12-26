package main

import (
	"fmt"
	"image/png"
	"os"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {

	fmt.Println("Qr Form Tamu Privy")

	// Create the barcode
	qrCode, _ := qr.Encode("www.bit.ly/FormTamuPrivy", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("file/qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}