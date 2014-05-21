Assembly in Go
======

--------------

Why use Assembly in Go? 
=======================

Pros:

* Speed up computationally intensive functions.
* Access CPU specific instructions or capabilities that are not available otherwise.
* Good way to learn and play around with assembly.
* go vet works on assembly files.

Cons:

* Very hard to write bug-free, safe and stable code.
* Binds your code to the CPU architectures you have written assembly for (unless a fallback function written in go is created. More on this later). 
* It may breake in the future if go shanges internaly.
* Currently the functions written in assembly will not be inlined like small go functions will.
* It will not be covered by the go tool cover. 
* It will not work with go fmt.	

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

Before using Assembly in Go 
===========================

Have you tried...

* Identifying what the program is spending most time on by profiling? 

* Algoritmic optimizations if it is applicable?

* Minimise the generation of garbage? (conversions from string to byte slice and vice versa etc.)

* Finding a good library that already solved this problem?

If nothing else works then you can consider rewriting in parts or as a 
whole the computationaly intence function in optimized assembly. 

--------------

Before writing Assembly in Go 
=============================

* Write test cases for every possible variation of input (there are too many pitfalls to do anything else).

Alignment, Page boundaries, Overflows, Special cases  etc.

--------------

Ok lets start with a basic example
==================================

package typehack

// ToByteSlice returns the string s as a []byte. Note that cap in bs is set to the length of s.
// WARNING! bs uses the same memory as s and has therefore broken the 
// immutability of the string s. This can lead to serious bugs, 
// vulnerabilities or other errors if used incorrectly. 
func ToByteSlice(s string) (bs []byte) 

// ToString returns the []byte bs as a string. Note that the cap value from bs is discarded.
// WARNING! s uses the same memory as bs and is therfor not Immutable. 
// This can lead to serious bugs, vulnerabilities or other errors if used incorrectly. 
func ToString(bs []byte) (s string) 

// Fallback for not supported architectures
func toByteSlice(s string) (bs []byte) {
	return []byte(s)
}

// Fallback for not supported architectures
func toString(bs []byte) (s string) {
	return string(bs)
}


--------------

Syntax
==================================




--------------

Thank you!
==========
