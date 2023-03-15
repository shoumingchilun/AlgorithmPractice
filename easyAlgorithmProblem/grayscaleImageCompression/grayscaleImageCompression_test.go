package grayscaleImageCompression

import (
	"fmt"
	"testing"
)

func TestCompression(t *testing.T) {
	i, j := Compression([]int{10, 12, 15, 255, 1, 2})
	fmt.Println(i, j)

}
