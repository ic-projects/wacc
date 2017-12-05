package main

import (
	"fmt"
)

// Expectance is an interface used to store what type is expected.
type Expectance interface {
	seen(*TypeChecker, TypeNode) TypeError
}

// SetExpectance is a struct that stores a set of acceptable types that can
// be seen. It implements Expectance.
type SetExpectance struct {
	set []TypeNode
}

func NewSetExpectance(ts []TypeNode) SetExpectance {
	return SetExpectance{
		set: ts,
	}
}

func contains(check *TypeChecker, arr []TypeNode, t TypeNode) bool {
	for _, a := range arr {
		if checkEquals(check, a, t) {
			return true
		}
	}
	return false
}

func checkEquals(check *TypeChecker, expecting TypeNode, seen TypeNode) bool {
	expectingValue := toValue(expecting)
	seenValue := toValue(seen)
	//fmt.Println(fmt.Sprintf("checkEquals expect %s and seen %s", expectingValue, seenValue))
	switch seenValue.(type) {
	case ArrayTypeNode:
		if found, ok := expectingValue.(ArrayTypeNode); ok {
			if seenValue == (ArrayTypeNode{}) {
				if found.Dim == 1 {
					check.expectRepeatUntilForce(found.T)
				} else {
					check.expectRepeatUntilForce(NewArrayTypeNode(found.T, found.Dim-1))
				}
				return true
			}
			return expectingValue == (ArrayTypeNode{}) || expectingValue.equals(seenValue)
		}
	case PairTypeNode:
		if found, ok := expectingValue.(PairTypeNode); ok {
			if seenValue == (PairTypeNode{}) {
				if expectingValue == (PairTypeNode{}) {
					return false
				}
				check.expect(found.T2)
				check.expect(found.T1)
				return true
			}
			return expectingValue == (PairTypeNode{}) || expectingValue.equals(seenValue)
		} else if _, ok := expectingValue.(BaseTypeNode); ok && expectingValue.equals(toValue(NewBaseTypeNode(PAIR))) {
			return true
		}
	case BaseTypeNode:
		if seenValue.equals(toValue(NewBaseTypeNode(PAIR))) {
			if _, ok := expectingValue.(PairTypeNode); ok {
				if seenValue == (PairTypeNode{}) {
					if expectingValue == (PairTypeNode{}) {
						return false
					}
					return true
				}
				return true
			}
		}

		return expectingValue.equals(seenValue)
	case StructTypeNode:
		return expectingValue.equals(seenValue)
	case PointerTypeNode:
		return expectingValue.equals(seenValue)
	case NullTypeNode:
		return true
	default:
		// fmt.Println(reflect.TypeOf(seen))
		// fmt.Println(reflect.TypeOf(seenValue))
		// fmt.Println(reflect.TypeOf(expecting))
		// fmt.Println(reflect.TypeOf(expectingValue))
		// fmt.Println("Unknown type for checkEquals")
	}
	return false
}

