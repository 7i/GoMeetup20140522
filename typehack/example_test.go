package typehack_test

import "fmt"
import typehack "."

func ExampleToByteSlice() {
	s := "Hello, Gophers!"
	fmt.Println("ToByteSlice:", ToByteSlice(s))
	fmt.Println("Raw s:", s)
}


func ExampleToString() {
	bs := []byte("Hello, Gophers!")
	fmt.Println("ToString:", ToString(bs))
	fmt.Println("Raw bs:", bs)
}
