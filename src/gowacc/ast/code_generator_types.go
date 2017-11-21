package ast

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Probably best to move all the structs and stuff here

type Predefined interface {
	add(*CodeGenerator)
}

type PreFunction struct {
  name string
  body []string
  dependancies []Predefined
}

func NewPreFunction(name string, dependancies []Predefined, body []string) Predefined {
	return PreFunction{
		name: name,
    body: body,
    dependancies: dependancies,
	}
}

func (f PreFunction) add(v *CodeGenerator) {
  // Add function to the code
  v.asm.global[f.name] = f.body

  // Add any dependancies
  for dep := range f.dependancies {
    dep.add(v)
  }
}

type PreData struct {
  name string
  text string
  dependancies []Predefined
}

func NewPreData(name string, text string) Predefined {
	return PreData{
		name: name,
    text: text,
	}
}

func (d PreData) add(v *CodeGenerator) {
  v.addDataWithLabel(d.name, d.text)

  // Add any dependancies
  for dep := range d.dependancies {
    dep.add(v)
  }
}

const (

  // Predefined Functions
  PRINT_LN Predefined = NewPreFunction("p_print_ln",
    []Predefined{
      MSG_NEWLINE
    }
    []string{
    	"PUSH {lr}",
    	"LDR r0, =msg_p_newline",
    	"ADD r0, r0, #4",
    	"BL puts",
    	"MOV r0, #0",
    	"BL fflush",
    	"POP {pc}"
    })
  PRINT_BOOL Predefined = NewPreFunction("p_print_bool",
    []Predefined{
      MSG_TRUE,
      MSG_FALSE
    }
    []string{
      "PUSH {lr}",
      "CMP r0, #0",
      "LDRNE r0, =msg_p_true",
      "LDREQ r0, =msg_p_false",
      "ADD r0, r0, #4",
      "BL printf",
      "MOV r0, #0",
      "BL fflush",
      "POP {pc})"
    })
  PRINT_INT Predefined = NewPreFunction("p_print_int",
    []Predefined{
      MSG_INT
    }
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
  MSG_TRUE Predefined = NewPreData("msg_p_true", "true\\0")
  MSG_FALSE Predefined = NewPreData("msg_p_false", "false\\0")
  MSG_NEWLINE Predefined = NewPreData("msg_p_newline", "\\0")
  MSG_INT Predefined = NewPreData("msg_p_int", "%d\\0")

)
