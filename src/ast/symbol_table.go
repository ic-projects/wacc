package ast

import (
	"bytes"
	"fmt"
	"utils"
)

/**************** SYMBOL TABLE ****************/

// SymbolTable is a struct that stores the symboltable and the CurrentScope that
// the program is in.
type SymbolTable struct {
	Head           *SymbolTableNode
	CurrentScope   *SymbolTableNode
	Functions      map[string]*FunctionNode
	Structs        map[string]*StructNode
	functionsOrder []string
}

// NewSymbolTable returns an initialised SymbolTable, with an empty Scope. The
// current Scope is initialised to the top level Scope.
func NewSymbolTable() *SymbolTable {
	head := NewSymbolTableNode(nil)
	return &SymbolTable{
		Head:           head,
		CurrentScope:   head,
		Functions:      make(map[string]*FunctionNode),
		Structs:        make(map[string]*StructNode),
		functionsOrder: make([]string, 0),
	}
}

/**************** SYMBOL TABLE NODE ****************/

// SymbolTableNode is a struct that stores its Scope, the scopes below it and a
// Pointer to the Scope above itself.
type SymbolTableNode struct {
	Scope       map[string]*IdentifierDeclaration
	ChildScopes []*SymbolTableNode
	ParentScope *SymbolTableNode
	lastScope   int
	ScopeSize   int
}

// NewSymbolTableNode builds a SymbolTableNode.
func NewSymbolTableNode(parentScope *SymbolTableNode) *SymbolTableNode {
	return &SymbolTableNode{
		Scope:       make(map[string]*IdentifierDeclaration),
		ChildScopes: make([]*SymbolTableNode, 0, 10),
		ParentScope: parentScope,
		lastScope:   -1,
	}
}

/**************** IDENTIFIER DECLARATION ****************/

// IdentifierDeclaration stores the type and identifier for a symbol.
type IdentifierDeclaration struct {
	Pos        utils.Position
	T          TypeNode
	Ident      *IdentifierNode
	Location   *utils.Location
	IsDeclared bool
	Value      ExpressionNode
	HasValue   bool
}

// NewIdentifierDeclaration builds an Identifier Declaration.
func NewIdentifierDeclaration(programNode ProgramNode) *IdentifierDeclaration {
	switch node := programNode.(type) {
	case *ParameterNode:
		return &IdentifierDeclaration{
			Pos:        node.Pos,
			T:          node.T,
			Ident:      node.Ident,
			IsDeclared: false,
			HasValue:   false,
		}
	case *DeclareNode:
		a := &IdentifierDeclaration{
			Pos:        node.Pos,
			T:          node.T,
			Ident:      node.Ident,
			IsDeclared: false,
			HasValue:   false,
		}
		return a
	default:
		return &IdentifierDeclaration{}
	}
}

// AddLocation will add a location to a declaration.
func (dec *IdentifierDeclaration) AddLocation(location *utils.Location) {
	dec.Location = location
}

func (dec *IdentifierDeclaration) SetValue(value ExpressionNode) {
	dec.Value = value
	dec.HasValue = true
}

func (dec *IdentifierDeclaration) RemoveValue() {
	dec.HasValue = false
}

/**************** MOVING SCOPE HELPER FUNCTIONS ****************/

func (table *SymbolTable) Reset() {
	// TODO
}

// MoveDownScope creates a new Scope such that it is a chile of the
// currentscope, and then sets the CurrentScope to be the new Scope.
func (table *SymbolTable) MoveDownScope() {
	newNode := NewSymbolTableNode(table.CurrentScope)
	table.CurrentScope.ChildScopes = append(
		table.CurrentScope.ChildScopes,
		newNode,
	)
	table.CurrentScope = newNode
}

// MoveNextScope builds a SymbolTableNode
func (table *SymbolTable) MoveNextScope() {
	table.CurrentScope.lastScope++
	if len(table.CurrentScope.ChildScopes) > table.CurrentScope.lastScope {
		table.CurrentScope =
			table.CurrentScope.ChildScopes[table.CurrentScope.lastScope]
	} else {
		fmt.Println(
			"Internal Error: no next Scope, CurrentScope has ",
			len(table.CurrentScope.ChildScopes),
			" childscopes",
		)
	}
}

// MoveUpScope will move the Scope one level up if it exists.
func (table *SymbolTable) MoveUpScope() {
	if table.CurrentScope.ParentScope != nil {
		table.CurrentScope = table.CurrentScope.ParentScope
	}
}

/**************** SEARCHING HELPER FUNCTIONS ****************/

// SearchForIdent will search for an identifier, first checking the CurrentScope
// and then will iterate through to the Head Scope. It will return false as its
// second return if the identifier is not in the CurrentScope or any
// parentScopes.
func (table *SymbolTable) SearchForIdent(
	identifier string,
) (*IdentifierDeclaration, bool) {
	for node := table.CurrentScope; node != nil; node = node.ParentScope {
		node, ok := node.Scope[identifier]
		if ok {
			return node, ok
		}
	}
	return &IdentifierDeclaration{}, false
}

