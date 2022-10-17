package schema

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/exp/constraints"
)

type Parameters []Parameter

func (p Parameters) String() string {
	return strings.Join(slicesRemap(p, func(p Parameter) string { return p.String() }), " ")
}

func (p Parameters) Comments() []string {
	paramMaxLen := slicesMaxFunc(p, func(p Parameter) int { return len([]rune(p.GetName())) })
	return slicesRemap(p, func(p Parameter) string {
		return fmt.Sprintf("// @param %-*v %v", paramMaxLen, p.GetName(), p.GetComment())
	})
}

type Parameter interface {
	_Parameter()

	GetName() string
	GetComment() string
	fmt.Stringer
}

var (
	_ Parameter = RequiredParameter{}
	_ Parameter = OptionalParameter{}
	_ Parameter = TriggerParameter{}
)

type RequiredParameter struct {
	Comment string
	Name    string
	Type    Type
}

func (_ RequiredParameter) _Parameter()        {}
func (t RequiredParameter) GetName() string    { return t.Name }
func (t RequiredParameter) GetComment() string { return t.Comment }
func (t RequiredParameter) String() string     { return t.Name + ":" + t.Type.String() }

type OptionalParameter struct {
	Comment     string
	Name        string
	Type        Type
	FlagTrigger string
	BitTrigger  int
}

func (_ OptionalParameter) _Parameter()        {}
func (t OptionalParameter) GetName() string    { return t.Name }
func (t OptionalParameter) GetComment() string { return t.Comment }
func (t OptionalParameter) String() string {
	return fmt.Sprintf("%v:%v.%v?%v", t.Name, t.FlagTrigger, t.BitTrigger, t.Type.String())
}

type TriggerParameter struct {
	Comment     string
	Name        string
	FlagTrigger string
	BitTrigger  int
}

func (_ TriggerParameter) _Parameter()        {}
func (t TriggerParameter) GetName() string    { return t.Name }
func (t TriggerParameter) GetComment() string { return t.Comment }
func (t TriggerParameter) String() string {
	return fmt.Sprintf("%v:%v.%v?true", t.Name, t.FlagTrigger, t.BitTrigger)
}

func slicesRemap[S ~[]I, I, K any](s S, f func(I) K) []K {
	k := make([]K, len(s))
	for i, item := range s {
		k[i] = f(item)
	}

	return k
}

func slicesMaxFunc[S ~[]T, T any](s S, f func(T) int) int {
	if len(s) == 0 {
		return 0
	}

	max := math.MinInt
	for _, item := range s {
		if m := f(item); m > max {
			max = m
		}
	}

	return max
}

func slicesMinFunc[S ~[]T, T any](s S, f func(T) int) int {
	if len(s) == 0 {
		return 0
	}

	min := math.MaxInt
	for _, item := range s {
		if m := f(item); m < min {
			min = m
		}
	}

	return min
}

func slicesSumFunc[S ~[]T, T any, C constraints.Ordered](s S, f func(T) C) (sum C) {
	if len(s) == 0 {
		return sum
	}

	for _, item := range s {
		sum += f(item)
	}

	return sum
}