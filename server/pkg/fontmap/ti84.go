package fontmap

type Font struct {
	Name     string       `json:"name"`
	Metadata MetadataType `json:"metadata"`
	Charmap  CharmapType  `json:"charmap"`
}

type MetadataType struct {
	TallerCharacters map[string]int `json:"tallerCharacters"`
	AverageHeight    int            `json:"averageHeight"`
	AverageWidth     int            `json:"averageWidth"`
}

type CharmapType map[string]Letter
type Letter []Row
type Row []string

// Hard coded stuff for this font
var TI84 Font = Font{
	Name: "TI84",
	Metadata: MetadataType{
		AverageHeight: 5,
		AverageWidth:  4,
		TallerCharacters: map[string]int{
			"Q": 6,
			"g": 6,
			"j": 6,
			"p": 6,
			"q": 6,
			"y": 6,
		},
	},
	Charmap: CharmapType{
		"A": Letter{
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"B": Letter{
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
		},
		"C": Letter{
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
		},
		"D": Letter{
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
		},
		"E": Letter{
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"F": Letter{
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
		},
		"G": Letter{
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
		},
		"H": Letter{
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"I": Letter{
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"J": Letter{
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"K": Letter{
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"L": Letter{
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"M": Letter{
			Row{"⚫️", "⚪️", "⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"N": Letter{
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"O": Letter{
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"P": Letter{
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
		},
		"Q": Letter{
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
		},
		"R": Letter{
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"S": Letter{
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
		},
		"T": Letter{
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"U": Letter{
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"V": Letter{
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"W": Letter{
			Row{"⚫️", "⚪️", "⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️", "⚫️", "⚪️"},
		},
		"X": Letter{
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"Y": Letter{
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"Z": Letter{
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"a": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
		},
		"b": Letter{
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
		},
		"c": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
		},
		"d": Letter{
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
		},
		"e": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
		},
		"f": Letter{
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
		},
		"g": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
		},
		"h": Letter{
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"i": Letter{
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
		},
		"j": Letter{
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"k": Letter{
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"l": Letter{
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
		},
		"m": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️", "⚫️", "⚪️"},
		},
		"n": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"o": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"p": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
		},
		"q": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚫️", "⚪️"},
		},
		"r": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
		},
		"s": Letter{
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️"},
		},
		"t": Letter{
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
		},
		"u": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"v": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"w": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"x": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
		},
		"y": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
		},
		"z": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"0": Letter{
			Row{"⚫️️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"1": Letter{
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"2": Letter{
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"3": Letter{
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
		},
		"4": Letter{
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪", "⚫️", "⚪️"},
			Row{"⚫️", "⚫", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
		},
		"5": Letter{
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
		},
		"6": Letter{
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"7": Letter{
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
		},
		"8": Letter{
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"9": Letter{
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
		},
		".": Letter{
			Row{"⚪️", "⚪️"},
			Row{"⚪️", "⚪️"},
			Row{"⚪️", "⚪️"},
			Row{"⚪️", "⚪️"},
			Row{"⚫️", "⚪️"},
		},
		",": Letter{
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
		},
		"\"": Letter{
			Row{"⚫️", "⚪️", "⚫️"},
			Row{"⚫️", "⚪️", "⚫️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
		},
		"'": Letter{
			Row{"⚫️", "⚫️"},
			Row{"⚪️", "⚫️"},
			Row{"⚪️", "⚪️"},
			Row{"⚪️", "⚪️"},
			Row{"⚪️", "⚪️"},
		},
		"?": Letter{
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"!": Letter{
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
		},
		"@": Letter{
			Row{"⚪️", "⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️", "⚫️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️", "⚫️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️"},
			Row{"⚪️", "⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"_": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚫️", "⚪️"},
		},
		"*": Letter{
			Row{"⚪️", "⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️"},
			Row{"⚪️", "⚫️", "⚫️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚫️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"#": Letter{
			Row{"⚪️","⚫️","⚪️","⚫️","⚪️"},
			Row{"⚫️","⚫️","⚫️","⚫️","⚫️"},
			Row{"⚪️","⚫️","⚪️","⚫️","⚪️"},
			Row{"⚫️","⚫️","⚫️","⚫️","⚫️"},
			Row{"⚪️","⚫️","⚪️","⚫️","⚪️"},
		},
		"$": Letter{
			Row{"⚪️", "⚫️", "⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️", "⚫️", "⚪️", "⚪️"},
		},
		"%": Letter{
			Row{"⚫️", "⚪️", "⚫️"},
			Row{"⚪️", "⚪️", "⚫️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️"},
		},
		"&": Letter{
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚫️"},
		},
		"(": Letter{
			Row{"⚪️", "⚫️"},
			Row{"⚫️", "⚪️"},
			Row{"⚫️", "⚪️"},
			Row{"⚫️", "⚪️"},
			Row{"⚪️", "⚫️"},
		},
		")": Letter{
			Row{"⚫️", "⚪️"},
			Row{"⚪️", "⚫️"},
			Row{"⚪️", "⚫️"},
			Row{"⚪️", "⚫️"},
			Row{"⚫️", "⚪️"},
		},
		"+": Letter{
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
		},
		"-": Letter{
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
		},
		"/": Letter{
			Row{"⚪️", "⚪️", "⚫️"},
			Row{"⚪️", "⚪️", "⚫️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
		},
		":": Letter{
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
		},
		";": Letter{
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
		},
		"<": Letter{
			Row{"⚪️", "⚪️", "⚫️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️"},
		},
		">": Letter{
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
		},
		"=": Letter{
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚫️"},
			Row{"⚪️", "⚪️", "⚪️"},
		},
		"[": Letter{
			Row{"⚫️", "⚫️"},
			Row{"⚫️", "⚪️"},
			Row{"⚫️", "⚪️"},
			Row{"⚫️", "⚪️"},
			Row{"⚫️", "⚫️"},
		},
		"]": Letter{
			Row{"⚫️", "⚫️"},
			Row{"⚪️", "⚫️"},
			Row{"⚪️", "⚫️"},
			Row{"⚪️", "⚫️"},
			Row{"⚫️", "⚫️"},
		},
		"\\": Letter{
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️"},
			Row{"⚪️", "⚪️", "⚫️"},
		},
		"^": Letter{
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚪️", "⚫️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
		},
		"`": Letter{
			Row{"⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️"},
		},
		"{": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚫️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
		},
		"}": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚫️"},
			Row{"⚪️", "⚪️", "⚫️", "⚪️"},
			Row{"⚫️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
		},
		"|": Letter{
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚪️"},
		},
		"~": Letter{
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚫️", "⚪️", "⚫️"},
			Row{"⚫️", "⚪️", "⚫️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
			Row{"⚪️", "⚪️", "⚪️", "⚪️"},
		},
	},
}
