// excercism Side excercise - leap
package leap

//  given a year, this functon will return a bool value true if it is Leap-year
//  else this will return false
func IsLeapYear(inputYear int) bool {
	divisibleBy4 := inputYear%4 == 0
	divisibleBy100 := inputYear%100 == 0
	divisibleBy400 := inputYear%400 == 0
	if divisibleBy4{
		if !divisibleBy100 {
			return true
		}
		if (divisibleBy100 && divisibleBy400){
			return true
		}
	}
	return false
}
