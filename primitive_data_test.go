/* Examples with primitive data types */

package nanoexamples

import (
	"fmt"
	"math"
)

func ExampleIntDefaultType() {
	i := 1000
	fmt.Printf("%T: %#v\n", i, i)
	// Output:
	// int: 1000
}

func ExampleUintOverflow() {
	maxUint := ^uint(0)
	log2 := math.Log2(float64(maxUint))
	fmt.Printf("   maximum %T: %d (2↑%.0f)\n", maxUint, maxUint, log2)
	maxUint++
	fmt.Printf("maximum %T ++: %d\n", maxUint, maxUint)
	// Output:
	//    maximum uint: 18446744073709551615 (2↑64)
	// maximum uint ++: 0
}

func ExampleIntOverflow() {
	maxInt := int(^uint(0) >> 1)
	log2 := math.Log2(float64(maxInt))
	fmt.Printf("   maximum %T: %d (2↑%.0f)\n", maxInt, maxInt, log2)
	maxInt++
	fmt.Printf("maximum %T ++: %d\n", maxInt, maxInt)
	// Output:
	//    maximum int: 9223372036854775807 (2↑63)
	// maximum int ++: -9223372036854775808
}

func ExampleRuneConstants() {
	char := 'A'
	fmt.Printf("%T: %#v %+q %0x\n", char, char, char, char)
	char = 'Ω'
	fmt.Printf("%T: %#v %+q %0x\n", char, char, char, char)
	// Output:
	// int32: 65 'A' 41
	// int32: 937 '\u03a9' 3a9
}

func ExampleStringBytes() {
	word := "café"
	fmt.Printf("%T: %#v %+q % x\n", word, word, word, word)
	fmt.Printf("len(word) %T: %#v\n", len(word), len(word))
	fmt.Printf("word[0] %T: %#v // %+q %q\n", word[0], word[0], word[0], word[0])
	fmt.Printf("word[3] %T: %#v // %+q %q\n", word[3], word[3], word[3], word[3])
	fmt.Printf("word[4] %T: %#v // %+q %q\n", word[4], word[4], word[4], word[4])
	// Output:
	// string: "café" "caf\u00e9" 63 61 66 c3 a9
	// len(word) int: 5
	// word[0] uint8: 0x63 // 'c' 'c'
	// word[3] uint8: 0xc3 // '\u00c3' 'Ã'
	// word[4] uint8: 0xa9 // '\u00a9' '©'
}

func ExampleStringSlices() {
	word := "café"
	fmt.Printf("%T: %#v %+q % x\n", word, word, word, word)
	fmt.Printf("len(word) %T: %#v\n", len(word), len(word))
	fmt.Printf("word[:3] %T: %#v %+q\n", word[:3], word[:3], word[:3])
	fmt.Printf("word[:4] %T: %#v %+q\n", word[:4], word[:4], word[:4])
	fmt.Printf("word[2:] %T: %#v %+q\n", word[2:], word[2:], word[2:])
	// Output:
	// string: "café" "caf\u00e9" 63 61 66 c3 a9
	// len(word) int: 5
	// word[:3] string: "caf" "caf"
	// word[:4] string: "caf\xc3" "caf\xc3"
	// word[2:] string: "fé" "f\u00e9"
}

func ExampleStringRange() {
	word := "café"
	fmt.Printf("range %q:\n", word)
	for index, runeVal := range word {
		fmt.Printf("    index %T: [%#v] → %T: %3d // %c\n",
			index, index, runeVal, runeVal, runeVal)
	}
	word = "ação"
	fmt.Printf("range %q:\n", word)
	for index, runeVal := range word {
		fmt.Printf("    index %T: [%#v] → %T: %3d // %c\n",
			index, index, runeVal, runeVal, runeVal)
	}
	// Output:
	// range "café":
	//     index int: [0] → int32:  99 // c
	//     index int: [1] → int32:  97 // a
	//     index int: [2] → int32: 102 // f
	//     index int: [3] → int32: 233 // é
	// range "ação":
	//     index int: [0] → int32:  97 // a
	//     index int: [1] → int32: 231 // ç
	//     index int: [3] → int32: 227 // ã
	//     index int: [5] → int32: 111 // o

}
