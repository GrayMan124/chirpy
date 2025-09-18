package chirp

import "strings"

func replaceWords(input string) string {
	split := strings.Split(input, " ")

	for idx, word := range split {
		if strings.ToLower(word) == "kerfuffle" || strings.ToLower(word) == "sharbert" || strings.ToLower(word) == "fornax" {
			split[idx] = "****"
		}
	}
	return strings.Join(split, " ")
}

func ValidateChirp(input string) (string, bool) {

	if len(input) > 140 {
		return "", false
	}
	chirp := replaceWords(input)
	return chirp, true
}
