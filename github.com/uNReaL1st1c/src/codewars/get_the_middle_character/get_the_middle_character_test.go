package getthemiddlecharacter

import "testing"

func TestGetMiddle(t *testing.T) {

	var (
		t1 = "test"
		t2 = "testing"
		t3 = "middle"
		t4 = "A"
	)

	result_1 := GetMiddle(t1)

	if result_1 != "es" {
		t.Errorf("incorrect result: expected %s, got %s", "es", result_1)
	}

	result_2 := GetMiddle(t2)

	if result_2 != "t" {
		t.Errorf("incorrect result: expected %s, got %s", "t", result_2)
	}

	result_3 := GetMiddle(t3)

	if result_3 != "dd" {
		t.Errorf("incorrect result: expected %s, got %s", "dd", result_3)
	}

	result_4 := GetMiddle(t4)

	if result_4 != "A" {
		t.Errorf("incorrect result: expected %s, got %s", "A", result_4)
	}

}
