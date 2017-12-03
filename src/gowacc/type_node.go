package main

import (
	"bytes"
	"fmt"
)

// TypeNode is an empty interface for all types to implement.
type TypeNode interface {
	equals(TypeNode) bool
	String() string
}

/**************** TYPE NODE HELPER FUNCTIONS ****************/

func SizeOf(t TypeNode) int {
	switch node := toValue(t).(type) {
	case BaseTypeNode:
		switch node.T {
		case CHAR, BOOL:
			return 1
		}
	}
	return 4
}

/**************** BASE TYPE ****************/

type BaseType int

const (
	INT    BaseType = iota // int
	BOOL                   // bool
	CHAR                   // char
	STRING                 // string, but internally represented as an array of chars
	PAIR                   // pair
	VOID                   // void, used for the return type of the main function, as you cannot return from main.
)

// String will return the string format of the BaseType. Void returns "int" to
// match the refCompile ast printing.
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

/**************** BASE TYPE NODE ****************/

// BaseTypeNode is a struct that stores a BaseType.
type BaseTypeNode struct {
	T BaseType
}

func NewBaseTypeNode(t BaseType) *BaseTypeNode {
	return &BaseTypeNode{
		T: t,
	}
}

func (node BaseTypeNode) String() string {
	return fmt.Sprintf("%s", node.T)
}

func (node BaseTypeNode) equals(t TypeNode) bool {
	if arr, ok := toValue(t).(BaseTypeNode); ok {
		return node.T == arr.T
	}
	return false
}

/**************** ARRAY TYPE NODE ****************/

// ArrayTypeNode stores the type, and dimension of the array. It stores if it is
// a string additionally to distinguish between a char array and a string.
type ArrayTypeNode struct {
	T        TypeNode
	Dim      int
	IsString bool
}

// NewArrayTypeNode returns an initialised ArrayTypeNode. If the type provided
// is an array, it will increase the dimensions of the given array, and return
// it.
func NewArrayTypeNode(t TypeNode, dim int) *ArrayTypeNode {
	if array, ok := t.(*ArrayTypeNode); ok {
		array.Dim += dim
		return array
	}
	return &ArrayTypeNode{
		T:        t,
		Dim:      dim,
		IsString: false,
	}
}

// NewStringArrayTypeNode returns an initialised ArrayTypeNode for a string.
func NewStringArrayTypeNode() *ArrayTypeNode {
	return &ArrayTypeNode{
		T:        NewBaseTypeNode(CHAR),
		Dim:      1,
		IsString: true,
	}
}

func (node ArrayTypeNode) String() string {
	if node == (ArrayTypeNode{}) {
		return fmt.Sprintf("array")
	}
	var buf bytes.Buffer
	if node.IsString {
		buf.WriteString(fmt.Sprintf("string"))
		for i := 0; i < node.Dim-1; i++ {
			buf.WriteString("[]")
		}
	} else {
		buf.WriteString(fmt.Sprintf("%s", node.T))
		for i := 0; i < node.Dim; i++ {
			buf.WriteString("[]")
		}
	}
	return buf.String()
}

func (node ArrayTypeNode) equals(t TypeNode) bool {
	if arr, ok := toValue(t).(ArrayTypeNode); ok {
		return arr.Dim == node.Dim && node.T.equals(arr.T)
	}
	return false
}

/**************** PAIR TYPE NODE ****************/

// PairTypeNode is a struct that stores the types of the first and second
// elements of the pair.
//
// E.g.
//
//  pair(int, int)
type PairTypeNode struct {
	T1 TypeNode
	T2 TypeNode
}

func NewPairTypeNode(t1 TypeNode, t2 TypeNode) *PairTypeNode {
	return &PairTypeNode{
		T1: t1,
		T2: t2,
	}
}

func (node PairTypeNode) String() string {
	if node == (PairTypeNode{}) {
		return fmt.Sprintf("pair")
	}
	return fmt.Sprintf("pair(%s, %s)", node.T1, node.T2)
}

func (node PairTypeNode) equals(t TypeNode) bool {
	if arr, ok := toValue(t).(PairTypeNode); ok {
		return node.T1.equals(arr.T1) && node.T2.equals(arr.T2)
	}
	return false
}

type StructTypeNode struct {
	Ident string
}

type NullTypeNode struct {
}

func NewNullTypeNode() *NullTypeNode {
	return &NullTypeNode{}
}

func (node NullTypeNode) equals(t TypeNode) bool {
	_, ok := toValue(t).(NullTypeNode)
	return ok
}

func (node NullTypeNode) String() string {
	return "null"
}

func NewStructTypeNode(i *IdentifierNode) *StructTypeNode {
	return &StructTypeNode{
		Ident: i.Ident,
	}
}

func (node StructTypeNode) String() string {
	return fmt.Sprintf("struct %s", node.Ident)
}

func (node StructTypeNode) equals(t TypeNode) bool {
	if arr, ok := toValue(t).(StructTypeNode); ok {
		return arr.Ident == node.Ident
	}
	return false
}

type PointerTypeNode struct {
	T TypeNode
}

func NewPointerTypeNode(t TypeNode) *PointerTypeNode {
	return &PointerTypeNode{
		T: t,
	}
}

func (node PointerTypeNode) String() string {
	return fmt.Sprintf("%s *", node.T.String())
}

func (node PointerTypeNode) equals(t TypeNode) bool {
	if arr, ok := toValue(t).(PointerTypeNode); ok {
		return arr.T == node.T
	}
	return false
}
