package pkpass

type BarcodeFormat string

const (
	PKBarcodeFormatQR      = "PKBarcodeFormatQR"
	PKBarcodeFormatPDF417  = "PKBarcodeFormatPDF417"
	PKBarcodeFormatAztec   = "PKBarcodeFormatAztec"
	PKBarcodeFormatCode128 = "PKBarcodeFormatCode128"
)

type BarcodeDict struct {
	AltText         string        `json:"altText"`
	Format          BarcodeFormat `json:"format"`
	Message         string        `json:"message"`
	MessageEncoding string        `json:"messageEncoding"`
}
