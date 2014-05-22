//; func ToByteSlice(s string) (bs []byte)
TEXT ·ToByteSlice(SB),NOSPLIT,$0
	JMP ·toByteSlice(SB)

//; func ToString(bs []byte) (s string)
TEXT ·ToString(SB),NOSPLIT,$0
	JMP ·toString(SB)
