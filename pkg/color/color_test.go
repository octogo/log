package color

import "testing"

func TestNew(t *testing.T) {
	defer t.Log(New(NormalDisplay))
	for a := range Attributes {
		c := New(Attributes[a])
		t.Logf("ANSII attribute %d: %#v", a, c)
		for ci := 30; ci <= len(FGColors)+30; ci++ {
			c = New(Attributes[a], Color(ci))
			t.Logf("ANSII color %d: %#v", ci, c)
		}
		for ci := 40; ci <= len(BGColors)+40; ci++ {
			c = New(Attributes[a], Color(ci))
			t.Logf("ANSII color %d: %#v", ci, c)
		}
	}
}
