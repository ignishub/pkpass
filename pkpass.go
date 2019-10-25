package pkpass

type PKPassImage struct {
}

type PKPass struct {
	Manifest Manifest
	Pass     Pass
	// Images
	Background PKPassImage
	Logo       PKPassImage
	Icon       PKPassImage
	Footer     PKPassImage
	Strip      PKPassImage
	Thumbnail  PKPassImage
}
