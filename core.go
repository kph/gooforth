package gooforth

import (
	"fmt"
)

type colonWord struct {
	words []Word
}

// Internal word pushed on return stack to continue with a colon definition
type colonWordNext struct {
	cw    *colonWord // Word we are executing
	index int        // Index into word
}

type dotWord struct {
}

func (cw *colonWord) Execute(t Task) (n Word) {
	if len(cw.words) == 0 {
		return nil // Consier panic - but like this EXIT is optional
	}
	n = cw.words[0]
	cwn := &colonWordNext{cw: cw, index: 1}
	t.RPush(cwn) // Return to process the next entry
	return       // Execute the first word in the list
}

func (cwn *colonWordNext) Execute(t Task) (n Word) {
	if cwn.index >= len(cwn.cw.words) {
		return nil // Consider panic - but like this EXIT is optional
	}
	n = cwn.cw.words[cwn.index]
	cwn.index++
	t.RPush(cwn) // Return to process the next entry
	return       // Execute the first word in the list
}

func (dw *dotWord) Execute(t Task) (n Word) {
	w := t.Pop()
	fmt.Printf("%v", w)
	return nil
}
