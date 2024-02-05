package main

import (
	"github.com/antlr4-go/antlr/v4"
	parser "go_javacfg/gen"
	"go_javacfg/internal/listener"
)

func main() {
	input, _ := antlr.NewFileStream("./javasamples/for.java")
	lexer := parser.NewJava20Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJava20Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.CompilationUnit()
	walker := antlr.NewParseTreeWalker()
	walker.Walk(listener.NewCFGExtractorListener(), tree)
}
