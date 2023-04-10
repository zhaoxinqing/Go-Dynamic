package public

import (
	"fmt"
	"testing"
)

func TestDivFloat64(t *testing.T) {

	x := 0.1
	y := 0.2

	fmt.Println(AddFloat64(x, y)) // 输出：0.3
	fmt.Println(SubFloat64(x, y)) // 输出：-0.1
	fmt.Println(MulFloat64(x, y)) // 输出：0.02
	fmt.Println(DivFloat64(x, y)) // 输出：0.5

}
