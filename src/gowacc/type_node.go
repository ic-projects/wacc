package main

import (
	"bytes"
	"fmt"
)

// TypeNode is an empty interface for all types to implement.
type TypeNode interface {
	equals(TypeNode) bool
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

func toValue(typeNode TypeNode) TypeNode {
	switch t := typeNode.(type) {
	case *ArrayTypeNode:
		if inside, ok := toValue(t.T).(ArrayTypeNode); ok {
			t.Dim += inside.Dim
			t.T = inside.T
		}
		return *t
	case *PairTypeNode:
		return *t
	case *BaseTypeNode:
		return *t
	case *StructTypeNode:
		return *t
	case *NullTypeNode:
		return *t
	case *DynamicTypeNode:
		t2 := t.getValue()
		if _, ok := t2.(*DynamicTypeNode); ok {
			return t2
		} else {
			return toValue(t2)
		}
	case ArrayTypeNode:
		if inside, ok := toValue(t.T).(ArrayTypeNode); ok {
			t.Dim += inside.Dim
			t.T = inside.T
		}
		return t
	default:
		return t
	}
}

func validType(T TypeNode, i *IdentifierNode) GenericError {
	switch t := toValue(T).(type) {
	case ArrayTypeNode:
		if validType(t.T, i) != nil {
			return NewCustomError(i.Pos, fmt.Sprint("Unknown type for ident %s, has type %s", i.Ident, t))
		}
		return nil
	case DynamicTypeNode:
		fmt.Println("here")
		if len(t.T.poss) != 1 || validType(t.T.poss[0], i) != nil {
			return NewCustomError(i.Pos, fmt.Sprint("Unknown type for ident %s, has type %s", i.Ident, t))
		}
		return nil
	case PairTypeNode:
		if validType(t.T1, i) != nil || validType(t.T2, i) != nil {
			return NewCustomError(i.Pos, fmt.Sprint("Unknown type for ident %s, has type %s", i.Ident, t))
		}
		return nil
	default:
		return nil
	}
}

/******************** BASE TYPE ********************/

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
	} else if _, ok := toValue(t).(PairTypeNode); ok && node.T == PAIR {
		return true
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
	fmt.Println(fmt.Sprintf("array equals on %s and %s", node, t))
	if arr, ok := toValue(t).(ArrayTypeNode); ok {
		if arr2, ok := toValue(node).(ArrayTypeNode); ok {
			return (arr == ArrayTypeNode{}) || (arr2 == ArrayTypeNode{}) || (arr.Dim == arr2.Dim && arr2.T.equals(arr.T))
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
	} else if arr, ok := toValue(t).(BaseTypeNode); ok && arr.T == PAIR {
		return true
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

type DynamicTypeNode struct {
	T            *InternalDynamicType
	insidePair   bool
	arrayPointer *ArrayTypeNode
}

type InternalDynamicType struct {
	init     bool
	poss     []TypeNode
	watchers *[]*DynamicTypeNode
}

func NewDynamicTypeNode() *DynamicTypeNode {
	watchers := make([]*DynamicTypeNode, 0)
	node := &DynamicTypeNode{
		T: &InternalDynamicType{
			init:     false,
			poss:     make([]TypeNode, 0),
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
			poss:     make([]TypeNode, 0),
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
		if len(node.poss) == 1 {
			return fmt.Sprintf("%s", node.poss[0])
		}
		return fmt.Sprintf("%s", node.poss)
	}
	return fmt.Sprintf("unknown")
}

func (node DynamicTypeNode) String() string {
	return fmt.Sprintf("dynamic (%s) %p", node.T, node.T)
}

func (node DynamicTypeNode) equals(t TypeNode) bool {
	fmt.Println(fmt.Sprintf("WARNING, equals called on dynamic type: (%s) and (%s)", node, t))
	_, ok := node.reduceSet([]TypeNode{t})
	return ok
}

func (node *DynamicTypeNode) getValue() TypeNode {
	if len(node.T.poss) == 1 {
		t := node.T.poss[0]
		if arr, ok := t.(*ArrayTypeNode); ok {
			if arr.T == nil {
				arr := NewArrayTypeNode(NewDynamicTypeNode(), 1)
				arr.T.(*DynamicTypeNode).arrayPointer = arr
				node.T.poss[0] = arr
				t = node.T.poss[0]
			}
			//}
		} else if pair, ok := t.(*PairTypeNode); ok {
			if !node.insidePair {
				if pair.T2 == nil && pair.T1 == nil {
					node.T.poss[0] = NewPairTypeNode(NewDynamicTypeInsidePairNode(), NewDynamicTypeInsidePairNode())
					t = node.T.poss[0]
				} else if pair.T1 == nil {
					pair.T1 = NewDynamicTypeInsidePairNode()
				} else if pair.T2 == nil {
					pair.T2 = NewDynamicTypeInsidePairNode()
				}
			} else {
				node.T.poss[0] = NewBaseTypeNode(PAIR)
				t = node.T.poss[0]
			}
		}
		return t
	}
	return node
}

func (node *DynamicTypeNode) reduce(dyn *DynamicTypeNode) (TypeNode, bool) {
	fmt.Println(fmt.Sprintf("Special double dynamic reduction"))
	if node.T.init && dyn.T.init {
		if len(node.T.poss) == 1 && len(dyn.T.poss) == 1 {
			if node.T.poss[0].equals(dyn.T.poss[0]) {
				return node.getValue(), true
			} else {
				return nil, false
			}
		}
		t, ok := node.reduceSet(dyn.T.poss)
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

func (node *DynamicTypeNode) reduceSet(ts []TypeNode) (TypeNode, bool) {
	fmt.Println(fmt.Sprintf("Reducing %s with %s", node, ts))
	// Dynamic type saw another dynamic type
	if dyn, ok := ts[0].(*DynamicTypeNode); len(ts) == 1 && ok {
		return node.reduce(dyn)
	}
	if node.T.init {
		newSet := make([]TypeNode, 0)
		for _, t := range ts {
			for _, t2 := range node.T.poss {
				if t.equals(t2) {
					newSet = append(newSet, t2)
				}
			}
		}

		// Reduce leaves no possibilities
		if len(newSet) == 0 {
			fmt.Println(fmt.Sprintf("Reducing failed"))
			return nil, false
		}
		fmt.Println(fmt.Sprintf("Reducing success"))
		node.T.poss = newSet
	} else {
		fmt.Println(fmt.Sprintf("Reducing %s with %s", node, ts))
		fmt.Println(fmt.Sprintf("Reducing caused init"))
		node.T.init = true
		node.T.poss = ts
		fmt.Println(fmt.Sprintf("Reducing %s with %s", node, ts))
	}

	return node.getValue(), true
}
