package pack

import "testing"

func TestCanAddNumber(t *testing.T)  {
	result := Add(1,2)
	if result != 3 {
		t.Log("Failed to add 1 and 2")
		t.Fail()
	}
}