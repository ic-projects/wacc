package main

// Predefined is an interface that defines a function or data that can be added
// to the CodeGenerators assembly.
type Predefined interface {
	add(*CodeGenerator, *Library)
}

/**************** PREDEFINED FUNCTIONS ****************/

// PreFunction defines a predefined function. The function has a name, the body
// of function containing the code and a dependency list of other
// LibraryFunctions required for the function to work.
type PreFunction struct {
	name         string
	body         []string
	dependencies []LibraryFunction
}

// NewPreFunction returns a PreFunction struct containing the name, body and
// dependencies of the function.
func NewPreFunction(
	name string,
	dependencies []LibraryFunction,
	body []string,
) Predefined {
	return PreFunction{
		name:         name,
		body:         body,
		dependencies: dependencies,
	}
}

// (f PreFunction) add adds the PreFunction to the assembly code being
// generated by the CodeGenerator. It adds any required dependencies and will
// only add the function to the assembly if it is not already present.
func (f PreFunction) add(codegen *CodeGenerator, lib *Library) {
	if _, already := codegen.asm.global[f.name]; !already {
		// Add function to the code
		codegen.asm.global[f.name] = f.body

		// Add any dependencies
		for _, dep := range f.dependencies {
			lib.add(codegen, dep)
		}
	}
}

/**************** PREDEFINED DATA ****************/

// PreData defines a predefined data. The data has a name and the actual text.
// It is normally a string to be printed in the event of a runtime error or
// for formatting other print functions.
type PreData struct {
	name string
	text string
}

// NewPreData returns a PreData struct containing the name and text of the data.
func NewPreData(name string, text string) Predefined {
	return PreData{
		name: name,
		text: text,
	}
}

// (d PreData) add adds the PreData to the assembly code being generated by the
// CodeGenerator. If called twice or more there will still be only one data in
// the final generated assembly code.
func (d PreData) add(codegen *CodeGenerator, lib *Library) {
	codegen.addDataWithLabel(d.name, d.text)
}

/**************** LIBRARY FUNCTIONS ****************/

// LibraryFunction is an enum used internally to reference specific predefined
// functions or data.
type LibraryFunction int

// (l LibraryFunction) String is the method used to convert the LibraryFunction
// enum into a string that can be used as a label in the generated assembly
// code.
func (l LibraryFunction) String() string {
	switch l {
	case printLn:
		return "p_print_ln"
	case printBool:
		return "p_print_bool"
	case printInt:
		return "p_print_int"
	case printString:
		return "p_print_string"
	case printReference:
		return "p_print_reference"
	case readInt:
		return "p_read_int"
	case readChar:
		return "p_read_char"
	case checkDivide:
		return "p_check_divide"
	case checkOverflow:
		return "p_check_overflow"
	case checkArrayIndex:
		return "p_check_array_index"
	case checkNullPointer:
		return "p_check_null_pointer"
	case throwRuntimeError:
		return "p_throw_runtime_error"
	case free:
		return "p_free"
	case msgTrue:
		return "msg_true"
	case msgFalse:
		return "msg_false"
	case msgInt:
		return "msg_int"
	case msgChar:
		return "msg_char"
	case msgNewline:
		return "msg_newline"
	case msgString:
		return "msg_string"
	case msgReference:
		return "msg_reference"
	case msgDivideByZero:
		return "msg_divide_by_zero"
	case msgOverflow:
		return "msg_overflow"
	case msgArrayNegativeIndex:
		return "msg_array_negative_index"
	case msgArrayOutBoundsIndex:
		return "msg_array_out_bounds_index"
	case msgNullPointerReference:
		return "msg_null_pointer_reference"
	default:
		return "unknown"
	}
}

const (
	printLn LibraryFunction = iota
	printBool
	printInt
	printString
	printReference

	readInt
	readChar

	checkDivide
	checkOverflow
	checkArrayIndex
	checkNullPointer
	throwRuntimeError

	free

	msgTrue
	msgFalse
	msgInt
	msgChar
	msgNewline
	msgString
	msgReference
	msgDivideByZero
	msgOverflow
	msgArrayNegativeIndex
	msgArrayOutBoundsIndex
	msgNullPointerReference
)