// seen is called when we have seen a SetExpectance.
func (exp SetExpectance) seen(check *TypeChecker, typeNode TypeNode) TypeError {
	validTypes := exp.set
	redoSeen := false
	//fmt.Println(fmt.Sprintf("Seen is now %s", typeNode))
	if dyn, ok := typeNode.(*DynamicTypeNode); ok {
		if _, ok := validTypes[0].(*NullTypeNode); !ok && typeNode != nil {
			if newType, ok := dyn.reduceSet(validTypes); ok {
				//fmt.Println(fmt.Sprintf("Reduced seen to %s", newType))
				redoSeen = true
				typeNode = newType
			} else {
				return NewTypeError(typeNode, validTypes)
			}
		}
	}
	//fmt.Println(fmt.Sprintf("Seen is now %s", typeNode))
	if dyn, ok := validTypes[0].(*DynamicTypeNode); len(validTypes) == 1 && ok {
		if _, ok := typeNode.(*NullTypeNode); !ok && typeNode != nil {
			//fmt.Println(fmt.Sprintf("Seen is now %s", typeNode))
			if newType, ok := dyn.reduceSet([]TypeNode{typeNode}); ok {
				//fmt.Println(fmt.Sprintf("Seen is now %s", typeNode))
				redoSeen = true
				validTypes[0] = newType
				//fmt.Println(fmt.Sprintf("Reduced expect to %s", validTypes[0]))
			} else {
				return NewTypeError(typeNode, validTypes)
			}
		}
	}
	if redoSeen {
		return NewSetExpectance(validTypes).seen(check, typeNode)
	}

	//fmt.Println(fmt.Sprintf("Final seen is: %s   Final expect is: %s", typeNode, validTypes))
	//fmt.Println(fmt.Sprintf("Seen is now %s", typeNode))
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
func (exp TwiceSameExpectance) seen(check *TypeChecker, t TypeNode) TypeError {
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
func (exp RepeatExpectance) seen(check *TypeChecker, t TypeNode) TypeError {
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
func (exp AnyExpectance) seen(check *TypeChecker, t TypeNode) TypeError {
	return TypeError{}
}

// TypeChecker stores a stack of expectance. Seeing a type will pop it off from
// the stack, while expecting a type will push the type onto the stack.
// It can be frozen at a ProgramNode to prevent incorrect errors which can
// happen after some errors.
type TypeChecker struct {
	stack      []Expectance
	frozenNode ProgramNode
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
func (check *TypeChecker) seen(t TypeNode) TypeError {
	if check.frozen() {
		return TypeError{}
	}
	if len(check.stack) < 1 {
		fmt.Println("Internal type checker error")
		return TypeError{}
	}
	if DebugMode {
		fmt.Printf("Seen %s  -- p %T: &p=%p p=&i=%p \n", t, t, &t, t)
	}

	expectance := check.stack[len(check.stack)-1]
	check.stack = check.stack[:len(check.stack)-1]

	e := expectance.seen(check, t)
	//fmt.Printf("Seen done %s  -- p %T: &p=%p p=&i=%p \n", t, t, &t, t)
	return e
}

// StripType is used to remove the type of Arrays and Pairs, which is useful
// for comparing types.
func StripType(t TypeNode) TypeNode {
	switch t.(type) {
	case *ArrayTypeNode:
		return ArrayTypeNode{}
	case *PairTypeNode:
		return PairTypeNode{}
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
	if DebugMode {
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
	if DebugMode {
		fmt.Println("Expecting any")
	}
	check.stack = append(check.stack, NewAnyExpectance())
}

// expectTwiceSame adds a TwiceSameExpectance to the stack, if not frozen.
func (check *TypeChecker) expectTwiceSame(ex Expectance) {
	if check.frozen() {
		return
	}
	if DebugMode {
		fmt.Println("Expecting twice")
	}
	check.stack = append(check.stack, NewTwiceSameExpectance(ex))
}

// expectRepeatUntilForce adds a RepeatExpectance to the stack, if not
// frozen.
func (check *TypeChecker) expectRepeatUntilForce(t TypeNode) {
	if check.frozen() {
		return
	}
	if DebugMode {
		fmt.Printf("Expecting repeat %s\n", t)
	}
	check.stack = append(check.stack, NewRepeatExpectance(NewSetExpectance([]TypeNode{t})))
}

// expect adds a SetExpectance with the given TypeNode the only element in the set,
// if not frozen.
func (check *TypeChecker) expect(t TypeNode) {
	if check.frozen() {
		return
	}
	if DebugMode {
		fmt.Printf("Expecting %s  -- p %T: &p=%p p=&i=%p \n", t, t, &t, t)
	}
	check.expectSet([]TypeNode{t})
}

// expectSet adds an SetExpectance to the stack, if not frozen.
func (check *TypeChecker) expectSet(ts []TypeNode) {
	if check.frozen() {
		return
	}
	check.stack = append(check.stack, NewSetExpectance(ts))
}

// frozen returns if the typechecker is frozen or not.
func (check *TypeChecker) frozen() bool {
	return check.frozenNode != nil
}

// freeze freezes the type checker on a node.
func (check *TypeChecker) freeze(node ProgramNode) {
	if check.frozen() {
		return
	}
	if DebugMode {
		fmt.Printf("Frozen on %s\n", node)
	}
	check.frozenNode = node
}

// isSameNode Compares equality of ProgramNodes. As FunctionCallNode and ArrayElementNode
// are not comparable with the == operator, we define our own function that compares types first.
func isSameNode(n1 ProgramNode, n2 ProgramNode) bool {
	_, n1FunctionCall := n1.(*FunctionCallNode)
	_, n2FunctionCall := n2.(*FunctionCallNode)
	_, n1ArrayElement := n1.(*ArrayElementNode)
	_, n2ArrayElement := n2.(*ArrayElementNode)

	if (n1FunctionCall && n2FunctionCall) ||
		(n1ArrayElement && n2ArrayElement) {
		return true
	} else if !n1FunctionCall && !n2FunctionCall &&
		!n1ArrayElement && !n2ArrayElement {
		return n1 == n2
	} else {
		return false
	}
}

// unfreeze will unfreeze the checker, if the given node is the same as the frozen node.
func (check *TypeChecker) unfreeze(node ProgramNode) {
	if node == check.frozenNode { // isSameNode(node, check.frozenNode) {
		if DebugMode {
			fmt.Printf("Unfrozen on %s\n", node)
		}
		check.frozenNode = nil
	}
}
