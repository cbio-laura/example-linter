package linters

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var TodoAnalyzer = &analysis.Analyzer{
	Name: "duplicate_error",
	Doc:  "Find duplicate error codes within repo.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	errorCodesSeen := make(map[string]bool)

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {

			// Make sure this is a function call
			functionCall, isFunctionCall := n.(*ast.CallExpr)
			if !isFunctionCall {
				// fmt.Printf("%v is not function call \n\n", n)
				return true
			}

			var baseLooksLike string
			var functionLooksLike string

			selectorExpression, ok := functionCall.Fun.(*ast.SelectorExpr)
			if !ok {

				functionIdent, ok := functionCall.Fun.(*ast.Ident)
				if !ok {
					return true
				}

				functionLooksLike = functionIdent.String()

			} else {
				baseIdent, ok := selectorExpression.X.(*ast.Ident)
				if !ok {
					return true
				}

				baseLooksLike = baseIdent.String()
				functionLooksLike = selectorExpression.Sel.String()
			}

			if (baseLooksLike == "errors" && strings.HasPrefix(functionLooksLike, "New")) ||
				functionLooksLike == "errorFx" {

				// Error expressions have either 2 or 3 arguments
				if len(functionCall.Args) < 2 {
					return true
				}

				rawErrorNumber := functionCall.Args[0]

				// The error number should be a string
				basicLiteralErrorNumber, ok := rawErrorNumber.(*ast.BasicLit)
				if !ok {
					// This would indicate a function call looked like
					// `errors.New` but first param isn't an identity.
					return true
				}

				// Check if this error code is a dup
				_, errorNumberAlreadyFound := errorCodesSeen[basicLiteralErrorNumber.Value]
				if errorNumberAlreadyFound {
					pass.Report(analysis.Diagnostic{
						Pos:            rawErrorNumber.Pos(),
						End:            rawErrorNumber.End(),
						Category:       "duplicate_error",
						Message:        fmt.Sprintf("Error number %s has already been seen", basicLiteralErrorNumber.Value),
						SuggestedFixes: nil,
					})
				}

				// If error code hasn't been seen yet, add it to the map and
				// continue.
				errorCodesSeen[basicLiteralErrorNumber.Value] = true
			}

			// If string doesn't have prefix, we don't care; some random
			// function call
			return true

		})
	}
	return nil, nil
}
