package semantic_checker

import (
	"fmt"
	"gowacc/ast"
)

// Expectance is an interface used to store what type is expected.
type Expectance interface {
	seen(*TypeChecker, ast.TypeNode) TypeError
}

// SetExpectance is a struct that stores a set of acceptable types that can
// be seen. It implements Expectance.
type SetExpectance struct {
	set map[ast.TypeNode]bool
}

func NewSetExpectance(ts []ast.TypeNode) SetExpectance {
	set := make(map[ast.TypeNode]bool)
	for _, t := range ts {
		set[t] = true
	}
	return SetExpectance{
		set: set,
	}
}

// arrayCase handles the multiple options where we have seen an Array.
func arrayCase(check *TypeChecker, validTypes map[ast.TypeNode]bool, t ast.ArrayTypeNode) bool {
	_, match := validTypes[t]
	nilArray := ast.ArrayTypeNode{}
	expectingAnyArray := false
	matchOnAnyArray := t == nilArray
	var found ast.ArrayTypeNode
	for key := range validTypes {
		if StripType(key) == nilArray {
			found = key.(ast.ArrayTypeNode)
			expectingAnyArray = true
			break
		}
	}

	// ArrayLiteral case, so expect an unknown amount of expressions
	if matchOnAnyArray {
		if found.Dim == 1 {
			check.expectRepeatUntilForce(found.T)
		} else {
			check.expectRepeatUntilForce(ast.NewArrayTypeNode(found.T, found.Dim-1))
		}
		return true
	}
	return expectingAnyArray || match
}

// pairCase handles the multiple options where we have seen an pair.
func pairCase(check *TypeChecker, validTypes map[ast.TypeNode]bool, basePairMatch bool, t ast.PairTypeNode) bool {
	_, match := validTypes[t]
	_, matchBase := validTypes[ast.NewBaseTypeNode(ast.PAIR)]
	nilPair := ast.PairTypeNode{}
	matchOnAnyPair := t == nilPair
	var nilMatch ast.PairTypeNode
	expectingAnyPair := false
	for key := range validTypes {
		if StripType(key) == nilPair {
			nilMatch = key.(ast.PairTypeNode)
			expectingAnyPair = true
			break
		}
	}

	if matchOnAnyPair {
		if !expectingAnyPair {
			return false
		}

		// newpair case, expect first and second types of pair
		if !basePairMatch {
			check.expect(nilMatch.T2)
			check.expect(nilMatch.T1)
		}
		return true
	}

	return match || matchBase || expectingAnyPair
}

// seen is called when we have seen a SetExpectance.
func (exp SetExpectance) seen(check *TypeChecker, typeNode ast.TypeNode) TypeError {
	validTypes := exp.set

	switch t := typeNode.(type) {
	case ast.ArrayTypeNode:
		found := arrayCase(check, validTypes, t)
		if !found {
			return NewTypeError(t, validTypes)
		}
	case ast.PairTypeNode:
		found := pairCase(check, validTypes, false, t)
		if !found {
			return NewTypeError(t, validTypes)
		}
	case ast.BaseTypeNode:
		if t.T == ast.PAIR {
			_, found := validTypes[t]
			if !found {
				found = pairCase(check, validTypes, true, ast.PairTypeNode{})
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
	if t == nil {
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
	if ast.DEBUG_MODE {
		fmt.Printf("Seen %s\n", t)
	}

	expectance := check.stack[len(check.stack)-1]
	check.stack = check.stack[:len(check.stack)-1]

	return expectance.seen(check, t)
}

// StripType is used to remove the type of Arrays and Pairs, which is useful
// for comparing types.
func StripType(t ast.TypeNode) ast.TypeNode {
	switch t.(type) {
	case ast.ArrayTypeNode:
		return ast.ArrayTypeNode{}
	case ast.PairTypeNode:
		return ast.PairTypeNode{}
	}
	return t
}

// forcePop will force an expectance off the stack, useful for RepeatExpectance.
// It will only change the stack if it is not frozen.
func (check *TypeChecker) forcePop() {
	if check.frozen() {
		return
	}
	if ast.DEBUG_MODE {
		fmt.Println("Force pop")
	}
	if len(check.stack) < 1 {
		fmt.Println("Internal type checker error")
		return
	}
	check.stack = check.stack[:len(check.stack)-1]
}

// expectAny adds a AnyExpectance to the stack, if not frozen.
func (check *TypeChecker) expectAny() {
	if check.frozen() {
		return
	}
	if ast.DEBUG_MODE {
		fmt.Println("Expecting any")
	}
	check.stack = append(check.stack, NewAnyExpectance())
}

// expectTwiceSame adds a TwiceSameExpectance to the stack, if not frozen.
func (check *TypeChecker) expectTwiceSame(ex Expectance) {
	if check.frozen() {
		return
	}
	if ast.DEBUG_MODE {
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
	if ast.DEBUG_MODE {
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
	if ast.DEBUG_MODE {
		fmt.Printf("Expecting %s\n", t)
	}
	check.expectSet([]ast.TypeNode{t})
}

// expectSet adds an SetExpectance to the stack, if not frozen.
func (check *TypeChecker) expectSet(ts []ast.TypeNode) {
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
func (check *TypeChecker) freeze(node ast.ProgramNode) {
	if check.frozen() {
		return
	}
	if ast.DEBUG_MODE {
		fmt.Printf("Frozen on %s\n", node)
	}
	check.frozenNode = node
}

// isSameNode Compares equality of ProgramNodes. As FunctionCallNode and ArrayElementNode
// are not comparable with the == operator, we define our own function that compares types first.
func isSameNode(n1 ast.ProgramNode, n2 ast.ProgramNode) bool {
	_, n1FunctionCall := n1.(ast.FunctionCallNode)
	_, n2FunctionCall := n2.(ast.FunctionCallNode)
	_, n1ArrayElement := n1.(ast.ArrayElementNode)
	_, n2ArrayElement := n2.(ast.ArrayElementNode)

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
func (check *TypeChecker) unfreeze(node ast.ProgramNode) {
	if isSameNode(node, check.frozenNode) {
		if ast.DEBUG_MODE {
			fmt.Printf("Unfrozen on %s\n", node)
		}
		check.frozenNode = nil
	}
}
