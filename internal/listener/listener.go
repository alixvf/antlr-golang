package listener

import (
	"fmt"
	parser "go_javacfg/gen"
	"log"
)

type CFGExtractorListener struct {
	*parser.BaseJava20ParserListener
}

func NewCFGExtractorListener() *CFGExtractorListener {
	return &CFGExtractorListener{}
}

func (c *CFGExtractorListener) EnterNormalClassDeclaration(ctx *parser.NormalClassDeclarationContext) {
	className := ctx.TypeIdentifier().GetText()
	log.Println("Entering class -->", className)
}

func (c *CFGExtractorListener) ExitNormalClassDeclaration(ctx *parser.NormalClassDeclarationContext) {
	log.Println("Exiting class -->", ctx.TypeIdentifier().GetText())
}

func (c *CFGExtractorListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	methodName := ctx.MethodHeader().MethodDeclarator().Identifier().GetText()
	log.Println(fmt.Sprintf("Entering method --> %s$$%s", methodName))
}

func (c *CFGExtractorListener) ExitMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	methodName := ctx.MethodHeader().MethodDeclarator().Identifier().GetText()
	log.Println(fmt.Sprintf("Exiting method --> %s$$%s", methodName))
}
