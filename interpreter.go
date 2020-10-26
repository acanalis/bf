package braingo

import (
	"io"
)

type Interpreter struct {
	program string
	input   io.Reader
	output  io.Writer
	ip      int
	dp      int
	tape    [3000000]byte
	jump    map[int]int
}

func NewInterpreter(program string) *Interpreter {
	return &Interpreter{
		program: program,
		dp:      1500000,
		jump:    matchbrackets(program),
	}
}

func matchbrackets(program string) (match map[int]int) {
	match = make(map[int]int)
	var brackets []int
	for i, s := range program {
		if s == '[' {
			brackets = append(brackets, i)
		}
		if s == ']' {
			// if len(brackets) == 0 {return nil}
			match[i] = brackets[len(brackets)-1]
			match[brackets[len(brackets)-1]] = i
			brackets = brackets[0 : len(brackets)-1]
		}
	}
	// if len(brackets) != 0 {return nil}
	return match
}

func (i *Interpreter) Run(input io.Reader, output io.Writer) {
	for {
		if len(i.program) <= i.ip {
			return
		}
		switch i.program[i.ip] {
		case '[':
			if i.tape[i.dp] == 0 {
				i.ip = i.jump[i.ip]
			}
		case ']':
			if i.tape[i.dp] != 0 {
				i.ip = i.jump[i.ip]
			}
		case '>':
			i.dp++
		case '<':
			i.dp--
		case '+':
			i.tape[i.dp]++
		case '-':
			i.tape[i.dp]--
		case '.':
			output.Write(i.tape[i.dp : i.dp+1])
		case ',':
			_, err := input.Read(i.tape[i.dp : i.dp+1])
			if err != nil {
				return
			}
		}
		i.ip++

	}
}
