Assembly in Go
======

--------------

Why use Assembly in Go? 
=======================

Pros:

* Speed up computationally intensive functions.
* Access CPU specific instructions or capabilities that are not available otherwise.
* Good way to learn and play around with assembly.

Cons:

* Very hard to write bug-free, safe and stable code.
* Binds your code to the CPU architectures you have written assembly for (unless a fallback function written in go is created. More on this later). 
* It may breake in the future if go shanges internaly.
* Currently the functions written in assembly will not be inlined like small go functions will.
* It will not be covered by the go tool cover. ###test this###
* It will not work with go fmt.	###test this###

--------------

How it works
============

There is no inlining of assembly code in a go file.

Special names for the assembly files (_386.s, _amd64.s, _arm.s).

Declare the assembly function in the go code.

--------------

The assembly function decleration
=================================

Special character for function declerations in the .s files

Middle dot: · (Unicode code point U+00B7)

Example:

TEXT ·Test(SB), 7, $0


--------------

Thank you!
==========
