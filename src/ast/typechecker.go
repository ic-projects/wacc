package ast

import (
	"fmt"
	"os"
	"bytes"
)

type Expectance interface {
	seen(*TypeChecker, TypeNode)
}

type SetExpectance struct {
	set map[TypeNode] bool
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

func (exp SetExpectance) seen(check *TypeChecker, t TypeNode) {
	validTypes := exp.set

	_, found := validTypes[t]
	if !found {
		typeErr(t, validTypes)
	}
}

type TwiceSameExpectance struct {
	exp Expectance
}

func NewTwiceSameExpectance(exp Expectance) TwiceSameExpectance {
	return TwiceSameExpectance{
		exp: exp,
	}
}

func (exp TwiceSameExpectance) seen(check *TypeChecker, t TypeNode) {
	exp.exp.seen(check, t)
	check.expect(t)
}

type AnyExpectance struct {}

func NewAnyExpectance() AnyExpectance {
	return AnyExpectance{}
}

func (exp AnyExpectance) seen(check *TypeChecker, t TypeNode) {}

type TypeChecker struct {
	stack []Expectance
}

func NewTypeChecker() *TypeChecker {
	stack := make([]Expectance, 0)
	return &TypeChecker{
		stack: stack,
	}
}


func (check *TypeChecker) seen(t TypeNode) {
	if len(check.stack) < 1 {
		//Oh no
	}

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Found %s\n", StripType(t)))
	fmt.Println(b.String())

	expectance := check.stack[len(check.stack) - 1]
	check.stack = check.stack[:len(check.stack) - 1]



	expectance.seen(check, StripType(t))
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

func typeErr(got TypeNode, validTypes map[TypeNode] bool) {
	var b bytes.Buffer
	b.WriteString("Expected type ")
	for key, _ := range validTypes {
			b.WriteString(fmt.Sprintf("%s ", key))
	}
	b.WriteString(fmt.Sprintf("but got type %s\n", got))
	fmt.Println(b.String())
	os.Exit(200)
}

func (check *TypeChecker) expectAny() {
	check.stack = append(check.stack, NewAnyExpectance())
}

func (check *TypeChecker) expectTwiceSame(ex Expectance) {
	check.stack = append(check.stack, NewTwiceSameExpectance(ex))
}

func (check *TypeChecker) expect(t TypeNode) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Adding %s\n", t))
	fmt.Println(b.String())

	check.expectSet([]TypeNode{t})
}

func (check *TypeChecker) expectSet(ts []TypeNode) {
	check.stack = append(check.stack, NewSetExpectance(ts))
}
