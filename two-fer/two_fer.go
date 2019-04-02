// Exercims track#2
package twofer

// given a name "X" , this function will return string "One for X, one for me."
// if blank name passed, this function will retunr string "One for you, one for me."
func ShareWith(name string) string {
	if len(name) >0 {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}
 