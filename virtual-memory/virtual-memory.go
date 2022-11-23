package virtualMemory

import (
	"fmt"
	mainMemory "unisinos/so/tgb/main-memory"
	"unisinos/so/tgb/utils"
)

const (
	invalidIndex = -1
	pageLength   = 8    // 8KB
	memoryLength = 1000 // 1000KB == 1MB
)

// VirtualMemory is a fictitious struct to represent the virtual memory, containing it mapped pages to main memory pages.
type VirtualMemory struct {
	Pages []*VirtualMemoryPage
}

// NewVirtualMemory creates a new VirtualMemory.
func NewVirtualMemory() *VirtualMemory {
	pageQuantity := memoryLength / pageLength
	pages := make([]*VirtualMemoryPage, pageQuantity)
	return &VirtualMemory{
		Pages: pages,
	}
}

// AddPage will add a new page into table page if it has availabe slots.
func (vm *VirtualMemory) AddPage(mainMemoryPage *mainMemory.MainMemoryPage) (index int) {
	i := vm.firstIndexAvailable()
	if i != invalidIndex {
		vm.Pages[i] = NewVirtualMemoryPage(mainMemoryPage, true)
	}
	return i
}

// firstIndexAvailable identify the first index that hasn't value recorded.
func (vm *VirtualMemory) firstIndexAvailable() int {
	for i, page := range vm.Pages {
		if page == nil {
			return i
		}
	}
	return invalidIndex
}

func (vm *VirtualMemory) MarkInvalidPage(invalidPage *mainMemory.MainMemoryPage) {
	for _, page := range vm.Pages {
		if page != nil && page.MainMemoryPage == invalidPage {
			page.ValidBit = false
			return
		}
	}
}

// PrintPages will print the virtual memory pages formatted and shows the id of each page in each slot.
func (vm *VirtualMemory) PrintPages() {
	fmt.Println("Páginas da memória virtual")
	utils.PrintSeparator()
	strIndex := ""
	strPage := ""

	for i, page := range vm.Pages {
		if i%25 == 0 && i > 0 {
			fmt.Println(strIndex)
			fmt.Println(strPage)
			utils.PrintSeparator()
			strIndex = ""
			strPage = ""
		}

		strIndex = utils.FormatOutputMemory(strIndex, i)

		if page != nil {
			strPage = utils.FormatOutputMemory(strPage, page.Id)
		} else {
			strPage += " --- "
		}
	}

	fmt.Println(strIndex)
	fmt.Println(strPage)
}
