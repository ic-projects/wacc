package ast

import (
	"bytes"
	"fmt"
)

type TypeNode interface {
}

type BaseType int

const (
	INT BaseType = iota
	BOOL
	CHAR
	STRING
	PAIR
	VOID
)

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
		if DEBUG_MODE {
			return "basePair"
		}
		return ""
	case VOID:
		return "int"
	}
	return ""
}

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

type ArrayTypeNode struct {
	t        TypeNode
	dim      int
	isString bool
}

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

func NewStringArrayTypeNode() ArrayTypeNode {
	return ArrayTypeNode{
		t:        NewBaseTypeNode(CHAR),
		dim:      1,
		isString: true,
	}
}

func (node ArrayTypeNode) String() string {
	var buf bytes.Buffer
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
	if DEBUG_MODE {
		return fmt.Sprintf("pair(%s,%s)", node.t1, node.t2)
	}
	return fmt.Sprintf("%s%s", node.t1, node.t2)
}
