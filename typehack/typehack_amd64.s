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
MOVQ $0x10, AX
