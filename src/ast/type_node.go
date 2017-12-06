package ast

import (
	"bytes"
	"fmt"
)

// TypeNode is an empty interface for all types to implement.
type TypeNode interface {
	fmt.Stringer
	Equals(TypeNode) bool
}

/**************** TYPE NODE HELPER FUNCTIONS ****************/

// SizeOf returns the size, in bytes, required to store an element of the given
// type.
func SizeOf(t TypeNode) int {
	switch node := ToValue(t).(type) {
	case BaseTypeNode:
		switch node.T {
		case CHAR, BOOL:
			return 1
		}
	}
	return 4
}

func ToValue(typeNode TypeNode) TypeNode {
	switch t := typeNode.(type) {
	case *ArrayTypeNode:
		return *t
	case *PairTypeNode:
		return *t
	case *BaseTypeNode:
		return *t
	case *PointerTypeNode:
		return *t
	case *StructTypeNode:
		return *t
	case *NullTypeNode:
		return *t
	case *DynamicTypeNode:
		t2 := t.getValue()
		if _, ok := t2.(*DynamicTypeNode); ok {
			return t2
		}
		return ToValue(t2)
	case ArrayTypeNode:
		return t
	default:
		return t
	}
}

/******************** BASE TYPE ********************/

// BaseType is a representation of a simple WACC type, that may form a type by
// itself, or as part of a more complex array or pair type.
type BaseType int

