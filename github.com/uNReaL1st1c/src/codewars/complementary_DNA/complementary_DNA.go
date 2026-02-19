package complementarydna

func DNAStrand(dna string) string {

	runes := []rune(dna)

	result := make([]rune, 0, cap(runes))

	for i := 0; i < len(runes); i++ {
		result = append(result, otherComplementaryDNA(runes[i]))
	}

	return string(result)
}

func otherComplementaryDNA(r rune) rune {
	var result rune

	switch r {
	case 'A':
		result = 'T'
	case 'T':
		result = 'A'
	case 'G':
		result = 'C'
	case 'C':
		result = 'G'
	}

	return result
}
