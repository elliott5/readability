package read

import (
	"math"
	"strings"
	"unicode/utf8"
)

type Assessment struct {
	AutomatedReadability, ColemanLiau, FleschKincaid, GunningFog, Smog, AvgYrsOfEd, StdDevYrsOfEd, FleschReadingEase float64
}

// Assess all of various indexes, returning the full picture.
// TODO make faster by not repeating the shared underlying calculations.
func Assess(text string) *Assessment {
	ret := new(Assessment)

	text = strings.TrimSpace(text)
	if text == "" {
		return ret // an empty text
	}

	r, _ := utf8.DecodeLastRuneInString(text)
	switch r {
	case '.', '!', '?': // do nothing
	default:
		text += "." // there must be at least one sentence
	}

	ret.AutomatedReadability = Ari(text)
	ret.ColemanLiau = Cli(text)
	ret.FleschKincaid = Fk(text)
	ret.GunningFog = Gfi(text)
	ret.Smog = Smog(text)

	values := []float64{ret.AutomatedReadability, ret.ColemanLiau, ret.FleschKincaid, ret.GunningFog, ret.Smog}
	ret.AvgYrsOfEd = 0.0
	for _, v := range values {
		ret.AvgYrsOfEd += v
	}
	ret.AvgYrsOfEd /= float64(len(values))
	sum := 0.0
	for _, v := range values {
		dev := v - ret.AvgYrsOfEd
		sum += dev * dev
	}
	ret.StdDevYrsOfEd = math.Sqrt(sum / float64(len(values)-1))

	ret.FleschReadingEase = Fre(text)

	return ret
}
