// Package dog provides a single function for working out dog years
// in human years
package dog

// Years converts human years into dog years.
// humanYears parameter is x7 and the result is returned as an int.
func Years(humanYears int) int {
	dogYears := humanYears * 7
	return dogYears
}
