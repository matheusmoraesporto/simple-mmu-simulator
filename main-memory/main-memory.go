package mainMemory

import (
	"fmt"
	"unisinos/so/tgb/utils"
)

const (
	invalidIndex = -1
	pageLength   = 8  // 8KB
	memoryLength = 64 // 64KB
)

// MainMemory is a fictitious struct to represent the physical memory, containing it pages.
type MainMemory struct {
	Pages []MainMemoryPage
}

// NewMainMemory creates a new MainMemory.
func NewMainMemory() *MainMemory {
	pageQuantity := memoryLength / pageLength
	pages := make([]MainMemoryPage, pageQuantity)
	return &MainMemory{
		Pages: pages,
	}
}

// AddPage will add a new page into table page if it has availabe slots.
// If hasn't, will replace a page.
// Returns the new page added and the page replaced if it was.
func (mm *MainMemory) AddPage(idPage, pageSize int) (newPages []*MainMemoryPage, replacedPages []*MainMemoryPage) {
	for pageSize >= 0 {
		size := 8
		if pageSize < 8 {
			size = pageSize
		}

		if i := mm.firstIndexAvailable(); i != invalidIndex {
			mm.Pages[i] = NewPage(idPage, size)
			newPages = append(newPages, &mm.Pages[i])
		} else {
			new, replaced := mm.replacePage(idPage, size)
			newPages = append(newPages, new)
			replacedPages = append(replacedPages, replaced)
		}

		pageSize = pageSize - pageLength
	}
	return
}

// firstIndexAvailable identify the first index that hasn't value recorded.
func (mm *MainMemory) firstIndexAvailable() int {
	for i, page := range mm.Pages {
		if page.Id == 0 {
			return i
		}
	}
	return invalidIndex
}

// replacePage is a LRU algorithm. It will replace the least recently used page by the new page.
// Returns the new page added and the page replaced if it was.
func (vm *MainMemory) replacePage(idPage, pageSize int) (*MainMemoryPage, *MainMemoryPage) {
	newPage := NewPage(idPage, pageSize)
	indexReplace := 0
	for i, page := range vm.Pages {
		if page.Id != newPage.Id && page.LastAccess.Before(vm.Pages[indexReplace].LastAccess) {
			indexReplace = i
		}
	}

	replacedPage := vm.Pages[indexReplace]
	vm.Pages[indexReplace] = newPage

	return &vm.Pages[indexReplace], &replacedPage
}

// PrintPages will print the main memory pages formatted and shows the id of each page in each slot.
func (mm *MainMemory) PrintPages() {
	utils.PrintSeparator()
	fmt.Println("Páginas da memória principal")
	utils.PrintSeparator()
	strIndex := ""
	strPage := ""

	for i, page := range mm.Pages {
		if i%25 == 0 && i > 0 {
			fmt.Println(strIndex)
			fmt.Println(strPage)
			utils.PrintSeparator()
			strIndex = ""
			strPage = ""
		}

		strIndex = utils.FormatOutputMemory(strIndex, i)

		if page.Id > 0 {
			strPage = utils.FormatOutputMemory(strPage, page.Id)
		} else {
			strPage += " --- "
		}
	}

	fmt.Println(strIndex)
	fmt.Println(strPage)
	utils.PrintSeparator()
}
