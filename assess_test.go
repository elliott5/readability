package read

import (
	"testing"
)

// TestAssess is a working example rather than a test.
func TestAssess(t *testing.T) {
	for tid, text := range []string{
		`The cat sat on the mat.`,
		"The Australian platypus is seemingly a hybrid of a mammal and reptilian creature",
		`Be not afraid of greatness: some are born great, some achieve greatness, and some have greatness thrust upon them.`,
		`The researchers discovered that the fang-blenny, a tiny reef-dwelling fish, has a venom that is laced with pain-killing chemicals.`,
	} {

		t.Log(tid, ")", text)
		res := Assess(text)
		t.Logf("Automated Readability: %0.2f\n", res.AutomatedReadability)
		t.Logf("Coleman-Liau: %0.2f\n", res.ColemanLiau)
		t.Logf("Flesch-Kincaid: %0.2f\n", res.FleschKincaid)
		t.Logf("Gunning fog: %0.2f\n", res.GunningFog)
		t.Logf("SMOG: %0.2f\n", res.Smog)
		t.Logf("Average years-of-education: %0.2f\n", res.AvgYrsOfEd)
		t.Logf("Std Dev of years-of-education: %0.2f\n", res.StdDevYrsOfEd)

		yrsStart := res.AvgYrsOfEd - res.StdDevYrsOfEd
		if yrsStart < 0 {
			yrsStart = 0
		}
		t.Logf("Years-of-education 95%% confidence range: %0.2f -- %0.2f\n", yrsStart, res.AvgYrsOfEd+res.StdDevYrsOfEd)

		t.Logf("Flesch Reading-Ease: %0.2f\n", res.FleschReadingEase)
	}
}
