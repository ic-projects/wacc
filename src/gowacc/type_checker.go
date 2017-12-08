package main

import (
	"ast"
	"bytes"
	"fmt"
)

// Expectance is an interface used to store what type is expected.
type Expectance interface {
	seen(*TypeChecker, ast.TypeNode) TypeError
}

// SetExpectance is a struct that stores a set of acceptable types that can
// be seen. It implements Expectance.
type SetExpectance struct {
	set []ast.TypeNode
}

func NewSetExpectance(ts []ast.TypeNode) SetExpectance {
	return SetExpectance{
		set: ts,
	}
}

func contains(check *TypeChecker, arr []ast.TypeNode, t ast.TypeNode) bool {
	for _, a := range arr {
		if checkEquals(check, a, t) {
			return true
		}
	}
	return false
}

func checkEquals(check *TypeChecker, expecting ast.TypeNode, seen ast.TypeNode) bool {
	expectingValue := ast.ToValue(expecting)
	seenValue := ast.ToValue(seen)
	switch seenValue.(type) {
	case ast.ArrayTypeNode:
		if found, ok := expectingValue.(ast.ArrayTypeNode); ok {
			if seenValue == (ast.ArrayTypeNode{}) {
				check.expectRepeatUntilForce(found.T)
				return true
			}
			return expectingValue == (ast.ArrayTypeNode{}) || expectingValue.Equals(seenValue)
		}
	case ast.PairTypeNode:
		if found, ok := expectingValue.(ast.PairTypeNode); ok {
			if seenValue == (ast.PairTypeNode{}) {
				if expectingValue == (ast.PairTypeNode{}) {
					return false
				}
				check.expect(found.T2)
				check.expect(found.T1)
				return true
			}
			return expectingValue == (ast.PairTypeNode{}) || expectingValue.Equals(seenValue)
		} else if _, ok := expectingValue.(ast.BaseTypeNode); ok && expectingValue.Equals(ast.ToValue(ast.NewBaseTypeNode(ast.PAIR))) {
			return true
		}
	case ast.BaseTypeNode:
		if seenValue.Equals(ast.ToValue(ast.NewBaseTypeNode(ast.PAIR))) {
			if _, ok := expectingValue.(ast.PairTypeNode); ok {
				if seenValue == (ast.PairTypeNode{}) {
					if expectingValue == (ast.PairTypeNode{}) {
						return false
					}
					return true
				}
				return true
			}
		}

		return expectingValue.Equals(seenValue)
	case ast.StructTypeNode:
		return expectingValue.Equals(seenValue)
	case ast.PointerTypeNode:
		return expectingValue.Equals(seenValue)
	case ast.NullTypeNode:
		return true
	}
	return false
}

// seen is called when we have seen a SetExpectance.
func (exp SetExpectance) seen(check *TypeChecker, typeNode ast.TypeNode) TypeError {
	validTypes := exp.set
	redoSeen := false
	if dyn, ok := typeNode.(*ast.DynamicTypeNode); ok {
		if _, ok := validTypes[0].(*ast.NullTypeNode); !ok && typeNode != nil {
			if newType, ok := dyn.ReduceSet(validTypes); ok {
				redoSeen = (newType != typeNode)
				typeNode = newType
			} else {
				return NewTypeError(typeNode, validTypes)
			}
		}
	}

	if dyn, ok := validTypes[0].(*ast.DynamicTypeNode); len(validTypes) == 1 && ok {
		if _, ok := typeNode.(*ast.NullTypeNode); !ok && typeNode != nil {
			if newType, ok := dyn.ReduceSet([]ast.TypeNode{typeNode}); ok {
				redoSeen = (newType != validTypes[0])
				validTypes[0] = newType
			} else {
				return NewTypeError(typeNode, validTypes)
			}
		}
	}
	if redoSeen {
		return NewSetExpectance(validTypes).seen(check, typeNode)
	}

	if _, ok := typeNode.(*ast.DynamicTypeNode); ok {
		return TypeError{}
	}
	if contains(check, validTypes, typeNode) {
		return TypeError{}
	}
	return NewTypeError(typeNode, validTypes)
}

// TwiceSameExpectance is a struct for when we want the next two types to be
// the same. This would be used for an assign statement, for example. It
// implements Expectance.
type TwiceSameExpectance struct {
	exp Expectance
}

func NewTwiceSameExpectance(exp Expectance) TwiceSameExpectance {
	return TwiceSameExpectance{
		exp: exp,
	}
}

// seen is called when we have seen a TwiceSameExpectance.
func (exp TwiceSameExpectance) seen(check *TypeChecker, t ast.TypeNode) TypeError {
	typeError := exp.exp.seen(check, t)
	if t == nil { //hmm
		check.expectAny()
	} else {
		check.expect(t)
	}

	return typeError
}

// RepeatExpectance is a struct for an expectance that is used multiple times,
// such as for an ArrayLiteral where all elements should be of a specific type.
// It implements Expectance.
type RepeatExpectance struct {
	exp Expectance
}

func NewRepeatExpectance(exp Expectance) RepeatExpectance {
	return RepeatExpectance{
		exp: exp,
	}
}

// seen is called when we have seen a RepeatExpectance. It will stop it from
// being removed from the stack by adding an extra expectance before seeing
// the expectance.
func (exp RepeatExpectance) seen(check *TypeChecker, t ast.TypeNode) TypeError {
	check.stack = append(check.stack, exp)
	return exp.exp.seen(check, t)
}

