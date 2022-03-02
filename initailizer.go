package infra

type Initializer interface {
	Init()
	//Stop()
}

type InitializeRegister struct {
	Initializers []Initializer
}

func (i *InitializeRegister) Register(ai Initializer) {
	i.Initializers = append(i.Initializers, ai)
}

type WebStarter struct {}
func (s *WebStarter) Init(){}
//func (s *WebStarter) Stop(){}