// Library is struct containing a map from LibraryFunction to Predefined. This
// allows a LibraryFunction to be retrieved from the name and then the add
// method can be called on Predefined to add the required function.
type Library struct {
	store map[LibraryFunction]Predefined
}

// add adds the given LibraryFunction function to the assembly code being
// generated by the CodeGenerator. It adds the predefined function or data in
// the correct location and also adds any required dependencies for that
// LibraryFunction.
func (lib *Library) add(codegen *CodeGenerator, function LibraryFunction) {
	lib.store[function].add(codegen, lib)
}

// NewPreData adds predefined data to the Library, it requires the name
// of the data being added and the text of the data. The name of the data in
// the generated assembly is derived from the LibraryFunction String method.
func (lib *Library) NewPreData(function LibraryFunction, text string) {
	lib.store[function] = NewPreData(function.String(), text)
}

// NewPreFunction add a predefined function to the Library, it requires the name
// of the function being added, a list of dependencies required for the given
// function to work and the body of the function. The name of the function in
// the generated assembly is derived from the LibraryFunction String method.
func (lib *Library) NewPreFunction(
	function LibraryFunction,
	dependencies []LibraryFunction,
	body []string,
) {
	lib.store[function] = NewPreFunction(function.String(), dependencies, body)
}

// GetLibrary returns a pointer to a complete Library of predefined functions
// and strings to be included in the assembly code if necessary.
func GetLibrary() *Library {
	library := &Library{
		store: make(map[LibraryFunction]Predefined),
	}

	// Predefined Functions
	// Printing Functions
	library.NewPreFunction(printLn,
		[]LibraryFunction{
			msgNewline,
		},
		[]string{
			"PUSH {lr}\n",
			"LDR r0, =" + msgNewline.String() + "\n",
			"ADD r0, r0, #4\n",
			"BL puts\n",
			"MOV r0, #0\n",
			"BL fflush\n",
			"POP {pc}\n",
		})
	library.NewPreFunction(printBool,
		[]LibraryFunction{
			msgTrue,
			msgFalse,
		},
		[]string{
			"PUSH {lr}\n",
			"CMP r0, #0\n",
			"LDRNE r0, =" + msgTrue.String() + "\n",
			"LDREQ r0, =" + msgFalse.String() + "\n",
			"ADD r0, r0, #4\n",
			"BL printf\n",
			"MOV r0, #0\n",
			"BL fflush\n",
			"POP {pc}\n",
		})
	library.NewPreFunction(printInt,
		[]LibraryFunction{
			msgInt,
		},
		[]string{
			"PUSH {lr}\n",
			"MOV r1, r0\n",
			"LDR r0, =" + msgInt.String() + "\n",
			"ADD r0, r0, #4\n",
			"BL printf\n",
			"MOV r0, #0\n",
			"BL fflush\n",
			"POP {pc}\n",
		})
	library.NewPreFunction(printString,
		[]LibraryFunction{
			msgString,
		},
		[]string{
			"PUSH {lr}\n",
			"LDR r1, [r0]\n",
			"ADD r2, r0, #4\n",
			"LDR r0, =" + msgString.String() + "\n",
			"ADD r0, r0, #4\n",
			"BL printf\n",
			"MOV r0, #0\n",
			"BL fflush\n",
			"POP {pc}\n",
		})
	library.NewPreFunction(printReference,
		[]LibraryFunction{
			msgReference,
		},
		[]string{
			"PUSH {lr}\n",
			"MOV r1, r0\n",
			"LDR r0, =" + msgReference.String() + "\n",
			"ADD r0, r0, #4\n",
			"BL printf\n",
			"MOV r0, #0\n",
			"BL fflush\n",
			"POP {pc}\n",
		})

	// Reading functions
	library.NewPreFunction(readInt,
		[]LibraryFunction{
			msgInt,
		},
		[]string{
			"PUSH {lr}\n",
			"MOV r1, r0\n",
			"LDR r0, =" + msgInt.String() + "\n",
			"ADD r0, r0, #4\n",
			"BL scanf\n",
			"POP {pc}\n",
		})
	library.NewPreFunction(readChar,
		[]LibraryFunction{
			msgChar,
		},
		[]string{
			"PUSH {lr}\n",
			"MOV r1, r0\n",
			"LDR r0, =" + msgChar.String() + "\n",
			"ADD r0, r0, #4\n",
			"BL scanf\n",
			"POP {pc}\n",
		})

	// Runtime Error Functions
	library.NewPreFunction(checkDivide,
		[]LibraryFunction{
			throwRuntimeError,
			msgDivideByZero,
		},
		[]string{
			"PUSH {lr}\n",
			"CMP r1, #0\n",
			"LDREQ r0, =" + msgDivideByZero.String() + "\n",
			"BLEQ " + throwRuntimeError.String() + "\n",
			"POP {pc}\n",
		})
	library.NewPreFunction(checkOverflow,
		[]LibraryFunction{
			throwRuntimeError,
			msgOverflow,
		},
		[]string{
			"LDR r0, =" + msgOverflow.String() + "\n",
			"BL " + throwRuntimeError.String() + "\n",
		})
	library.NewPreFunction(checkArrayIndex,
		[]LibraryFunction{
			throwRuntimeError,
			msgArrayOutBoundsIndex,
			msgArrayNegativeIndex,
		},
		[]string{
			"PUSH {lr}\n",
			"CMP r0, #0\n",
			"LDRLT r0, =" + msgArrayNegativeIndex.String() + "\n",
			"BLLT " + throwRuntimeError.String() + "\n",
			"LDR r1, [r1]\n",
			"CMP r0, r1\n",
			"LDRCS r0, =" + msgArrayOutBoundsIndex.String() + "\n",
			"BLCS " + throwRuntimeError.String() + "\n",
			"POP {pc}\n",
		})
	library.NewPreFunction(checkNullPointer,
		[]LibraryFunction{
			msgNullPointerReference,
			throwRuntimeError,
		},
		[]string{
			"PUSH {lr}\n",
			"CMP r0, #0\n",
			"LDREQ r0, =" + msgNullPointerReference.String() + "\n",
			"BLEQ " + throwRuntimeError.String() + "\n",
			"POP {pc}\n",
		})
	library.NewPreFunction(throwRuntimeError,
		[]LibraryFunction{
			printString,
		},
		[]string{
			"BL " + printString.String() + "\n",
			"MOV r0, #-1\n",
			"BL exit\n",
		})

	// Free functions
	library.NewPreFunction(free,
		[]LibraryFunction{
			msgNullPointerReference,
			throwRuntimeError,
		},
		[]string{
			"PUSH {lr}\n",
			"CMP r0, #0\n",
			"LDREQ r0, =" + msgNullPointerReference.String() + "\n",
			"BEQ " + throwRuntimeError.String() + "\n",
			"PUSH {r0}\n",
			"LDR r0, [r0]\n",
			"BL free\n",
			"LDR r0, [sp]\n",
			"LDR r0, [r0, #4]\n",
			"BL free\n",
			"POP {r0}\n",
			"BL free\n",
			"POP {pc}\n",
		})

	// Predefined strings
	library.NewPreData(msgTrue, "true\\0")
	library.NewPreData(msgFalse, "false\\0")
	library.NewPreData(msgNewline, "\\0")
	library.NewPreData(msgInt, "%d\\0")
	library.NewPreData(msgChar, " %c\\0")
	library.NewPreData(msgString, "%.*s\\0")
	library.NewPreData(msgReference, "%p\\0")
	library.NewPreData(
		msgDivideByZero,
		"DivideByZeroError: divide or modulo by zero\\n\\0",
	)
	library.NewPreData(
		msgOverflow,
		"OverflowError: the result is too small/large to store in a 4-byte "+
			"signed-integer.\\n",
	)
	library.NewPreData(
		msgArrayNegativeIndex,
		"ArrayIndexOutOfBoundsError: negative index\\n\\0",
	)
	library.NewPreData(
		msgArrayOutBoundsIndex,
		"ArrayIndexOutOfBoundsError: index too large\\n\\0",
	)
	library.NewPreData(
		msgNullPointerReference,
		"NullReferenceError: dereference a null reference\\n\\0",
	)

	return library
}