// AnyExpectance is an empty struct, allowing for any type.
type AnyExpectance struct{}

func NewAnyExpectance() AnyExpectance {
	return AnyExpectance{}
}

// seen is called when we have seen an AnyExpectance. It allows anything, and
// returns an empty error, i.e. no error.
func (exp AnyExpectance) seen(check *TypeChecker, t ast.TypeNode) TypeError {
	return TypeError{}
}

// TypeChecker stores a stack of expectance. Seeing a type will pop it off from
// the stack, while expecting a type will push the type onto the stack.
// It can be frozen at a ProgramNode to prevent incorrect errors which can
// happen after some errors.
type TypeChecker struct {
	stack      []Expectance
	frozenNode ast.ProgramNode
}

// NewTypeChecker will return an initialised TypeChecker, with an empty stack.
func NewTypeChecker() *TypeChecker {
	stack := make([]Expectance, 0)
	return &TypeChecker{
		stack: stack,
	}
}

// seen will pop the type from the stack, and return a TypeError corresponding
// to the mismatch between the type popped off the stack and the TypeNode given.
func (check *TypeChecker) seen(t ast.TypeNode) TypeError {
	if check.frozen() {
		return TypeError{}
	}
	if len(check.stack) < 1 {
		fmt.Println("Internal type checker error")
		return TypeError{}
	}
	if ast.DebugMode {
		fmt.Printf("Seen %s  -- p %T: &p=%p p=&i=%p \n", t, t, &t, t)
	}

	expectance := check.stack[len(check.stack)-1]
	check.stack = check.stack[:len(check.stack)-1]

	return expectance.seen(check, t)
}

// StripType is used to remove the type of Arrays and Pairs, which is useful
// for comparing types.
func StripType(t ast.TypeNode) ast.TypeNode {
	switch t.(type) {
	case *ast.ArrayTypeNode:
		return ast.ArrayTypeNode{}
	case *ast.PairTypeNode:
		return ast.PairTypeNode{}
	default:
		fmt.Println("Internal type checker error, unknown typenode")
	}
	return t
}

// forcePop will force an expectance off the stack, useful for RepeatExpectance.
// It will only change the stack if it is not frozen.
func (check *TypeChecker) forcePop() {
	if check.frozen() {
		return
	}
	if ast.DebugMode {
		fmt.Println("Force pop")
	}
	if len(check.stack) < 1 {
		fmt.Println("Internal type checker error, stack ran out")
		return
	}
	check.stack = check.stack[:len(check.stack)-1]
}

// expectAny adds a AnyExpectance to the stack, if not frozen.
func (check *TypeChecker) expectAny() {
	if check.frozen() {
		return
	}
	if ast.DebugMode {
		fmt.Println("Expecting any")
	}
	check.stack = append(check.stack, NewAnyExpectance())
}

// expectTwiceSame adds a TwiceSameExpectance to the stack, if not frozen.
func (check *TypeChecker) expectTwiceSame(ex Expectance) {
	if check.frozen() {
		return
	}
	if ast.DebugMode {
		fmt.Println("Expecting twice")
	}
	check.stack = append(check.stack, NewTwiceSameExpectance(ex))
}

// expectRepeatUntilForce adds a RepeatExpectance to the stack, if not
// frozen.
func (check *TypeChecker) expectRepeatUntilForce(t ast.TypeNode) {
	if check.frozen() {
		return
	}
	if ast.DebugMode {
		fmt.Printf("Expecting repeat %s\n", t)
	}
	check.stack = append(check.stack, NewRepeatExpectance(NewSetExpectance([]ast.TypeNode{t})))
}

// expect adds a SetExpectance with the given TypeNode the only element in the set,
// if not frozen.
func (check *TypeChecker) expect(t ast.TypeNode) {
	if check.frozen() {
		return
	}
	check.expectSet([]ast.TypeNode{t})
}

// expectSet adds an SetExpectance to the stack, if not frozen.
func (check *TypeChecker) expectSet(ts []ast.TypeNode) {
	if check.frozen() {
		return
	}
	if ast.DebugMode {
		var b bytes.Buffer
		b.WriteString("Expecting ")
		for i, t := range ts {
			b.WriteString(fmt.Sprintf("[%s  -- p %T: &p=%p p=&i=%p]", t, t, &t, t))
			if i != len(ts)-1 {
				b.WriteString(" or ")
			}
		}
		fmt.Println(b.String())
	}
	check.stack = append(check.stack, NewSetExpectance(ts))
}

// frozen returns if the typechecker is frozen or not.
func (check *TypeChecker) frozen() bool {
	return check.frozenNode != nil
}

// freeze freezes the type checker on a node.
func (check *TypeChecker) freeze(node ast.ProgramNode) {
	if check.frozen() {
		return
	}
	if ast.DebugMode {
		fmt.Printf("Frozen on %s\n", node)
	}
	check.frozenNode = node
}

// unfreeze will unfreeze the checker, if the given node is the same as the frozen node.
func (check *TypeChecker) unfreeze(node ast.ProgramNode) {
	if node == check.frozenNode {
		if ast.DebugMode {
			fmt.Printf("Unfrozen on %s\n", node)
		}
		check.frozenNode = nil
	}
}
