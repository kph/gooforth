package gooforth

import (
	"fmt"
)

type dotWord struct {
}

func (dw *dotWord) Execute(t Task) (n Word) {
	w := t.Pop()
	fmt.Printf("%v\n", w)
	return nil
}
