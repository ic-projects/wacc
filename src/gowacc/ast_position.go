package main

import (
	"fmt"
)

// Position stores the position of a node within the original code. The linenumber,
// column number and offset from the beginning of the file.
type Position struct {
	lineNumber int
	colNumber  int
	offset     int
}

// NewPosition builds an AST position.
func NewPosition(lineNumber int, colNumber int, offset int) Position {
	return Position{
		lineNumber: lineNumber,
		colNumber:  colNumber,
		offset:     offset,
	}
}

// LineNumber returns the line number of a Position.
func (p Position) LineNumber() int {
	return p.lineNumber
}

// ColNumber returns the column number of a Position.
func (p Position) ColNumber() int {
	colNum := p.colNumber
	if colNum != 0 {
		colNum--
	}
	return colNum
}

func (p Position) String() string {
	colNum := p.colNumber
	if colNum != 0 {
		colNum--
	}

	if DebugMode {
		offsetNum := p.offset
		if offsetNum != 0 {
			offsetNum--
		}
		return fmt.Sprintf(
			"line %d, column %d, offset %d",
			p.lineNumber,
			colNum,
			offsetNum,
		)
	}

	return fmt.Sprintf("%d:%d", p.lineNumber, colNum)
}
