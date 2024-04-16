//go:build !go1.20
// +build !go1.20

package syntax

import "cmd/compile/internal/syntax"

func (ctx *BlockContext) traverseCallStmtCallExpr(node *syntax.CallExpr, globaleNames map[string]*DeclInfo, imports map[string]string) *syntax.CallExpr {
	return ctx.traverseCallExpr(node, globaleNames, imports)
}

func copyCallExpr(expr *syntax.CallExpr) *syntax.CallExpr {
	if expr == nil {
		return nil
	}
	x := *expr
	x.Fun = copyExpr(expr.Fun)
	x.ArgList = copyExprs(expr.ArgList)
	return &x
}
