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
}

func (wt wordTest) Execute(t Task) (n Word) {
	fmt.Printf("Hello world\n")
	return nil
}
