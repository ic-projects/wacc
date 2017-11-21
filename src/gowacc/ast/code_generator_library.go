package ast

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
const (
  PRINT_LN LibraryFunction = iota
  PRINT_BOOL
  PRINT_INT
  MSG_TRUE
  MSG_FALSE
  MSG_INT
  MSG_NEWLINE
)

type Library struct {
  lib map[LibraryFunction]Predefined
}

func (l *Library) add(v *CodeGenerator, function LibraryFunction) {
  l.lib[function].add(v, l)
}

func GetLibrary() *Library {
  library := make(map[LibraryFunction]Predefined)

  // Predefined Functions
  library[PRINT_LN] = NewPreFunction("p_print_ln",
    []LibraryFunction{
      MSG_NEWLINE,
    },
    []string{
    	"PUSH {lr}",
    	"LDR r0, =msg_p_newline",
    	"ADD r0, r0, #4",
    	"BL puts",
    	"MOV r0, #0",
    	"BL fflush",
    	"POP {pc}",
    })
  library[PRINT_BOOL] = NewPreFunction("p_print_bool",
    []LibraryFunction{
      MSG_TRUE,
      MSG_FALSE,
    },
    []string{
      "PUSH {lr}",
      "CMP r0, #0",
      "LDRNE r0, =msg_p_true",
      "LDREQ r0, =msg_p_false",
      "ADD r0, r0, #4",
      "BL printf",
      "MOV r0, #0",
      "BL fflush",
      "POP {pc}",
    })
  library[PRINT_INT] = NewPreFunction("p_print_int",
    []LibraryFunction{
      MSG_INT,
    },
    []string{
      "PUSH {lr}",
      "MOV r1, r0",
      "LDR r0, =msg_p_int",
      "ADD r0, r0, #4",
      "BL printf",
      "MOV r0, #0",
      "BL fflush",
      "POP {pc}",
    })
  //PRINT_STRING

  // Predefined strings
  library[MSG_TRUE] = NewPreData("msg_p_true", "true\\0", 5)
  library[MSG_FALSE] = NewPreData("msg_p_false", "false\\0", 6)
  library[MSG_NEWLINE] = NewPreData("msg_p_newline", "\\0", 1)
  library[MSG_INT] = NewPreData("msg_p_int", "%d\\0", 3)


  return &Library {
    lib: library,
  }
}
