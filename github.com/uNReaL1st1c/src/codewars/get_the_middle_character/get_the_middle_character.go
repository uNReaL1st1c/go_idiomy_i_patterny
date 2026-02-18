package getthemiddlecharacter

func GetMiddle(s string) string {

	var result []rune

	runes := []rune(s)

	if len(runes)%2 == 0 {
		result = append(result, runes[(len(runes)/2)-1], runes[(len(runes)/2)])
	} else {
		result = append(result, runes[len(runes)/2])
	}

	return string(result)

}
