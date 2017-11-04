package ast

type SymbolTable struct {
	head         SymbolTableNode
	currentScope *SymbolTableNode
}

func NewSymbolTable() SymbolTable {
	head := NewSymbolTableNode(nil)
	return SymbolTable{
		head:         head,
		currentScope: &head,
	}
}

type SymbolTableNode struct {
	scope       map[string]ProgramNode
	childScopes []SymbolTableNode
	parentScope *SymbolTableNode
}

func NewSymbolTableNode(parentScope *SymbolTableNode) SymbolTableNode {
	return SymbolTableNode{
		scope:       make(map[string]ProgramNode),
		childScopes: make([]SymbolTableNode, 0, 10),
		parentScope: parentScope,
	}
}

func (table SymbolTable) MoveDownScope() SymbolTable {
	newNode := NewSymbolTableNode(table.currentScope)
	table.currentScope.childScopes = append(table.currentScope.childScopes, newNode)
	table.currentScope = &newNode
	return table
}

func (table SymbolTable) MoveUpScope() SymbolTable {
	table.currentScope = table.currentScope.parentScope
	return table
}

func (table SymbolTable) SearchFor(identifier string) (ProgramNode, bool) {
	for node := table.currentScope; node != nil; node = node.parentScope {
		node, ok := node.scope[identifier]
		if ok {
			return node, ok
		}
	}
	return nil, false
}

func (table SymbolTable) AddToScope(identifier string, node ProgramNode) SymbolTable {
	table.currentScope.scope[identifier] = node
	return table
}
