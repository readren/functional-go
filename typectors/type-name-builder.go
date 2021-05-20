package typectors

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func BuildTypeName(tpe string) string {
	node, err := parser.ParseExpr(tpe + "{}")
	if err != nil {
		panic(fmt.Errorf(`unable to parse the type "%s" : %w`, tpe, err))
	}
	load := load{tpe, strings.Builder{}}
	v := skipCompositeLitVisitor{&load}
	ast.Walk(v, node)
	return v.loadPtr.builder.String()
}

type load struct {
	tpe     string
	builder strings.Builder
}
type skipCompositeLitVisitor struct{ loadPtr *load }
type buildNameVisitor struct{ loadPtr *load }
type arrayLenVisitor struct{ loadPtr *load }
type identifierVisitor struct{ loadPtr *load }

func (v skipCompositeLitVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.CompositeLit:
		return buildNameVisitor{v.loadPtr}
	default:
		panic(fmt.Errorf(`invalid or unsupported type "%s"`, v.loadPtr.tpe))
	}
}

func (v buildNameVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case nil:
		return nil
	case *ast.MapType:
		v.loadPtr.builder.WriteString("mapFrom_")
		ast.Walk(v, n.Key)
		v.loadPtr.builder.WriteString("_to_")
		ast.Walk(v, n.Value)
		return nil

	case *ast.ArrayType:
		if n.Len == nil {
			v.loadPtr.builder.WriteString("slice_")
		} else {
			v.loadPtr.builder.WriteString("array")
			ast.Walk(arrayLenVisitor{v.loadPtr}, n.Len)
			v.loadPtr.builder.WriteRune('_')
		}
		ast.Walk(v, n.Elt)
		return nil

	case *ast.Ident:
		v.loadPtr.builder.WriteString(n.Name)
		return nil

	case *ast.SelectorExpr:
		ast.Walk(identifierVisitor{v.loadPtr}, n.X)
		ast.Walk(identifierVisitor{v.loadPtr}, n.Sel)
		return nil

	default:
		panic(fmt.Errorf(`invalid or unsupported node "%T:%v" when parsing the type "%s"`, n, n, v.loadPtr.tpe))
	}
}

func (v arrayLenVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.BasicLit:
		if n.Kind == token.INT {
			v.loadPtr.builder.WriteString(string(n.Value))
			return nil
		}
	}
	panic(fmt.Errorf(`the array size in the type "%s" should be a literal int expression`, v.loadPtr.tpe))
}

func (v identifierVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.Ident:
		v.loadPtr.builder.WriteString(string(n.Name))
		return nil
	}
	panic(fmt.Errorf(`the array size in the type "%s" should be a literal int expression`, v.loadPtr.tpe))
}
