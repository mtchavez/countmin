package main

import (
	"fmt"

	"github.com/mtchavez/countmin/countmin"
)

func main() {
	cm := countmin.New(10, 100000000)
	for _, i := range []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1} {
		cm.Add([]byte(fmt.Sprintf("%d", i)), int64(i))
	}
	fmt.Printf("Estimate of %d is %d\n", 1, cm.Count([]byte("1")))
	fmt.Printf("Estimate of %d is %d\n", 3, cm.Count([]byte("3")))
	fmt.Printf("Estimate of %d is %d\n", 9, cm.Count([]byte("9")))
	fmt.Println("Size: ", cm.Size())
	fmt.Println("Err: ", cm.RelativeError())
	fmt.Println("Confidence: ", cm.Confidence())
}
