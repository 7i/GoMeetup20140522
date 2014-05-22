Assembly in Go
==============

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
* Currently the functions written in assembly will not be inlined like small go functions.
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

.. sourcecode:: python

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

Syntax
======

[Instruction] [src], [dst]

Immediate values:

* MOVQ $1234, AX //; moves 1234 in to AX

* MOVQ $0x12EF, AX //; moves the hex value 0x12EF in to AX

Size of the operands:

* MOVQ is a MOV of size Quad word (8 byte)

* MOVL is a MOV of size Double word or Long word (4 byte)

* MOVW is a MOV of size Word (2 byte)

* MOVB is a MOV of size Byte (1 byte)

Registers:

* SP, AX, BX, CX, DX, BP, DI, SI, R8-R15 

* X0-X15 (XMM registers - 16 byte)

* M0-M7 (MMX registers - 8 byte)

--------------

More Syntax
===========

Addressing:

* AX, (AX), (AX)(BX*4), 10(AX), and 10(AX)(BX*4)

The offsets from AX can be replaced by offsets from FP or SB to access names:

* extern+5(SB)(AX*2)

Examples:

MOVQ AX, BX //; Move the value of AX in to BX

MOVQ (AX), BX //; Move 8 bytes from where AX is pointing in to BX

MOVQ (AX)(BX*4), CX //; Move 8 bytes from where AX + 4*BX is pointing

MOVQ 10(AX)(BX*4), CX //; Move 8 bytes from where AX + 4*BX + 10 is pointing

MOVQ asdf+0(FP), AX //; Move the first 8 bytes located at the Frame pointer in to AX.

//; This is usualy the first argument, or part of the first argument to a function.

//; asdf in this case is just a lable for conviniance.

--------------

More Documentation and examples
===============================

Docs:

http://plan9.bell-labs.com/sys/doc/asm.html

http://golang.org/doc/asm

http://www.doxsey.net/blog/go-and-assembly/index.htm

Examples (all .s files for the architecture you are interested in):

http://golang.org/src/pkg/math/

http://golang.org/src/pkg/runtime/

http://golang.org/src/pkg/crypto/aes/

http://golang.org/src/pkg/crypto/md5/

http://golang.org/src/pkg/crypto/sha1/

http://golang.org/src/pkg/crypto/rc4/

--------------

Starting Point 
==============

Usually when optimising a function you already have the function written in go, 
if this is the case you can easely get a good starting point for your assembly 
by outputing the compiler generated assembly for the function. 

go tool 6g -S gofile.go

--------------

Ok lets start with a basic example
==================================

--------------

typehack.go
===========

.. sourcecode:: go

 package typehack
 
 // ToByteSlice returns the string s as a []byte. Note that cap in bs is 
 // set to the length of s. WARNING! bs uses the same memory as s and has 
 // therefore broken the immutability of the string s. This can lead to 
 // serious bugs, vulnerabilities or other errors if used incorrectly. 
 func ToByteSlice(s string) (bs []byte) 
 
 // ToString returns the []byte bs as a string. Note that the cap value 
 // from bs is discarded. WARNING! s uses the same memory as bs and is 
 // therfor not Immutable. This can lead to serious bugs, vulnerabilities 
 // or other errors if used incorrectly. 
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

example_test.go - ToByteSlice 
=============================

.. sourcecode:: go

 //; func ToByteSlice(s string) (bs []byte)
 TEXT ·ToByteSlice(SB),7,$0
 //       s+0(FP) contains pointer to the s data
 //       s+8(FP) contains length of s
 //       bs+16(FP) contains pointer to bs data
 //       bs+24(FP) contains length of bs
 //       bs+32(FP) contains cap of bs
 
 //; Pointer to data in s is moved to bs pointer so that bs now refers to 
 //; the same data.  
 	MOVQ s+0(FP), AX 
 	MOVQ s+8(FP), BX
 	MOVQ AX, bs+16(FP)  
 	MOVQ BX, bs+24(FP)
 	MOVQ BX, bs+32(FP)
 	RET

--------------

example_test.go - ToString 
==========================

.. sourcecode:: go

 //; func ToString(bs []byte) (s string)
 TEXT ·ToString(SB),7,$0
 //       bs+0(FP) contains pointer to the bs data
 //       bs+8(FP) contains length of bs
 //       bs+16(FP) contains cap of bs
 //       s+24(FP) contains pointer to s data
 //       s+32(FP) contains length of s
 
 //; Pointer to data in bs is moved to s pointer so that bs now refers to 
 //; the same data.  
 	MOVQ bs+0(FP), AX
 	MOVQ bs+8(FP), BX 
 	MOVQ AX, s+24(FP)  
 	MOVQ BX, s+32(FP)
 	RET
--------------

Side Note
=========

To get more information on how the internal structure of []byte and strings please visit:

http://blog.golang.org/go-slices-usage-and-internals

--------------

Fallback for 386 and ARM
========================

typehack_arm.s:

.. sourcecode:: go

 //; func ToByteSlice(s string) (bs []byte)
 TEXT ·ToByteSlice(SB),NOSPLIT,$0
 	B ·toByteSlice(SB)
 
 //; func ToString(bs []byte) (s string)
 TEXT ·ToString(SB),NOSPLIT,$0
 	B ·toString(SB)

typehack_386.s:

.. sourcecode:: go

 //; func ToByteSlice(s string) (bs []byte)
 TEXT ·ToByteSlice(SB),NOSPLIT,$0
 	JMP ·toByteSlice(SB)
 
 //; func ToString(bs []byte) (s string)
 TEXT ·ToString(SB),NOSPLIT,$0
 	JMP ·toString(SB)

--------------

example_test.go 
===============

.. sourcecode:: go

 package typehack_test
 
 import "fmt"
 import typehack "."
 
 func ExampleToByteSlice() {
 	s := "Hello, Gophers!"
 	fmt.Println("ToByteSlice:", typehack.ToByteSlice(s))
 	fmt.Println("Raw s:", s)
 // Output:
 // ToByteSlice: [72 101 108 108 111 44 32 71 111 112 104 101 114 115 33]
 // Raw s: Hello, Gophers!
 }
 
 func ExampleToString() {
 	bs := []byte("Hello, Gophers!")
 	fmt.Println("ToString:", typehack.ToString(bs))
 	fmt.Println("Raw bs:", bs)
 // Output:
 // ToString: Hello, Gophers!
 // Raw bs: [72 101 108 108 111 44 32 71 111 112 104 101 114 115 33]
 }

--------------

Other examples
==============


--------------

Thank you!
==========
