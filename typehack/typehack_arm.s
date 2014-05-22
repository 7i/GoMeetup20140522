//; func ToByteSlice(s string) (bs []byte)
TEXT ·ToByteSlice(SB),NOSPLIT,$0
	B ·toByteSlice(SB)

//; func ToString(bs []byte) (s string)
TEXT ·ToString(SB),NOSPLIT,$0
	B ·toString(SB)
