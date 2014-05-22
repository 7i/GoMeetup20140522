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
