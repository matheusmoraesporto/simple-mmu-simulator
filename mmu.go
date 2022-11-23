package main

import (
	"log"
	"sync"
	"time"
	mainMemory "unisinos/so/tgb/main-memory"
	virtualMemory "unisinos/so/tgb/virtual-memory"
)

const (
	invalidIndex = -1
	pagesLen     = 10
)

// MMU is a fictitious struct that represents the real MMU, it has a point to main memory and virtual memory.
type MMU struct {
	*mainMemory.MainMemory
	*virtualMemory.VirtualMemory
	mu *sync.Mutex
}

// NewMMU creates new MMU.
func NewMMU() *MMU {
	return &MMU{
		MainMemory:    mainMemory.NewMainMemory(),
		VirtualMemory: virtualMemory.NewVirtualMemory(),
		mu:            new(sync.Mutex),
	}
}

// AccessPage simulate the behaviour of an access to a page. If the page is loaded in the main memory, it shows.
// If it isn't, add the new page in the main memory and virtual memory.
func (mmu *MMU) AccessPage(idPage int) {
	mmu.mu.Lock()
	defer mmu.mu.Unlock()
	log.Printf("Solicitando acesso à página %d\n", idPage)

	index := mmu.MainMemory.GetPage(idPage) // TODO: Usar o bit de validação da virtual memory
	if index == invalidIndex {
		mmu.AddPage(idPage)
		log.Printf("A página %d não estava na memória e foi adicionada.\n", idPage)
	} else {
		mmu.MainMemory.Pages[index].LastAccess = time.Now()
		log.Printf("A página %d já estava na memória e foi acessada.\n", idPage)
	}

	mmu.VirtualMemory.PrintPages()
	mmu.MainMemory.PrintPages()
}

// AddPage adds the new page into main memory and virtual memory.
func (mmu *MMU) AddPage(idPage int) {
	newPage, replacedPage := mmu.MainMemory.AddPage(idPage)
	mmu.VirtualMemory.MarkInvalidPage(replacedPage)
	mmu.VirtualMemory.AddPage(newPage)
}
