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
