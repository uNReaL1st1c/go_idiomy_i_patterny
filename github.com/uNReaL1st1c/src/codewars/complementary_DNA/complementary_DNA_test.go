package complementarydna

import "testing"

func TestDNAStrand(t *testing.T) {

	test_1_DNA := "AAAA"
	test_2_DNA := "ATTGC"
	test_3_DNA := "GTAT"

	result_1 := DNAStrand(test_1_DNA)

	if result_1 != "TTTT" {
		t.Errorf("incorrect result: expected %s, got %s", "TTTT", result_1)
	}

	result_2 := DNAStrand(test_2_DNA)

	if result_2 != "TAACG" {
		t.Errorf("incorrect result: expected %s, got %s", "TAACG", result_2)
	}

	result_3 := DNAStrand(test_3_DNA)

	if result_3 != "CATA" {
		t.Errorf("incorrect result: expected %s, got %s", "CATA", result_3)
	}

}

func Test_otherComplementaryDNA(t *testing.T) {

	var (
		A = 'A'
		T = 'T'
		C = 'C'
		G = 'G'
	)

	result_A := otherComplementaryDNA(A)

	if result_A != 'T' {
		t.Errorf("incorrect result: expected %c, got %c", 'T', result_A)
	}

	result_C := otherComplementaryDNA(C)

	if result_C != 'G' {
		t.Errorf("incorrect result: expected %c, got %c", 'G', result_C)
	}

	result_G := otherComplementaryDNA(G)

	if result_G != 'C' {
		t.Errorf("incorrect result: expected %c, got %c", 'C', result_G)
	}

	result_T := otherComplementaryDNA(T)

	if result_T != 'A' {
		t.Errorf("incorrect result: expected %c, got %c", 'A', result_T)
	}

}
