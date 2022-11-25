# Simple MMU simulator
This project is a simples Memory Management Unit simulator, developed to system operationals subject's degree.

### How to run
Into repository root directory, just run the following command:
```
go run .
```

## Introduction
In the main file, there is a global MMU, to represent OS MMU, it will manage the main memory and the virtual memory, mapping pages references. It is been using go routines to run parallel call to simulate an access page, the idea is create concurrency to load pages in our memory and the mmu should manage it and load each memory according to page size.

Memory | Memory size | Page size | Behaviour when page table is full
--- | --- | --- | ---
Main memory | 64Kb | 8Kb | Replace pages using LRU algorithm
Virtual memory | 1Mb | NA | Do not add more pages

### MMU
When a page is requested to mmu, first it will check in the virtual memory if the requested page already is loaded in the virtual memory, if true just change the last access to this page (it will be important in the main memory explanation). Case it is not loaded, so the MMU need to first load the page in the main memory, it will create a reference to this memory and it will be mapped to the virtual memory.
Just an important point, if the virtual memory is already full, it will not add anymore pages, just access if the random page selected is loaded in the virtual memory.

### Virtual memory
It is the first memory layer that will be accessed by the MMU, it has pointers to the page in the main memory, that is identified by the valid bit. When a requested page is not loaded in the main memory, it will communicate to MMU, that will load the page in the main memory and after register into virtual memory.

### Main memory
It is the last memory layer and simulate the physical memory. Here the pages are in fact loaded and could be accessed. When a page is not loaded but all slots are filled, it will use a LRU (Least Recently Used) algorithm, to determine a page to be replaced. It will replace how many pages were necessary to complete the new page regarding it size. To identify the page that will be replaced, the LRU will choose the last recently accessed page to remove it and add the new page. It will be repeated if the page has a size over than 8Kb, that is the page slot size in the main memory, therefore the page is split to be loaded in all slots.
