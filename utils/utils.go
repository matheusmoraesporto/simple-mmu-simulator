package utils

import "fmt"

// FormatOutputMemory will format values to show as a table.
func FormatOutputMemory(output string, value int) string {
	if value < 10 {
		output += fmt.Sprintf("  %d  ", value)
	} else if value < 100 {
		output += fmt.Sprintf("  %d ", value)
	} else {
		output += fmt.Sprintf(" %d ", value)
	}

	return output
}

// PrintSeparator will print the default separtor.
func PrintSeparator() {
	fmt.Println("============================================================================================================================")
}
