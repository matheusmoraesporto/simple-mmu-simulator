# Simple MMU simulator
This project is a simples Memory Management Unit simulator, developed to system operationals subject's degree.


### How to run
Into repository root directory, just run the following command:
```
go run .
```

## Introduction
In the main file, there is a global MMU, to represent OS MMU, it will manage the main memory and the virtual memory, mapping pages references. It's been using go routines to run parallel call to simulate an access page, the idea is create concurrency to load pages in our memory and the mmu should manage it and load each memory according to page size.

### MMU
When a page is requested to mmu, first it will check in the virtual memory if the requested page already is loaded in the virtual memory, if true just change the last access to this page (it will be important in the main memory explanation). Case it's not loaded, so the MMU need to first load the page in the main memory, it will create a reference to this memory and it will be mapped to the virtual memory.
Just an important point, if the virtual memory is already full, it will not add anymore pages, just access if the random page selected is loaded in the virtual memory.

