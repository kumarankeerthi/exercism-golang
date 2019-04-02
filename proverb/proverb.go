// excercism Proverb
package proverb

import (
	"strings"
)

// given a list of string, this function will return proverbs
func Proverb(rhyme []string) []string {
	line := "For want of a X the Y was lost."
	endLine := "And all for the want of a X."

	if len(rhyme) == 0 {
		return []string{}
	}
	proverbs := make([]string, len(rhyme))
	for i := 0; i < len(rhyme)-1; i++ {
		temp := strings.Replace(line, "X", rhyme[i], 1)
		proverbs[i] = strings.Replace(temp, "Y", rhyme[i+1], 1)
	}
	proverbs[len(rhyme)-1] = strings.Replace(endLine, "X", rhyme[0], 1)
	return proverbs
}
