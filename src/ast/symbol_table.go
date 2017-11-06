package ast

import (
	"fmt"
)

type SymbolTable struct {
	head         SymbolTableNode
	currentScope *SymbolTableNode
	functions    map[string]FunctionNode
}

func NewSymbolTable() *SymbolTable {
	head := NewSymbolTableNode(nil)
	return &SymbolTable{
		head:         head,
		currentScope: &head,
		functions:    make(map[string]FunctionNode),
	}
}

type SymbolTableNode struct {
	scope       map[string]IdentifierDeclaration
	childScopes []SymbolTableNode
	parentScope *SymbolTableNode
}

func NewSymbolTableNode(parentScope *SymbolTableNode) SymbolTableNode {
	return SymbolTableNode{
		scope:       make(map[string]IdentifierDeclaration),
		childScopes: make([]SymbolTableNode, 0, 10),
		parentScope: parentScope,
	}
}

type IdentifierDeclaration struct {
	pos   Position
	t     TypeNode
	ident IdentifierNode
}

func NewIdentifierDeclaration(programNode ProgramNode) IdentifierDeclaration {
	switch node := programNode.(type) {
	case ParameterNode:
		return IdentifierDeclaration{
			pos:   node.pos,
			t:     node.t,
			ident: node.ident,
		}
	case DeclareNode:
		return IdentifierDeclaration{
			pos:   node.pos,
			t:     node.t,
			ident: node.ident,
		}
	default:
		return IdentifierDeclaration{}
	}
}

func (table *SymbolTable) MoveDownScope() {
	newNode := NewSymbolTableNode(table.currentScope)
	table.currentScope.childScopes = append(table.currentScope.childScopes, newNode)
	table.currentScope = &newNode
}

func (table *SymbolTable) MoveUpScope() {
	if table.currentScope.parentScope != nil {
		table.currentScope = table.currentScope.parentScope
	}
}

func (table *SymbolTable) SearchForIdent(identifier string) (IdentifierDeclaration, bool) {
	for node := table.currentScope; node != nil; node = node.parentScope {
		node, ok := node.scope[identifier]
		if ok {
			return node, ok
		}
	}
	return IdentifierDeclaration{}, false
}

func (table *SymbolTable) SearchForIdentInCurrentScope(identifier string) (IdentifierDeclaration, bool) {
	node, ok := table.currentScope.scope[identifier]
	return node, ok
}

func (table *SymbolTable) SearchForFunction(identifier string) (FunctionNode, bool) {
	node, ok := table.functions[identifier]
	return node, ok
}

func (table *SymbolTable) AddToScope(identifier string, programNode ProgramNode) {
	table.currentScope.scope[identifier] = NewIdentifierDeclaration(programNode)
}

func (table *SymbolTable) AddFunction(identifier string, node FunctionNode) {
	table.functions[identifier] = node
}

func (node SymbolTableNode) Print() {
	for _, ident := range node.scope {
		fmt.Printf("%s of type %s\n", ident.ident, ident.t)
	}
	fmt.Println("Parent Scope ---------------------")
	if node.parentScope != nil {
		node.parentScope.Print()
	}
}

func (table *SymbolTable) Print() {
	fmt.Println("------- Begin Symbol table -------")
	fmt.Println("Functions ------------------------")
	for _, f := range table.functions {
		fmt.Printf("%s of type %s\n", f.ident, f.t)
	}
	fmt.Println("Scopes ---------------------------")
	table.currentScope.Print()
	fmt.Println("-------- End Symbol table --------")
}
