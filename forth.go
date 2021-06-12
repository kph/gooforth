package gooforth

type Task interface {
	Push(Word)
	Pop() Word
	RPush(Word)
	RPop() Word
}

type Word interface {
	Execute(Task) Task
}
