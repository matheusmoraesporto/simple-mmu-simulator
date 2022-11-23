package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"unisinos/so/tgb/utils"
)

const timeToAcessMemory = 2

var mmu = NewMMU()

func main() {
	showProgramInfo()
	accessMultipleRandomPages(mmu)
}

func accessMultipleRandomPages(mmu *MMU) {
	wg := new(sync.WaitGroup)
	for {
		go accessRandomPage(mmu, wg)
		wg.Add(1)
	}
}

func accessRandomPage(mmu *MMU, wg *sync.WaitGroup) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	idPage := rand.Intn(125)
	timeout := time.After(time.Second * timeToAcessMemory)

	select {
	case <-timeout:
		mmu.AccessPage(idPage)
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
