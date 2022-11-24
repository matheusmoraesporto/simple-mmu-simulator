# Simple MMU simulator
This project is a simples Memory Management Unit simulator, developed to system operationals subject's degree.


### How to run
Into repository root directory, just run the following command:
```
go run .
```

## Introduction
In the main file, there is a global MMU, to represent OS MMU, it will manage the main memory and the virtual memory. It's been using go routines to run parallel call to simulate an access page, the idea is create concurrency to load pages in our memory and the mmu should manage it and load each memory according to page size.
