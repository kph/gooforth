package gooforth

type fTask struct {
	stack  []Word
	rstack []Word
}

func NewTask() (t Task) {
	ft := &fTask{}
	ft.stack = make([]Word, 0)
	ft.rstack = make([]Word, 0)
	return ft
}

func (ft *fTask) Execute() {
	for len(ft.rstack) > 0 {
		w := ft.RPop()
		for w != nil {
			w = w.Execute(ft)
		}
	}
}

func (ft *fTask) Push(w Word) {
	ft.stack = append(ft.stack, w)
}

func (ft *fTask) Pop() (w Word) {
	l := len(ft.stack)
	if l > 0 {
		r := ft.stack[l-1]
		ft.stack = ft.stack[:l-1]
		return r
	}
	panic("Stack underflow")
}

func (ft *fTask) RPush(w Word) {
	ft.rstack = append(ft.rstack, w)
}

func (ft *fTask) RPop() (w Word) {
	l := len(ft.rstack)
	if l > 0 {
		r := ft.rstack[l-1]
		ft.rstack = ft.rstack[:l-1]
		return r
	}
	panic("Return stack underflow")
}
