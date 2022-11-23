package virtualMemory

import mainMemory "unisinos/so/tgb/main-memory"

// VirtualMemoryPage is a fictitious struct that represent the mapped memory in the main memory.
type VirtualMemoryPage struct {
	*mainMemory.MainMemoryPage
	// When is true, it means the page is loaded into main memory.
	ValidBit bool
}

// NewVirtualMemoryPage create a new VirtualMemoryPage.
func NewVirtualMemoryPage(mainMemoryPage *mainMemory.MainMemoryPage, validBit bool) *VirtualMemoryPage {
	return &VirtualMemoryPage{
		MainMemoryPage: mainMemoryPage,
		ValidBit:       validBit,
	}
}
