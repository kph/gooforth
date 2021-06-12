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
	task.Push(word)
	task.Pop().Execute(task)
}

func (wt wordTest) Execute(t Task) (n Task) {
	fmt.Printf("Hello world\n")
	return nil
}
