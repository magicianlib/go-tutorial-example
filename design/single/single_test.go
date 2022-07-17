package single

import "testing"

func TestGetInstance(t *testing.T) {

	i1 := GetInstance()
	i2 := GetInstance()

	if i1 != i2 {
		t.Fatalf("i1(%p) equale i2(%p): %t\n", i1, i2, i1 == i2)
	} else {
		t.Logf("i1(%p) equale i2(%p): %t\n", i1, i2, i1 == i2)
	}
}
