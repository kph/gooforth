package gooforth

import (
	"fmt"
)

type dotWord struct {
}

func (dw *dotWord) Execute(t Task) (n Word) {
	w := t.Pop()
	fmt.Printf("%v", w)
	return nil
}
