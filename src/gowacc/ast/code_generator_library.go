package ast

import (
	"strconv"
)

type Predefined interface {
	add(*CodeGenerator, *Library)
}

type PreFunction struct {
  name string
  body []string
  dependancies []LibraryFunction
}

func NewPreFunction(name string, dependancies []LibraryFunction, body []string) Predefined {
	return PreFunction{
		name: name,
    body: body,
    dependancies: dependancies,
	}
}

func (f PreFunction) add(v *CodeGenerator, lib *Library) {
  if _, already := v.asm.global[f.name]; !already {
    // Add function to the code
    v.asm.global[f.name] = f.body

    // Add any dependancies
    for _, dep := range f.dependancies {
      lib.add(v, dep)
    }
	}
}

type PreData struct {
  name string
  text string
  length int
  dependancies []LibraryFunction
}

func NewPreData(name string, text string, length int) Predefined {
	return PreData{
		name: name,
    text: text,
    length: length,
	}
}

func (d PreData) add(v *CodeGenerator, lib *Library) {
  v.addDataWithLabel(d.name, d.text, d.length)

  // Add any dependancies
  for _, dep := range d.dependancies {
    lib.add(v, dep)
  }
}

type LibraryFunction int

func (l LibraryFunction) String() string {
	return "lib_" + strconv.Itoa(int(l))
}
const (
  PRINT_LN LibraryFunction = iota
  PRINT_BOOL
  PRINT_INT
	PRINT_STRING

	CHECK_DIVIDE
	THROW_RUNTIME_ERROR

  MSG_TRUE
  MSG_FALSE
  MSG_INT
  MSG_NEWLINE
	MSG_STRING
	MSG_DIVIDE_BY_ZERO
)

type Library struct {
  lib map[LibraryFunction]Predefined
}

func (l *Library) add(v *CodeGenerator, function LibraryFunction) {
  l.lib[function].add(v, l)
}

func (l *Library) NewPreData(f LibraryFunction, text string, length int) {
  l.lib[f] = NewPreData(f.String(), text, length)
}

func (l *Library) NewPreFunction(f LibraryFunction, dependancies []LibraryFunction, body []string) {
  l.lib[f] = NewPreFunction(f.String(), dependancies, body)
}

func GetLibrary() *Library {
  library := &Library {
		lib: make(map[LibraryFunction]Predefined),
	}
  // Predefined Functions
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

	library.NewPreFunction(THROW_RUNTIME_ERROR,
		[]LibraryFunction{
				PRINT_STRING,
			},
		[]string{
			"BL " + PRINT_STRING.String(),
			"MOV r0, #-1",
			"BL exit",
		})

  // Predefined strings
  library.NewPreData(MSG_TRUE, "true\\0", 5)
  library.NewPreData(MSG_FALSE, "false\\0", 6)
  library.NewPreData(MSG_NEWLINE, "\\0", 1)
  library.NewPreData(MSG_INT, "%d\\0", 3)
	library.NewPreData(MSG_STRING, "%.*s\\0", 5)
	library.NewPreData(MSG_DIVIDE_BY_ZERO, "DivideByZeroError: divide or modulo by zero\\n\\0", 45)

  return library
}
