package keys

type KeyboardKey interface {
	GetRune() rune
	string() string
}

type symbolChar rune

func (s symbolChar) string() string {
	return string(s)
}

func (s symbolChar) GetRune() rune {
	return rune(s)
}

// .+\t([0-9A-F][0-9A-F])\t0[0-1].+\t(.+)$ conversion regex

type controlChar rune

func (c controlChar) string() string {
	return string(c)
}

func (c controlChar) GetRune() rune {
	return rune(c)
}

// Control chars
const (
	NullChar                controlChar = '\x00'
	StartOfHeading          controlChar = '\x01'
	StartOfText             controlChar = '\x02'
	EndOfText               controlChar = '\x03'
	EndOfTransmission       controlChar = '\x04'
	Enquiry                 controlChar = '\x05'
	Acknowledgment          controlChar = '\x06'
	Bell                    controlChar = '\x07'
	BackSpace               controlChar = '\x08'
	HorizontalTab           controlChar = '\x09'
	LineFeed                controlChar = '\x0A'
	VerticalTab             controlChar = '\x0B'
	FormFeed                controlChar = '\x0C'
	CarriageReturn          controlChar = '\x0D'
	ShiftOutOn              controlChar = '\x0E'
	ShiftInOff              controlChar = '\x0F'
	DataLineEscape          controlChar = '\x10'
	DeviceControl1          controlChar = '\x11'
	DeviceControl2          controlChar = '\x12'
	DeviceControl3          controlChar = '\x13'
	DeviceControl4          controlChar = '\x14'
	NegativeAcknowledgement controlChar = '\x15'
	SynchronousIdle         controlChar = '\x16'
	EndOfTransmitBlock      controlChar = '\x17'
	Cancel                  controlChar = '\x18'
	EndOfMedium             controlChar = '\x19'
	Substitute              controlChar = '\x1A'
	Escape                  controlChar = '\x1B'
	FileSeparator           controlChar = '\x1C'
	GroupSeparator          controlChar = '\x1D'
	RecordSeparator         controlChar = '\x1E'
	UnitSeparator           controlChar = '\x1F'
)

// Symbol chars
const (
	Space                symbolChar = '\x20'
	ExclamationMark      symbolChar = '\x21'
	DoubleQuotes         symbolChar = '\x22'
	Number               symbolChar = '\x23'
	Dollar               symbolChar = '\x24'
	PercentSign          symbolChar = '\x25'
	Ampersand            symbolChar = '\x26'
	SingleQuote          symbolChar = '\x27'
	OpenParenthesis      symbolChar = '\x28'
	CloseParenthesis     symbolChar = '\x29'
	Asterisk             symbolChar = '\x2A'
	Plus                 symbolChar = '\x2B'
	Comma                symbolChar = '\x2C'
	Hyphen               symbolChar = '\x2D'
	Dot                  symbolChar = '\x2E'
	Slash                symbolChar = '\x2F'
	Zero                 symbolChar = '\x30'
	One                  symbolChar = '\x31'
	Two                  symbolChar = '\x32'
	Three                symbolChar = '\x33'
	Four                 symbolChar = '\x34'
	Five                 symbolChar = '\x35'
	Six                  symbolChar = '\x36'
	Seven                symbolChar = '\x37'
	Eight                symbolChar = '\x38'
	Nine                 symbolChar = '\x39'
	Colon                symbolChar = '\x3A'
	Semicolon            symbolChar = '\x3B'
	LessThan             symbolChar = '\x3C'
	Equals               symbolChar = '\x3D'
	GreaterThan          symbolChar = '\x3E'
	QuestionMark         symbolChar = '\x3F'
	AtSymbol             symbolChar = '\x40'
	UppercaseA           symbolChar = '\x41'
	UppercaseB           symbolChar = '\x42'
	UppercaseC           symbolChar = '\x43'
	UppercaseD           symbolChar = '\x44'
	UppercaseE           symbolChar = '\x45'
	UppercaseF           symbolChar = '\x46'
	UppercaseG           symbolChar = '\x47'
	UppercaseH           symbolChar = '\x48'
	UppercaseI           symbolChar = '\x49'
	UppercaseJ           symbolChar = '\x4A'
	UppercaseK           symbolChar = '\x4B'
	UppercaseL           symbolChar = '\x4C'
	UppercaseM           symbolChar = '\x4D'
	UppercaseN           symbolChar = '\x4E'
	UppercaseO           symbolChar = '\x4F'
	UppercaseP           symbolChar = '\x50'
	UppercaseQ           symbolChar = '\x51'
	UppercaseR           symbolChar = '\x52'
	UppercaseS           symbolChar = '\x53'
	UppercaseT           symbolChar = '\x54'
	UppercaseU           symbolChar = '\x55'
	UppercaseV           symbolChar = '\x56'
	UppercaseW           symbolChar = '\x57'
	UppercaseX           symbolChar = '\x58'
	UppercaseY           symbolChar = '\x59'
	UppercaseZ           symbolChar = '\x5A'
	OpeningBracket       symbolChar = '\x5B'
	Backslash            symbolChar = '\x5C'
	ClosingBracket       symbolChar = '\x5D'
	CaretCircumflex      symbolChar = '\x5E'
	Underscore           symbolChar = '\x5F'
	GraveAccent          symbolChar = '\x60'
	LowercaseA           symbolChar = '\x61'
	LowercaseB           symbolChar = '\x62'
	LowercaseC           symbolChar = '\x63'
	LowercaseD           symbolChar = '\x64'
	LowercaseE           symbolChar = '\x65'
	LowercaseF           symbolChar = '\x66'
	LowercaseG           symbolChar = '\x67'
	LowercaseH           symbolChar = '\x68'
	LowercaseI           symbolChar = '\x69'
	LowercaseJ           symbolChar = '\x6A'
	LowercaseK           symbolChar = '\x6B'
	LowercaseL           symbolChar = '\x6C'
	LowercaseM           symbolChar = '\x6D'
	LowercaseN           symbolChar = '\x6E'
	LowercaseO           symbolChar = '\x6F'
	LowercaseP           symbolChar = '\x70'
	LowercaseQ           symbolChar = '\x71'
	LowercaseR           symbolChar = '\x72'
	LowercaseS           symbolChar = '\x73'
	LowercaseT           symbolChar = '\x74'
	LowercaseU           symbolChar = '\x75'
	LowercaseV           symbolChar = '\x76'
	LowercaseW           symbolChar = '\x77'
	LowercaseX           symbolChar = '\x78'
	LowercaseY           symbolChar = '\x79'
	LowercaseZ           symbolChar = '\x7A'
	OpeningBrace         symbolChar = '\x7B'
	VerticalBar          symbolChar = '\x7C'
	ClosingBrace         symbolChar = '\x7D'
	EquivalencySignTilde symbolChar = '\x7E'
	Delete               symbolChar = '\x7F'
)
