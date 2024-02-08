package cms

import (
	"io"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/gomarkdown/markdown/ast"
)

var (
	formatter *html.Formatter = html.New(html.WithClasses(true), html.TabWidth(4), html.WrapLongLines(true))
	style     *chroma.Style   = styles.Fallback
)

// func makeStyle() *chroma.Style {
// 	b := chroma.NewStyleBuilder("boons")
// 	b.Add(chroma.Keyword, "")
// 	b.Add(chroma.)

// 	style, err := b.Build()
// 	if err != nil {

// 	}
// 	return style
// }

// adapted from https://github.com/gomarkdown/markdown/blob/master/examples/code_hightlight.go
func highlight(writer io.Writer, source string, lang string) error {
	lexer := lexers.Get(lang)
	if lexer == nil {
		lexer = lexers.Analyse(source)
	}
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	iterator, err := lexer.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return formatter.Format(writer, style, iterator)
}

func renderCode(writer io.Writer, codeBlock *ast.CodeBlock, entering bool) {
	lang := string(codeBlock.Info)
	highlight(writer, string(codeBlock.Literal), lang)
}

func codeRenderHook(writer io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if code, ok := node.(*ast.CodeBlock); ok {
		renderCode(writer, code, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}
