package routeoption

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"professorc/src/environment"
	"professorc/src/evaluate"
	"professorc/src/lex"
	"professorc/src/parse"
	"strings"
)

func isFile(name string) bool {
	if !strings.Contains(name, ".") {
		fmt.Printf("Wrong File. got=%s\n", name)
		return false
	}

	splitedFileName := strings.Split(name, ".")
	if splitedFileName[1] != "gyosoo" {
		fmt.Printf("Not Professorlang File. got=%s\n", name)
		return false
	}

	return true
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Interpreter Error. don't ran Interpreting\n")
	io.WriteString(out, " parser Error:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func Interpreter(name string) {
	if !isFile(name) {
		fmt.Printf("Enter the professorlang file.\n")
		return
	}

	dat, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Professorlang File Error. We don't Interpreting this\n")
		return
	}
	input := string(dat)

	out := os.Stdout

	l := lex.New(input)
	p := parse.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	env := environment.New()
	pc := evaluate.NewPC(len(program.Statements))

	evaluate.Eval(program, env, pc)
}
