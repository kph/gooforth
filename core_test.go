package gooforth

import (
	"fmt"
	"testing"
)

type wordTest struct {
}

func TestForth(t *testing.T) {
	task := NewTask()
	word := wordTest{}
	task.RPush(word)
	task.Execute()

	task.Push(&TypeIntOne)
	dw := &dotWord{}
	dw.Execute(task)
}

func (wt wordTest) Execute(t Task) (n Word) {
	fmt.Printf("Hello world\n")
	return nil
}