const (
	// INT int type
	INT BaseType = iota
	// BOOL bool type
	BOOL
	// CHAR char type
	CHAR
	// STRING string type, but internally represented as an array of chars
	STRING
	// PAIR pair base type, used for nested pairs
	PAIR
	// VOID void type, used for the return type of the main function,
	// as you cannot return from main.
	VOID
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

// NewBaseTypeNode builds a BaseTypeNode.
func NewBaseTypeNode(t BaseType) *BaseTypeNode {
	return &BaseTypeNode{
		T: t,
	}
}

func (node BaseTypeNode) String() string {
	return node.T.String()
}

func (node BaseTypeNode) Equals(t TypeNode) bool {
	if arr, ok := ToValue(t).(BaseTypeNode); ok {
		return node.T == arr.T
	} else if _, ok := ToValue(t).(PairTypeNode); ok && node.T == PAIR {
		return true
	}
	return false
}

/**************** ARRAY TYPE NODE ****************/

// ArrayTypeNode stores the type, and dimension of the array. It stores if it is
// a string additionally to distinguish between a char array and a string.
type ArrayTypeNode struct {
	T        TypeNode
	IsString bool
}

// NewArrayTypeNode returns an initialised ArrayTypeNode. If the type provided
// is an array, it will increase the dimensions of the given array, and return
// it.
func NewArrayTypeNode(t TypeNode) *ArrayTypeNode {
	return &ArrayTypeNode{
		T:        t,
		IsString: false,
	}
}

func NewArrayTypeDimNode(t TypeNode, dim int) *ArrayTypeNode {
	var arr *ArrayTypeNode
	arr = &ArrayTypeNode{
		T:        t,
		IsString: false,
	}
	for i := 1; i < dim; i++ {
		arr = &ArrayTypeNode{
			T:        arr,
			IsString: false,
		}
	}
	return arr
}

// NewStringArrayTypeNode returns an initialised ArrayTypeNode for a string.
func NewStringArrayTypeNode() *ArrayTypeNode {
	return &ArrayTypeNode{
		T:        NewBaseTypeNode(CHAR),
		IsString: true,
	}
}

// GetDimElement will return the type of the element in the array at depth dim,
// e.g. an array of char[][] has char[][] at depth 0, char[] at depth 1 and
// char at depth 2
func (node ArrayTypeNode) GetDimElement(dim int) TypeNode {
	var t TypeNode
	t = node
	for i := 0; i < dim; i++ {
		t = ToValue(t).(ArrayTypeNode).T
	}
	return t
}

func (node ArrayTypeNode) String() string {
	if node == (ArrayTypeNode{}) {
		return fmt.Sprintf("array")
	}
	var buf bytes.Buffer
	if t, ok := ToValue(node.T).(BaseTypeNode); ok &&
		t.T == CHAR && node.IsString {
		buf.WriteString("string[]")
	} else {
		buf.WriteString(node.T.String())
		buf.WriteString("[]")
	}
	return buf.String()
}

func (node ArrayTypeNode) Equals(t TypeNode) bool {
	if arr, ok := ToValue(t).(ArrayTypeNode); ok {
		if arr2, ok := ToValue(node).(ArrayTypeNode); ok {
			return (arr == ArrayTypeNode{}) || (arr2 == ArrayTypeNode{}) || arr2.T.Equals(arr.T)
		}
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

// NewPairTypeNode builds a PairTypeNode.
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

func (node PairTypeNode) Equals(t TypeNode) bool {
	if arr, ok := ToValue(t).(PairTypeNode); ok {
		return node.T1.Equals(arr.T1) && node.T2.Equals(arr.T2)
	} else if arr, ok := ToValue(t).(BaseTypeNode); ok && arr.T == PAIR {
		return true
	}
	return false
}

/**************** NULL TYPE NODE ****************/

// NullTypeNode is an empty struct used to represent a null type.
type NullTypeNode struct {
}

// NewNullTypeNode builds a NullTypeNode.
func NewNullTypeNode() *NullTypeNode {
	return &NullTypeNode{}
}

func (node NullTypeNode) Equals(t TypeNode) bool {
	_, ok := ToValue(t).(NullTypeNode)
	return ok
}

func (node NullTypeNode) String() string {
	return "null"
}

/**************** STRUCT TYPE NODE ****************/

// StructTypeNode stores a user-defined type.
type StructTypeNode struct {
	Ident string
	poss  []string
}

// NewStructTypeNode builds a StructTypeNode.
func NewStructTypeNode(i *IdentifierNode) *StructTypeNode {
	return &StructTypeNode{
		Ident: i.Ident,
		poss:  make([]string, 0),
	}
}

func NewStrucDynamictTypeNode() *StructTypeNode {
	return &StructTypeNode{
		Ident: "",
		poss:  make([]string, 0),
	}
}

func (node StructTypeNode) String() string {
	return fmt.Sprintf("struct %s", node.Ident)
}

func (node StructTypeNode) Equals(t TypeNode) bool {
	if arr, ok := ToValue(t).(StructTypeNode); ok {
		/*if arr.Ident == "" && node.Ident == "" {
			newSet := make([]string, 0)
			for _, t := range arr.Poss {
				for _, t2 := range node.Poss {
					if t == t2 {
						newSet = append(newSet, t)
					}
				}
			}

			arr.Poss = newSet
			node.Poss = newSet
			if len(newSet) == 1 {
				node.Ident = newSet[0]
				arr.Ident = newSet[0]
			}
			return len(newSet) > 0
		}
		if arr.Ident == "" {
			for _, t := range arr.Poss {
				if t == node.Ident {
					arr.Ident = t
					return true
				}
			}
		}
		if node.Ident == "" {
			for _, t := range node.Poss {
				if t == arr.Ident {
					node.Ident = t
					return true
				}
			}
		}*/
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

func (node PointerTypeNode) Equals(t TypeNode) bool {
	if arr, ok := ToValue(t).(PointerTypeNode); ok {
		return arr.T.Equals(node.T)
	}
	return false
}

type DynamicTypeNode struct {
	T            *InternalDynamicType
	insidePair   bool
	arrayPointer *ArrayTypeNode
}

type InternalDynamicType struct {
	init     bool
	Poss     []TypeNode
	watchers *[]*DynamicTypeNode
}

func NewDynamicTypeNode() *DynamicTypeNode {
	watchers := make([]*DynamicTypeNode, 0)
	node := &DynamicTypeNode{
		T: &InternalDynamicType{
			init:     false,
			Poss:     make([]TypeNode, 0),
			watchers: &watchers,
		},
		insidePair: false,
	}
	*node.T.watchers = append(*node.T.watchers, node)
	return node
}

func NewDynamicTypeInsidePairNode() *DynamicTypeNode {
	watchers := make([]*DynamicTypeNode, 0)
	node := &DynamicTypeNode{
		T: &InternalDynamicType{
			init:     false,
			Poss:     make([]TypeNode, 0),
			watchers: &watchers,
		},
		insidePair: true,
	}
	*node.T.watchers = append(*node.T.watchers, node)
	return node
}

func (node *DynamicTypeNode) changeToWatch(other *DynamicTypeNode) {
	*other.T.watchers = append(*other.T.watchers, *node.T.watchers...)
	for _, watcher := range *node.T.watchers {
		watcher.T = node.T
	}
	node.T = other.T
}

func (node InternalDynamicType) String() string {
	if node.init {
		if len(node.Poss) == 1 {
			return fmt.Sprintf("%s", node.Poss[0])
		}
		return fmt.Sprintf("%s", node.Poss)
	}
	return fmt.Sprintf("unknown")
}

func (node DynamicTypeNode) String() string {
	return fmt.Sprintf("dynamic (%s) %p", node.T, node.T)
}

func (node DynamicTypeNode) Equals(t TypeNode) bool {
	_, ok := node.ReduceSet([]TypeNode{t})
	return ok
}

func (node *DynamicTypeNode) getValue() TypeNode {
	if len(node.T.Poss) == 1 {
		t := node.T.Poss[0]
		if arr, ok := t.(*ArrayTypeNode); ok {
			if arr.T == nil {
				arr := NewArrayTypeNode(NewDynamicTypeNode())
				arr.T.(*DynamicTypeNode).arrayPointer = arr
				node.T.Poss[0] = arr
				t = node.T.Poss[0]
			}
			//}
		} else if pair, ok := t.(*PairTypeNode); ok {
			if !node.insidePair {
				if pair.T2 == nil && pair.T1 == nil {
					node.T.Poss[0] = NewPairTypeNode(
						NewDynamicTypeInsidePairNode(),
						NewDynamicTypeInsidePairNode())
					t = node.T.Poss[0]
				} else if pair.T1 == nil {
					pair.T1 = NewDynamicTypeInsidePairNode()
				} else if pair.T2 == nil {
					pair.T2 = NewDynamicTypeInsidePairNode()
				}
			} else {
				node.T.Poss[0] = NewBaseTypeNode(PAIR)
				t = node.T.Poss[0]
			}
		}
		return t
	}
	return node
}

func (node *DynamicTypeNode) reduce(dyn *DynamicTypeNode) (TypeNode, bool) {
	//fmt.Println(fmt.Sprintf("Special double dynamic reduction"))
	if node.T.init && dyn.T.init {
		if len(node.T.Poss) == 1 && len(dyn.T.Poss) == 1 {
			if node.T.Poss[0].Equals(dyn.T.Poss[0]) {
				return node.getValue(), true
			} else {
				return nil, false
			}
		}
		t, ok := node.ReduceSet(dyn.T.Poss)
		if ok {
			dyn.changeToWatch(node)
		}
		return t, ok
	} else if node.T.init && !dyn.T.init {
		dyn.changeToWatch(node)
	} else if !node.T.init && dyn.T.init {
		node.changeToWatch(dyn)
	} else {
		node.changeToWatch(dyn)
		dyn.changeToWatch(node)
		return nil, true
	}
	dyn.changeToWatch(node)
	return node.getValue(), true
}

func (node *DynamicTypeNode) ReduceSet(ts []TypeNode) (TypeNode, bool) {
	//fmt.Println(fmt.Sprintf("Reducing %s with %s", node, ts))
	// Dynamic type saw another dynamic type
	if dyn, ok := ts[0].(*DynamicTypeNode); len(ts) == 1 && ok {
		return node.reduce(dyn)
	}
	if node.T.init {
		newSet := make([]TypeNode, 0)
		for _, t := range ts {
			for _, t2 := range node.T.Poss {
				if t.Equals(t2) {
					newSet = append(newSet, t2)
				}
			}
		}

		// Reduce leaves no possibilities
		if len(newSet) == 0 {
			//fmt.Println(fmt.Sprintf("Reducing failed"))
			return nil, false
		}
		//fmt.Println(fmt.Sprintf("Reducing success"))
		node.T.Poss = newSet
	} else {
		//fmt.Println(fmt.Sprintf("Reducing caused init"))
		node.T.init = true
		node.T.Poss = ts
	}

	return node.getValue(), true
}
