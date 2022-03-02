package infra

type BootApplication struct {
	IsTest     bool
}

func New() *BootApplication {
	e := &BootApplication{}
	return e
}

func (b *BootApplication) Start() {
	b.init()
	b.setup()
	b.start()
}

func (b *BootApplication) init() {
	for _, v := range GetStarters() {
		v.Init()
	}
}

func (b *BootApplication) setup() {
	for _, v := range GetStarters() {
		v.Setup()
	}
}

func (b *BootApplication) start() {
	for i, v := range GetStarters() {
		if v.StartBlocking() {
			if i+1 == len(GetStarters()) {
				v.Start()
			} else {
				go v.Start()
			}
		} else {
			v.Start()
		}
	}
}

func (b *BootApplication) Stop() {
	for _, v := range GetStarters() {
		v.Stop()
	}
}