// SearchForDeclaredIdent will search for an identifier that has the IsDeclared
// flag to be true. It will search the currentScope first, before checking
// parentScopes.
func (table *SymbolTable) SearchForDeclaredIdent(
	identifier string,
) *IdentifierDeclaration {
	for node := table.CurrentScope; node != nil; node = node.ParentScope {
		node, ok := node.Scope[identifier]
		if ok {
			if node.IsDeclared {
				return node
			}
		}
	}
	return &IdentifierDeclaration{}
}

// SearchForIdentInCurrentScope will search for an identifier, only in the
// CurrentScope. It will return false as its second return false
// if the identifier is not in the CurrentScope.
func (table *SymbolTable) SearchForIdentInCurrentScope(
	identifier string,
) (*IdentifierDeclaration, bool) {
	node, ok := table.CurrentScope.Scope[identifier]
	return node, ok
}

// SearchForFunction will search for a function, returning false as its second
// return if it is not found.
func (table *SymbolTable) SearchForFunction(
	identifier string,
) (*FunctionNode, bool) {
	node, ok := table.Functions[identifier]
	return node, ok
}

// SearchForStruct will search for a struct, returning false as its second
// return if not found.
func (table *SymbolTable) SearchForStruct(
	identifier string,
) (*StructNode, bool) {
	node, ok := table.Structs[identifier]
	return node, ok
}

// SearchForStructByUsage will search for a struct, returning false as its
// second return if not found. This version finds a struct from a member.
func (table *SymbolTable) SearchForStructByUsage(
	usage string) ([]TypeNode, []*StructNode, []*StructInternalNode) {
	set := make([]TypeNode, 0)
	structSet := make([]*StructNode, 0)
	structInternalSet := make([]*StructInternalNode, 0)
	for _, s := range table.Structs {
		if t, ok := s.TypesMap[usage]; ok {
			set = append(set, s.Types[t].T)
			structSet = append(structSet, s)
			structInternalSet = append(structInternalSet, s.Types[t])
		}
	}
	return set, structSet, structInternalSet
}

/**************** ADDING HELPER FUNCTIONS ****************/

// AddToScope will add an identifier to the CurrentScope.
func (table *SymbolTable) AddToScope(
	identifier string,
	programNode ProgramNode,
) {
	table.CurrentScope.Scope[identifier] = NewIdentifierDeclaration(programNode)
}

// AddFunction will add a function to the symbolTable
func (table *SymbolTable) AddFunction(identifier string, node *FunctionNode) {
	table.Functions[identifier] = node
}

// AddStruct will add a struct to the symbolTable
func (table *SymbolTable) AddStruct(identifier string, node *StructNode) {
	table.Structs[identifier] = node
}

/******************** PRINTING HELPER FUNCTIONS ********************/

// Print will print a Node, and all of its parents
func (node SymbolTableNode) Print() {
	for _, ident := range node.Scope {
		fmt.Printf("%s of type %s\n", ident.Ident, ident.T)
	}
	fmt.Println("Parent Scope ---------------------")
	if node.ParentScope != nil {
		node.ParentScope.Print()
	}
}

// Print will print a symbolTable, relating from the CurrentScope. I.e. it will
// print the CurrentScope and all parentScopes, along with the Functions.
func (table *SymbolTable) Print() {
	fmt.Println("------- Begin Symbol table -------")
	fmt.Println("Structs ------------------------")
	for _, f := range table.Structs {
		fmt.Printf("%s", f)
	}
	fmt.Println("Functions ------------------------")
	for _, f := range table.Functions {
		fmt.Printf("%s of type %s\n", f.Ident, f.T)
	}
	fmt.Println("Scopes ---------------------------")
	table.CurrentScope.Print()
	fmt.Println("-------- End Symbol table --------")
}

// String will return a string representation of the SymbolTable, from the
// top level Scope down.
func (table *SymbolTable) String() string {
	var buf bytes.Buffer
	buf.WriteString("- Structs:\n")
	for _, f := range table.Structs {
		buf.WriteString(f.String())
	}
	buf.WriteString("- Functions:\n")
	for _, f := range table.Functions {
		buf.WriteString(fmt.Sprintf("  - %s %s(", f.T, f.Ident.String()[2:]))
		for i, p := range f.Params {
			if i == 0 {
				buf.WriteString(p.String())
			} else {
				buf.WriteString(fmt.Sprintf(", %s", p))
			}
		}
		buf.WriteString(fmt.Sprintln(")"))
	}
	buf.WriteString("- Scopes:\n")
	for _, s := range table.Head.ChildScopes {
		buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", s), "  "))
	}
	return buf.String()
}

// String will return a string representation of the SymbolTableNode, and all
// of its children.
func (node *SymbolTableNode) String() string {
	var buf bytes.Buffer
	buf.WriteString("- Scope:\n")
	for _, s := range node.Scope {
		buf.WriteString(fmt.Sprintf(
			"  - Ident: %s, with type: %s\n",
			s.Ident.Ident, s.T,
		))
	}
	if len(node.ChildScopes) > 0 {
		buf.WriteString(" - With child scopes:\n")
		for _, s := range node.ChildScopes {
			buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", s), "  "))
		}
	}
	return buf.String()
}
