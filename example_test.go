package countmin

import (
	"fmt"
)

func ExampleCountMin_Add() {
	cm := New(10, 100000000)
	for _, i := range []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1} {
		cm.Add([]byte(fmt.Sprintf("%d", i)), int64(i))
	}
	fmt.Printf("Estimate of %d is %d\n", 1, cm.Count([]byte("1")))
	fmt.Printf("Estimate of %d is %d\n", 3, cm.Count([]byte("3")))
	fmt.Printf("Estimate of %d is %d\n", 9, cm.Count([]byte("9")))
	fmt.Println("Size: ", cm.Size())
	fmt.Println("Err: ", cm.RelativeError())
	fmt.Println("Confidence: ", cm.Confidence())
	// Output:
	// Estimate of 1 is 3
	// Estimate of 3 is 6
	// Estimate of 9 is 18
	// Size:  91
	// Err:  2e-08
	// Confidence:  0.9990234375
}
