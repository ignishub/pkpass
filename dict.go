package pkpass

type DataDetectorType string

const (
	PKDataDetectorTypePhoneNumber   = DataDetectorType("PKDataDetectorTypePhoneNumber")
	PKDataDetectorTypeLink          = DataDetectorType("PKDataDetectorTypeLink")
	PKDataDetectorTypeAddress       = DataDetectorType("PKDataDetectorTypeAddress")
	PKDataDetectorTypeCalendarEvent = DataDetectorType("PKDataDetectorTypeCalendarEvent")
)

type TextAlignment string

const (
	PKTextAlignmentLeft    = TextAlignment("PKTextAlignmentLeft")
	PKTextAlignmentCenter  = TextAlignment("PKTextAlignmentCenter")
	PKTextAlignmentRight   = TextAlignment("PKTextAlignmentRight")
	PKTextAlignmentNatural = TextAlignment("PKTextAlignmentNatural")
)

type StandartFieldDict struct {
	AttributedValue   string             `json:"attributedValue,omitempty"`
	ChangeMessage     string             `json:"changeMessage,omitempty"`
	DataDetectorTypes []DataDetectorType `json:"dataDetectorTypes,omitempty"`
	Key               string             `json:"key"`
	Label             string             `json:"label,omitempty"`
	TextAlignment     TextAlignment      `json:"textAlignment,omitempty"`
	Value             interface{}        `json:"value"`
}
