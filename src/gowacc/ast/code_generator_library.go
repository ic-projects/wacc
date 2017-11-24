package ast

// Predefined is an interface that defines a function or data that can be added
// to the CodeGenerators assembly.
type Predefined interface {
	add(*CodeGenerator, *Library)
}

// PreFunction defines a predefined function. The function has a name, the body of
// function containing the code and a dependency list of other LibraryFunctions
// required for the function to work.
type PreFunction struct {
	name         string
	body         []string
	dependencies []LibraryFunction
}

// NewPreFunction returns a PreFunction struct containing the name, body and dependencies
// of the function.
func NewPreFunction(name string, dependencies []LibraryFunction, body []string) Predefined {
	return PreFunction{
		name:         name,
		body:         body,
		dependencies: dependencies,
	}
}

// (f PreFunction) add adds the PreFunction to the assembly code being generated by the
// CodeGenerator. It adds any required dependencies and will only add the function
// to the assembly if it is not already present.
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

// LibraryFunction is an enum used internally to reference specific predefined
// functions or data.
type LibraryFunction int

// (l LibraryFunction) String is the method used to convert the LibraryFunction enum
// into a string that can be used as a label in the generated assembly code.
func (l LibraryFunction) String() string {
	switch l {
	case PRINT_LN:
		return "p_print_ln"
	case PRINT_BOOL:
		return "p_print_bool"
	case PRINT_INT:
		return "p_print_int"
	case PRINT_STRING:
		return "p_print_string"
	case PRINT_REFERENCE:
		return "p_print_reference"
	case READ_INT:
		return "p_read_int"
	case READ_CHAR:
		return "p_read_char"
	case CHECK_DIVIDE:
		return "p_check_divide"
	case CHECK_OVERFLOW:
		return "p_check_overflow"
	case CHECK_ARRAY_INDEX:
		return "p_check_array_index"
	case CHECK_NULL_POINTER:
		return "p_check_null_pointer"
	case THROW_RUNTIME_ERROR:
		return "p_throw_runtime_error"
	case FREE:
		return "p_free"
	case MSG_TRUE:
		return "msg_true"
	case MSG_FALSE:
		return "msg_false"
	case MSG_INT:
		return "msg_int"
	case MSG_CHAR:
		return "msg_char"
	case MSG_NEWLINE:
		return "msg_newline"
	case MSG_STRING:
		return "msg_string"
	case MSG_REFERENCE:
		return "msg_reference"
	case MSG_DIVIDE_BY_ZERO:
		return "msg_divide_by_zero"
	case MSG_OVERFLOW:
		return "msg_overflow"
	case MSG_ARRAY_NEGATIVE_INDEX:
		return "msg_array_negative_index"
	case MSG_ARRAY_OUT_BOUNDS_INDEX:
		return "msg_array_out_bounds_index"
	case MSG_NULL_POINTER_REFERENCE:
		return "msg_null_pointer_reference"
	default:
		return "unknown"
	}
}

const (
	PRINT_LN LibraryFunction = iota
	PRINT_BOOL
	PRINT_INT
	PRINT_STRING
	PRINT_REFERENCE

	READ_INT
	READ_CHAR

	CHECK_DIVIDE
	CHECK_OVERFLOW
	CHECK_ARRAY_INDEX
	CHECK_NULL_POINTER
	THROW_RUNTIME_ERROR

	FREE

	MSG_TRUE
	MSG_FALSE
	MSG_INT
	MSG_CHAR
	MSG_NEWLINE
	MSG_STRING
	MSG_REFERENCE
	MSG_DIVIDE_BY_ZERO
	MSG_OVERFLOW
	MSG_ARRAY_NEGATIVE_INDEX
	MSG_ARRAY_OUT_BOUNDS_INDEX
	MSG_NULL_POINTER_REFERENCE
	MSG_FREE
)

