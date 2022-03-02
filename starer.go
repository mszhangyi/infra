package infra

import (
	"sort"
)

type Starter interface {
	Init()
	Setup()
	Start()
	StartBlocking() bool
	Stop()
	PriorityGroup() PriorityGroup
	Priority() int
}

type starterRegister struct {
	nonBlockingStarters []Starter
	blockingStarters    []Starter
}

func (r *starterRegister) AllStarters() []Starter {
	starters := make([]Starter, 0)
	starters = append(starters, r.nonBlockingStarters...)
	starters = append(starters, r.blockingStarters...)
	return starters

}

func (r *starterRegister) Register(starter Starter) {
	if starter.StartBlocking() {
		r.blockingStarters = append(r.blockingStarters, starter)
	} else {
		r.nonBlockingStarters = append(r.nonBlockingStarters, starter)
	}
}

var StarterRegister *starterRegister = &starterRegister{}

type Starters []Starter

func (s Starters) Len() int      { return len(s) }
func (s Starters) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Starters) Less(i, j int) bool {
	return s[i].PriorityGroup() > s[j].PriorityGroup() && s[i].Priority() > s[j].Priority()
}

func Register(starter Starter) {
	StarterRegister.Register(starter)
}

func SortStarters() {
	sort.Sort(Starters(StarterRegister.AllStarters()))
}

func GetStarters() []Starter {
	return StarterRegister.AllStarters()
}

type PriorityGroup int

const (
	SystemGroup         PriorityGroup = 30
	BasicResourcesGroup PriorityGroup = 20
	AppGroup            PriorityGroup = 10

	INT_MAX          = int(^uint(0) >> 1)
	DEFAULT_PRIORITY = 10000
)

type BaseStarter struct {}

func (s *BaseStarter) Init()      {}
func (s *BaseStarter) Setup()     {}
func (s *BaseStarter) Start()     {}
func (s *BaseStarter) Stop()      {}
func (s *BaseStarter) StartBlocking() bool          { return false }
func (s *BaseStarter) PriorityGroup() PriorityGroup { return BasicResourcesGroup }
func (s *BaseStarter) Priority() int                { return DEFAULT_PRIORITY }
