package gooforth

import (
	"testing"
)

type wordTest struct {
}

var stringHello = TypeString{V: "Hello world"}
var byteNL = TypeByte{V: '\n'}

var intOne = TypeInt{V: 1}
var intTwo = TypeInt{V: 2}

func TestForth(t *testing.T) {
	task := NewTask()

	task.Push(&byteNL)
	task.Push(&intTwo)
	task.Push(&byteNL)
	task.Push(&stringHello)
	task.Push(&byteNL)
	task.Push(&intOne)

	task.RPush(&dotWord{})
	task.RPush(&dotWord{})
	task.RPush(&dotWord{})
	task.RPush(&dotWord{})
	task.RPush(&dotWord{})
	task.RPush(&dotWord{})

	task.Execute()
}
