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
		showAccessedPage(index, idPage)
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

// showAccessedPage will show some details of the page and will set the last access.
func showAccessedPage(index, idPage int) {
	lastAccess := mmu.VirtualMemory.Pages[index].MainMemoryPage.LastAccess
	mmu.VirtualMemory.Pages[index].MainMemoryPage.LastAccess = time.Now()
	log.Printf("A página %d(%dKB) já estava na memória e foi acessada.\n", idPage, mmu.VirtualMemory.Pages[index].MainMemoryPage.Size)
	log.Printf("A página %d tinha sido acessada pela última vez em %v, agora o última acesso é %v.\n", idPage, lastAccess, mmu.VirtualMemory.Pages[index].MainMemoryPage.LastAccess)
}
