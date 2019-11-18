package pkpass

const PKFormatVersion = 1

type Pass struct {
	// Standart Keys
	Description        string `json:"description"`
	FormatVersion      int    `json:"formatVersion"`
	OrganizationName   string `json:"organizationName"`
	PassTypeIdentifier string `json:"passTypeIdentifier"`
	SerialNumber       string `json:"serialNumber"`
	TeamIdentifier     string `json:"teamIdentifier"`

	// Associated App Keys
	AppLaunchURL              string `json:"appLaunchURL,omitempty"`
	AssociatedStoreIdentifier []int  `json:"associatedStoreIdentifier,omitempty"`

	// Companion App Keys
	UserInfo interface{} `json:"userInfo,omitempty"`

	// Expiration Keys
	ExpirateionDate string `json:"expirateionDate,omitempty"`
	Voided          bool   `json:"voided,omitempty"`

	// TODO: Relevance Keys
	// Beacons
	// Locations []Location
	// MaxDistance int
	// RelevantDate string

	// TODO: Style Keys

	// Visual Appearance Keys
	Barcode            *BarcodeDict  `json:"barcode,omitempty"`
	Barcodes           []BarcodeDict `json:"barcodes,omitempty"`
	BackgroundColor    Color         `json:"backgroundColor,omitempty"`
	ForegroundColor    Color         `json:"foregroundColor,omitempty"`
	GroupingIdentifier string        `json:"groupingIdentifier,omitempty"`
	LabelColor         Color         `json:"labelColor,omitempty"`
	LogoText           string        `json:"logoText,omitempty"`
	SuppressStripShine bool          `json:"suppressStripShine,omitempty"`

	// Web Service Keys
	AuthenticationToken string `json:"authenticationToken,omitempty"`
	WebServiceURL       string `json:"webServiceURL,omitempty"`

	// NFC-Enabled Pass Keys
	NFC *NFCDict `json:"nfc,omitempty"`

	StoreCard map[string]interface{} `json:"storeCard,omitempty"`
}
