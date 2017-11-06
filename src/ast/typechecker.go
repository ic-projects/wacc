package ast

import (
	"fmt"
	"os"
	"bytes"
)

var DEBUG_MODE = false

type Expectance interface {
	seen(*TypeChecker, TypeNode) TypeError
}

type SetExpectance struct {
	set map[TypeNode]bool
}

func NewSetExpectance(ts []TypeNode) SetExpectance {
	set := make(map[TypeNode]bool)
	for _, t := range ts {
		set[t] = true
	}
	return SetExpectance{
		set: set,
	}
}

func arrayCase(check *TypeChecker, validTypes map[TypeNode]bool, t ArrayTypeNode) bool {
	_, match := validTypes[t]
	nilArray := ArrayTypeNode{}
	setHasNil := false
	findNil := t == nilArray
	var found ArrayTypeNode
	for key, _ := range validTypes {
		if StripType(key) == nilArray {
			found = key.(ArrayTypeNode)
			setHasNil = true
			break
		}
	}

	if findNil {
		if found.dim == 1 {
			check.expectRepeatUntilForce(found.t)
		} else {
			check.expectRepeatUntilForce(NewArrayTypeNode(found.t, found.dim-1))
		}
		return true
	}
	return setHasNil || match
}

func pairCase(check *TypeChecker, validTypes map[TypeNode]bool, basePairMatch bool, t PairTypeNode) bool {
	_, match := validTypes[t]
	_, matchBase := validTypes[NewBaseTypeNode(PAIR)]
	nilPair := PairTypeNode{}
	findNil := t == nilPair
	var nilMatch PairTypeNode
	hasNilMatch := false
	for key, _ := range validTypes {
		if StripType(key) == nilPair {
			nilMatch = key.(PairTypeNode)
			hasNilMatch = true
			break
		}
	}

	if findNil {
		if !hasNilMatch {
			return false
		}
		if !basePairMatch {
			check.expect(nilMatch.t2)
			check.expect(nilMatch.t1)
		}
		return true
	}

	return match || matchBase || hasNilMatch
}

func (exp SetExpectance) seen(check *TypeChecker, typeNode TypeNode) TypeError {
	validTypes := exp.set

	switch t := typeNode.(type) {
	case ArrayTypeNode:
		found := arrayCase(check, validTypes, t)
		if !found {
			return NewTypeError(t, validTypes)
		}
	case PairTypeNode:
		found := pairCase(check, validTypes, false, t)
		if !found {
			return NewTypeError(t, validTypes)
		}
	case BaseTypeNode:
		if t.t == PAIR {
			_, found := validTypes[t]
			if !found {
				found = pairCase(check, validTypes, true, PairTypeNode{})
			}
			if !found {
				return NewTypeError(t, validTypes)
			}
		} else {
			_, found := validTypes[t]
			if !found {
				return NewTypeError(t, validTypes)
			}
		}
	default:
		_, found := validTypes[t]
		if !found {
			return NewTypeError(t, validTypes)
		}
	}

	return TypeError{}
}

type TwiceSameExpectance struct {
	exp Expectance
}

func NewTwiceSameExpectance(exp Expectance) TwiceSameExpectance {
	return TwiceSameExpectance{
		exp: exp,
	}
}

func (exp TwiceSameExpectance) seen(check *TypeChecker, t TypeNode) TypeError {
	typeError := exp.exp.seen(check, t)
	check.expect(t)
	return typeError
}

type RepeatExpectance struct {
	exp Expectance
}

func NewRepeatExpectance(exp Expectance) RepeatExpectance {
	return RepeatExpectance{
		exp: exp,
	}
}

func (exp RepeatExpectance) seen(check *TypeChecker, t TypeNode) TypeError {
	check.stack = append(check.stack, exp)
	return exp.exp.seen(check, t) // ERROR: Probably this
}

type AnyExpectance struct{}

func NewAnyExpectance() AnyExpectance {
	return AnyExpectance{}
}

func (exp AnyExpectance) seen(check *TypeChecker, t TypeNode) TypeError {
	return TypeError{}
}

type TypeChecker struct {
	stack []Expectance
}

func NewTypeChecker() *TypeChecker {
	stack := make([]Expectance, 0)
	return &TypeChecker{
		stack: stack,
	}
}

func (check *TypeChecker) seen(t TypeNode) TypeError {
	if len(check.stack) < 1 {
		fmt.Println("Oh no! Seen type when nun expected.")
		return TypeError{}
	}

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Seen %s\n", t))
	if DEBUG_MODE {
		fmt.Println(b.String())
	}

	expectance := check.stack[len(check.stack)-1]
	check.stack = check.stack[:len(check.stack)-1]

	return expectance.seen(check, t)
}

func StripType(t TypeNode) TypeNode {
	switch t.(type) {
	case ArrayTypeNode:
		return ArrayTypeNode{}
	case PairTypeNode:
		return PairTypeNode{}
	}
	return t
}

func (check *TypeChecker) forcePop() {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Force pop\n"))
	if DEBUG_MODE {
		fmt.Println(b.String())
	}
	check.stack = check.stack[:len(check.stack)-1]
}

func (check *TypeChecker) expectAny() {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Expecting any\n"))
	if DEBUG_MODE {
		fmt.Println(b.String())
	}
	check.stack = append(check.stack, NewAnyExpectance())
}

func (check *TypeChecker) expectTwiceSame(ex Expectance) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Expecting twice\n"))
	if DEBUG_MODE {
		fmt.Println(b.String())
	}
	check.stack = append(check.stack, NewTwiceSameExpectance(ex))
}

func (check *TypeChecker) expectRepeatUntilForce(t TypeNode) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Expecting repeat %s\n", t))
	if DEBUG_MODE {
		fmt.Println(b.String())
	}
	check.stack = append(check.stack, NewRepeatExpectance(NewSetExpectance([]TypeNode{t})))
}

func (check *TypeChecker) expect(t TypeNode) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Expecting %s\n", t))
	if DEBUG_MODE {
		fmt.Println(b.String())
	}

	check.expectSet([]TypeNode{t})
}

func (check *TypeChecker) expectSet(ts []TypeNode) {
	check.stack = append(check.stack, NewSetExpectance(ts))
}