// Library is struct containing a map from LibraryFunction to Predefined. This
// allows a LibraryFunction to be retrived from the name and then the add method
// can be called on Predefined to add the required function.
type Library struct {
	store map[LibraryFunction]Predefined
}

// add adds the given LibraryFunction function to the assembly code being generated
// by the CodeGenerator. It adds the predefined function or data in the correct location
// and also adds any required dependencies for that LibraryFunction.
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
// of the function being added, a list of dependencies required for the given function
// to work and the body of the function. The name of the function in the generated
// assembly is derived from the LibraryFunction String method.
func (lib *Library) NewPreFunction(function LibraryFunction, dependencies []LibraryFunction, body []string) {
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
	library.NewPreFunction(PRINT_LN,
		[]LibraryFunction{
			MSG_NEWLINE,
		},
		[]string{
			"PUSH {lr}",
			"LDR r0, =" + MSG_NEWLINE.String(),
			"ADD r0, r0, #4",
			"BL puts",
			"MOV r0, #0",
			"BL fflush",
			"POP {pc}",
		})
	library.NewPreFunction(PRINT_BOOL,
		[]LibraryFunction{
			MSG_TRUE,
			MSG_FALSE,
		},
		[]string{
			"PUSH {lr}",
			"CMP r0, #0",
			"LDRNE r0, =" + MSG_TRUE.String(),
			"LDREQ r0, =" + MSG_FALSE.String(),
			"ADD r0, r0, #4",
			"BL printf",
			"MOV r0, #0",
			"BL fflush",
			"POP {pc}",
		})
	library.NewPreFunction(PRINT_INT,
		[]LibraryFunction{
			MSG_INT,
		},
		[]string{
			"PUSH {lr}",
			"MOV r1, r0",
			"LDR r0, =" + MSG_INT.String(),
			"ADD r0, r0, #4",
			"BL printf",
			"MOV r0, #0",
			"BL fflush",
			"POP {pc}",
		})
	library.NewPreFunction(PRINT_STRING,
		[]LibraryFunction{
			MSG_STRING,
		},
		[]string{
			"PUSH {lr}",
			"LDR r1, [r0]",
			"ADD r2, r0, #4",
			"LDR r0, =" + MSG_STRING.String(),
			"ADD r0, r0, #4",
			"BL printf",
			"MOV r0, #0",
			"BL fflush",
			"POP {pc}",
		})
	library.NewPreFunction(PRINT_REFERENCE,
		[]LibraryFunction{
			MSG_REFERENCE,
		},
		[]string{
			"PUSH {lr}",
			"MOV r1, r0",
			"LDR r0, =" + MSG_REFERENCE.String(),
			"ADD r0, r0, #4",
			"BL printf",
			"MOV r0, #0",
			"BL fflush",
			"POP {pc}",
		})

	// Reading functions
	library.NewPreFunction(READ_INT,
		[]LibraryFunction{
			MSG_INT,
		},
		[]string{
			"PUSH {lr}",
			"MOV r1, r0",
			"LDR r0, =" + MSG_INT.String(),
			"ADD r0, r0, #4",
			"BL scanf",
			"POP {pc}",
		})
	library.NewPreFunction(READ_CHAR,
		[]LibraryFunction{
			MSG_CHAR,
		},
		[]string{
			"PUSH {lr}",
			"MOV r1, r0",
			"LDR r0, =" + MSG_CHAR.String(),
			"ADD r0, r0, #4",
			"BL scanf",
			"POP {pc}",
		})

	// Runtime Error Functions
	library.NewPreFunction(CHECK_DIVIDE,
		[]LibraryFunction{
			THROW_RUNTIME_ERROR,
			MSG_DIVIDE_BY_ZERO,
		},
		[]string{
			"PUSH {lr}",
			"CMP r1, #0",
			"LDREQ r0, =" + MSG_DIVIDE_BY_ZERO.String(),
			"BLEQ " + THROW_RUNTIME_ERROR.String(),
			"POP {pc}",
		})
	library.NewPreFunction(CHECK_OVERFLOW,
		[]LibraryFunction{
			THROW_RUNTIME_ERROR,
			MSG_OVERFLOW,
		},
		[]string{
			"LDR r0, =" + MSG_OVERFLOW.String(),
			"BL " + THROW_RUNTIME_ERROR.String(),
		})
	library.NewPreFunction(CHECK_ARRAY_INDEX,
		[]LibraryFunction{
			THROW_RUNTIME_ERROR,
			MSG_ARRAY_OUT_BOUNDS_INDEX,
			MSG_ARRAY_NEGATIVE_INDEX,
		},
		[]string{
			"PUSH {lr}",
			"CMP r0, #0",
			"LDRLT r0, =" + MSG_ARRAY_NEGATIVE_INDEX.String(),
			"BLLT " + THROW_RUNTIME_ERROR.String(),
			"LDR r1, [r1]",
			"CMP r0, r1",
			"LDRCS r0, =" + MSG_ARRAY_OUT_BOUNDS_INDEX.String(),
			"BLCS " + THROW_RUNTIME_ERROR.String(),
			"POP {pc}",
		})
	library.NewPreFunction(CHECK_NULL_POINTER,
		[]LibraryFunction{
			MSG_NULL_POINTER_REFERENCE,
			THROW_RUNTIME_ERROR,
		},
		[]string{
			"PUSH {lr}",
			"CMP r0, #0",
			"LDREQ r0, =" + MSG_NULL_POINTER_REFERENCE.String(),
			"BLEQ " + THROW_RUNTIME_ERROR.String(),
			"POP {pc}",
		})
	library.NewPreFunction(THROW_RUNTIME_ERROR,
		[]LibraryFunction{
			PRINT_STRING,
		},
		[]string{
			"BL " + PRINT_STRING.String(),
			"MOV r0, #-1",
			"BL exit",
		})

	// Free functions
	library.NewPreFunction(FREE,
		[]LibraryFunction{
			MSG_NULL_POINTER_REFERENCE,
			THROW_RUNTIME_ERROR,
		},
		[]string{
			"PUSH {lr}",
			"CMP r0, #0",
			"LDREQ r0, =" + MSG_NULL_POINTER_REFERENCE.String(),
			"BEQ " + THROW_RUNTIME_ERROR.String(),
			"PUSH {r0}",
			"LDR r0, [r0]",
			"BL free",
			"LDR r0, [sp]",
			"LDR r0, [r0, #4]",
			"BL free",
			"POP {r0}",
			"BL free",
			"POP {pc}",
		})

	// Predefined strings
	library.NewPreData(MSG_TRUE, "true\\0")
	library.NewPreData(MSG_FALSE, "false\\0")
	library.NewPreData(MSG_NEWLINE, "\\0")
	library.NewPreData(MSG_INT, "%d\\0")
	library.NewPreData(MSG_CHAR, " %c\\0")
	library.NewPreData(MSG_STRING, "%.*s\\0")
	library.NewPreData(MSG_REFERENCE, "%p\\0")
	library.NewPreData(MSG_DIVIDE_BY_ZERO, "DivideByZeroError: divide or modulo by zero\\n\\0")
	library.NewPreData(MSG_OVERFLOW, "OverflowError: the result is too small/large to store in a 4-byte signed-integer.\\n")
	library.NewPreData(MSG_ARRAY_NEGATIVE_INDEX, "ArrayIndexOutOfBoundsError: negative index\\n\\0")
	library.NewPreData(MSG_ARRAY_OUT_BOUNDS_INDEX, "ArrayIndexOutOfBoundsError: index too large\\n\\0")
	library.NewPreData(MSG_NULL_POINTER_REFERENCE, "NullReferenceError: dereference a null reference\\n\\0")

	return library
}
