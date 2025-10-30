package url_format

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"
)

func FormatNameToUrl(name string) string {
	var res string

	formatName := strings.ToLower(strings.TrimSpace(name))
	for _, char := range formatName {
		if unicode.IsSpace(char) {
			res += "_"
		} else {
			res += string(char)
		}
	}

	randomNumber := rand.Intn(1000) + 1
	res += fmt.Sprintf("_%v", randomNumber)

	return res
}
