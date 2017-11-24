package ast

import (
	"bytes"
	"fmt"
)

// SymbolTable is a struct that stores the symboltable and the currentScope that
// the program is in.
type SymbolTable struct {
	head           *SymbolTableNode
	currentScope   *SymbolTableNode
	functions      map[string]FunctionNode
	functionsOrder []string
}

// NewSymbolTable returns an initialised SymbolTable, with an empty scope. The
// current scope is initialised to the top level scope.
func NewSymbolTable() *SymbolTable {
	head := NewSymbolTableNode(nil)
	return &SymbolTable{
		head:           head,
		currentScope:   head,
		functions:      make(map[string]FunctionNode),
		functionsOrder: make([]string, 0),
	}
}

// SymbolTableNode is a struct that stores its scope, the scopes below it and a
// Pointer to the scope above itself.
type SymbolTableNode struct {
	scope       map[string]*IdentifierDeclaration
	childScopes []*SymbolTableNode
	parentScope *SymbolTableNode
	lastScope   int
	scopeSize   int
}

func NewSymbolTableNode(parentScope *SymbolTableNode) *SymbolTableNode {
	return &SymbolTableNode{
		scope:       make(map[string]*IdentifierDeclaration),
		childScopes: make([]*SymbolTableNode, 0, 10),
		parentScope: parentScope,
		lastScope:   -1,
	}
}

// IdentifierDeclaration stores the type and identifier for a symbol.
type IdentifierDeclaration struct {
	pos        Position
	t          TypeNode
	ident      IdentifierNode
	location   *Location
	isDeclared bool
}

func NewIdentifierDeclaration(programNode ProgramNode) *IdentifierDeclaration {
	switch node := programNode.(type) {
	case ParameterNode:
		return &IdentifierDeclaration{
			pos:        node.Pos,
			t:          node.T,
			ident:      node.Ident,
			isDeclared: false,
		}
	case DeclareNode:
		return &IdentifierDeclaration{
			pos:        node.Pos,
			t:          node.T,
			ident:      node.Ident,
			isDeclared: false,
		}
	default:
		return &IdentifierDeclaration{}
	}
}

func (dec *IdentifierDeclaration) AddLocation(location *Location) {
	dec.location = location
}

// MoveDownScope creates a new scope such that it is a chile of the currentscope,
// and then sets the currentScope to be the new scope.
func (table *SymbolTable) MoveDownScope() {
	newNode := NewSymbolTableNode(table.currentScope)
	table.currentScope.childScopes = append(table.currentScope.childScopes, newNode)
	table.currentScope = newNode
}

func (table *SymbolTable) MoveNextScope() {
	table.currentScope.lastScope++
	if len(table.currentScope.childScopes) > table.currentScope.lastScope {
		table.currentScope = table.currentScope.childScopes[table.currentScope.lastScope]
	} else {
		fmt.Println("Internal Error: no next scope, currentScope has ", len(table.currentScope.childScopes), " childscopes")
	}
}

// MoveUpScope will move the scope one level up if it exists.
func (table *SymbolTable) MoveUpScope() {
	if table.currentScope.parentScope != nil {
		table.currentScope = table.currentScope.parentScope
	}
}

// SearchForIdent will search for an identifier, first checking the currentScope
// and then will iterate through to the head scope. It will return false as its second return
// if the identifier is not in the currentScope or any parentScopes.
func (table *SymbolTable) SearchForIdent(identifier string) (*IdentifierDeclaration, bool) {
	for node := table.currentScope; node != nil; node = node.parentScope {
		node, ok := node.scope[identifier]
		if ok {
			return node, ok
		}
	}
	return &IdentifierDeclaration{}, false
}

func (table *SymbolTable) SearchForDeclaredIdent(identifier string) *IdentifierDeclaration {
	for node := table.currentScope; node != nil; node = node.parentScope {
		node, ok := node.scope[identifier]
		if ok {
			if node.isDeclared {
				return node
			}
		}
	}
	return &IdentifierDeclaration{}
}

// SearchForIdentInCurrentScope will search for an identifier, only in the
// currentScope. It will return false as its second return false
// if the identifier is not in the currentScope.
func (table *SymbolTable) SearchForIdentInCurrentScope(identifier string) (*IdentifierDeclaration, bool) {
	node, ok := table.currentScope.scope[identifier]
	return node, ok
}

// SearchForFunction will search for a function, returning false as its second
// return if it is not found.
func (table *SymbolTable) SearchForFunction(identifier string) (FunctionNode, bool) {
	node, ok := table.functions[identifier]
	return node, ok
}

// AddToScope will add an identifier to the currentScope.
func (table *SymbolTable) AddToScope(identifier string, programNode ProgramNode) {
	table.currentScope.scope[identifier] = NewIdentifierDeclaration(programNode)
}

// AddFunction will add a function to the symbolTable
func (table *SymbolTable) AddFunction(identifier string, node FunctionNode) {
	table.functions[identifier] = node
}

// Print will print a Node, and all of its parents
func (node SymbolTableNode) Print() {
	for _, ident := range node.scope {
		fmt.Printf("%s of type %s\n", ident.ident, ident.t)
	}
	fmt.Println("Parent Scope ---------------------")
	if node.parentScope != nil {
		node.parentScope.Print()
	}
}

// Print will print a symbolTable, relating from the currentScope. I.e. it will
// print the currentScope and all parentScopes, along with the Functions.
func (table *SymbolTable) Print() {
	fmt.Println("------- Begin Symbol table -------")
	fmt.Println("Functions ------------------------")
	for _, f := range table.functions {
		fmt.Printf("%s of type %s\n", f.Ident, f.T)
	}
	fmt.Println("Scopes ---------------------------")
	table.currentScope.Print()
	fmt.Println("-------- End Symbol table --------")
}

// String will return a string representation of the SymbolTable, from the
// top level scope down.
func (table *SymbolTable) String() string {
	var buf bytes.Buffer
	buf.WriteString("- Functions:\n")
	for _, f := range table.functions {
		buf.WriteString(fmt.Sprintf("  - %s %s(", f.T, f.Ident.String()[2:]))
		for i, p := range f.Params {
			if i == 0 {
				buf.WriteString(fmt.Sprintf("%s", p))
			} else {
				buf.WriteString(fmt.Sprintf(", %s", p))
			}
		}
		buf.WriteString(fmt.Sprintln(")"))
	}
	buf.WriteString("- Scopes:\n")
	for _, s := range table.head.childScopes {
		buf.WriteString(indent(fmt.Sprintf("%s\n", s), "  "))
	}
	return buf.String()
}

// String will return a string representation of the SymbolTableNode, and all
// of its children.
func (node *SymbolTableNode) String() string {
	var buf bytes.Buffer
	buf.WriteString("- Scope:\n")
	for _, s := range node.scope {
		buf.WriteString(fmt.Sprintf("  - Ident: %s, with type: %s\n", s.ident.Ident, s.t))
	}
	if len(node.childScopes) > 0 {
		buf.WriteString(" - With child scopes:\n")
		for _, s := range node.childScopes {
			buf.WriteString(indent(fmt.Sprintf("%s\n", s), "  "))
		}
	}
	return buf.String()
}
