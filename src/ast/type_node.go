package ast

import (
	"bytes"
	"fmt"
)

// TypeNode is an empty interface for all types to implement.
type TypeNode interface {
}

type BaseType int

const (
	INT    BaseType = iota // int
	BOOL                   // bool
	CHAR                   // char
	STRING                 // string, but internally represented as an array of chars
	PAIR                   // pair
	VOID                   // void, used for the return type of the main function, as you cannot return from main.
)

// String will return the string format of the BaseType. Pair returns empty
// string to match the refCompile ast printing. Void returns int for the same
// reason.
func (t BaseType) String() string {
	switch t {
	case INT:
		return "int"
	case BOOL:
		return "bool"
	case CHAR:
		return "char"
	case STRING:
		return "string"
	case PAIR:
		return "basePair"
	case VOID:
		return "int"
	}
	return ""
}

// BaseTypeNode is a struct that stores a BaseType.
type BaseTypeNode struct {
	t BaseType
}

func NewBaseTypeNode(t BaseType) BaseTypeNode {
	return BaseTypeNode{
		t: t,
	}
}

func (node BaseTypeNode) String() string {
	return fmt.Sprintf("%s", node.t)
}

// ArrayTypeNode stores the type, and dimension of the array. It stores if it is
// a string additionally to distinguish between a char array and a string.
type ArrayTypeNode struct {
	t        TypeNode
	dim      int
	isString bool
}

// NewArrayTypeNode returns an initialised ArrayTypeNode. If the type provided
// is an array, it will increase the dimensions of the given array, and return
// it.
func NewArrayTypeNode(t TypeNode, dim int) ArrayTypeNode {
	if array, ok := t.(ArrayTypeNode); ok {
		array.dim += dim
		return array
	}
	return ArrayTypeNode{
		t:        t,
		dim:      dim,
		isString: false,
	}
}

// NewStringArrayTypeNode returns an initialised ArrayTypeNode for a string.
func NewStringArrayTypeNode() ArrayTypeNode {
	return ArrayTypeNode{
		t:        NewBaseTypeNode(CHAR),
		dim:      1,
		isString: true,
	}
}

func (node ArrayTypeNode) String() string {
	var buf bytes.Buffer
	if node == (ArrayTypeNode{}) {
		return fmt.Sprintf("array")
	}
	if node.isString {
		buf.WriteString(fmt.Sprintf("string"))
		for i := 0; i < node.dim-1; i++ {
			buf.WriteString("[]")
		}
	} else {
		buf.WriteString(fmt.Sprintf("%s", node.t))
		for i := 0; i < node.dim; i++ {
			buf.WriteString("[]")
		}
	}
	return buf.String()
}

// PairTypeNode is a struct that stores the types of the first and second
// elements of the pair.
//
// E.g.
//
//  pair(int, int)
type PairTypeNode struct {
	t1 TypeNode
	t2 TypeNode
}

func NewPairTypeNode(t1 TypeNode, t2 TypeNode) PairTypeNode {
	return PairTypeNode{
		t1: t1,
		t2: t2,
	}
}

func (node PairTypeNode) String() string {
	if node == (PairTypeNode{}) {
		return fmt.Sprintf("pair")
	}
	return fmt.Sprintf("pair(%s, %s)", node.t1, node.t2)
}
