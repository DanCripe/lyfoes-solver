package main

type Color int

const (
	DarkBlue Color = iota
	MediumBlue
	LightBlue
	DarkRed
	MediumRed
	LightRed
	DarkGreen
	MediumGreen
	LightGreen
	White
	Grey
	Orange
	Yellow
	Purple
)

var ColorMap = map[Color]string{
	DarkBlue:    "DB",
	MediumBlue:  "MB",
	LightBlue:   "LB",
	DarkRed:     "DR",
	MediumRed:   "MR",
	LightRed:    "LR",
	DarkGreen:   "DG",
	MediumGreen: "MG",
	LightGreen:  "LG",
	White:       "Wh",
	Grey:        "Gr",
	Orange:      "Or",
	Yellow:      "Ye",
	Purple:      "Pu",
}

var ReverseColorMap = map[string]Color{
	"DB": DarkBlue,
	"MB": MediumBlue,
	"LB": LightBlue,
	"DR": DarkRed,
	"MR": MediumRed,
	"LR": LightRed,
	"DG": DarkGreen,
	"MG": MediumGreen,
	"LG": LightGreen,
	"Wh": White,
	"Gr": Grey,
	"Or": Orange,
	"Ye": Yellow,
	"Pu": Purple,
}
