package icinga

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	s := ServiceStatusOk
	fmt.Printf("%T %v %d\n", s, s, s)
}
