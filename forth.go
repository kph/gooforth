package gooforth

type Task interface {
	Execute()   // Execute the task
	Push(Word)  // Push a word on the task stack
	Pop() Word  // Pop a word from the task stack
	RPush(Word) // Push a word on the task return stack
	RPop() Word // Pop a word from the task return stack
}

type Word interface {
	Execute(Task) Word
}

type DictionaryEntry interface {
	Name() string
}
