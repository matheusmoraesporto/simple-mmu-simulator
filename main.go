package main

import (
	"math/rand"
	"sync"
	"time"
	"unisinos/so/tgb/utils"
)

const (
	maxPageId   = 125
	maxPageSize = 64
)

var mmu = NewMMU()

func main() {
	rand.Seed(time.Now().UnixNano())

	utils.ShowProgramInfo()
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

	idPage := randomGenerator(maxPageId)
	pageSize := randomGenerator(maxPageSize)
	mmu.AccessPage(idPage, pageSize)
}

func randomGenerator(max int) int {
	return rand.Intn(max) + 1
}
