package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"unisinos/so/tgb/utils"
)

const timeToAcessMemory = 2

var mmu = NewMMU()

func main() {
	showProgramInfo()
	// accessMultipleRandomPages(mmu)
	accessRandomPage(mmu)
}

func accessMultipleRandomPages(mmu *MMU) {
	for {
		go accessMultipleRandomPages(mmu)
	}
}

func accessRandomPage(mmu *MMU) {
	for {
		rand.Seed(time.Now().UnixNano())
		idPage := rand.Intn(125)
		timeout := time.After(time.Second * timeToAcessMemory)
		select {
		case <-timeout:
			log.Printf("Solicitando acesso à página %d\n", idPage)
			mmu.AccessPage(idPage)
		}
	}
}

func showProgramInfo() {
	fmt.Println("Iniciando a simulação de gerenciamento de memória")
	utils.PrintSeparator()
	fmt.Println("MEMÓRIA VIRTUAL:")
	fmt.Println("Tamanho total de 1MB - com 125 slots de páginas de 8k")
	utils.PrintSeparator()
	fmt.Println("MEMÓRIA PRINCIPAL:")
	fmt.Println("Tamanho total de 64KB - com 8 slots de páginas de 8k")
	utils.PrintSeparator()
}
