package mainMemory

import "time"

// MainMemoryPage is a fictitious structus that represents a page in the main memory.
type MainMemoryPage struct {
	Id         int
	LastAccess time.Time
	Size       int
}

// NewPage creates a new MainMemoryPage.
func NewPage(id, size int) MainMemoryPage {
	return MainMemoryPage{
		Id:         id,
		LastAccess: time.Now(),
		Size:       size,
	}
}
