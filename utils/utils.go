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

func ShowProgramInfo() {
	fmt.Println("Iniciando a simulação de gerenciamento de memória")
	PrintSeparator()
	fmt.Println("MEMÓRIA VIRTUAL:")
	fmt.Println("Tamanho total de 1MB - com 125 slots de páginas de 8Kb")
	PrintSeparator()
	fmt.Println("MEMÓRIA PRINCIPAL:")
	fmt.Println("Tamanho total de 64KB - com 8 slots de páginas de 8Kb")
	PrintSeparator()
}
