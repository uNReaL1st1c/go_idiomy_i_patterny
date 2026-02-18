package getthemiddlecharacter

func GetMiddle(s string) string {

	runes := []rune(s)
	result := make([]rune, 0, len(runes))

	if len(runes)%2 == 0 {
		result = append(result, runes[(len(runes)/2)-1], runes[(len(runes)/2)])
	} else {
		result = append(result, runes[len(runes)/2])
	}

	return string(result)

}
