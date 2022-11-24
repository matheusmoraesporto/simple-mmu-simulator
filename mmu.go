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
func (mmu *MMU) AccessPage(idPage, pageSize int) {
	mmu.mu.Lock()
	defer mmu.mu.Unlock()
	log.Printf("Solicitando acesso à página %d\n", idPage)

	index := mmu.VirtualMemory.GetValidPage(idPage)
	if index == invalidIndex {
		mmu.AddPage(idPage, pageSize)
	} else {
		// TODO: talvez adicionar alguma informação, por exemplo em quais indícies da mm ela está
		mmu.VirtualMemory.Pages[index].MainMemoryPage.LastAccess = time.Now()
		log.Printf("A página %d(%dKB) já estava na memória e foi acessada.\n", idPage, pageSize)
	}

	mmu.VirtualMemory.PrintPages()
	mmu.MainMemory.PrintPages()
}

// AddPage adds the new page into main memory and virtual memory.
func (mmu *MMU) AddPage(idPage, pageSize int) {
	vmIsFull := mmu.VirtualMemory.FirstIndexAvailable() == invalidIndex

	// Don't add if vm pages are all filled
	if !vmIsFull {
		newPages, replacedPages := mmu.MainMemory.AddPage(idPage, pageSize)

		for _, page := range replacedPages {
			mmu.VirtualMemory.MarkInvalidPage(page)
		}

		for _, page := range newPages {
			mmu.VirtualMemory.AddPage(page)
		}

		log.Printf("A página %d(%dKB) não estava na memória e foi adicionada.\n", idPage, pageSize)
	}
}
