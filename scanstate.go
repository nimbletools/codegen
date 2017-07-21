package main

type ScanScope struct {
	Name string
	Depth int
}

type ScanState struct {
	Scopes *Stack
	ScopeDepth int
}

func (self *ScanState) Init() {
	self.Scopes = &Stack{}
	self.ScopeDepth = 0
}

func (self *ScanState) PushScope(name string) {
	self.Scopes.Push(ScanScope{ name, self.ScopeDepth + 1 })
}

func (self *ScanState) PopScope() {
	scanState.Scopes.Pop()
}
