package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/buildtag"
	"golang.org/x/tools/go/analysis/passes/errorsas"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/structtag"

	"github.com/cucumberjaye/url-shortener/pkg/myanalyzers"
)

func main() {
	analyzers := []*analysis.Analyzer{
		// passes analyzers
		printf.Analyzer,
		shift.Analyzer,
		structtag.Analyzer,
		shadow.Analyzer,
		errorsas.Analyzer,
		buildtag.Analyzer,
		assign.Analyzer,

		//custom os.Exit analyzer
		myanalyzers.ExitCheckAnalyzer,
	}

	/*count := 0
	for _, v := range staticcheck.Analyzers {
		// staticchec.io SA analyzers
		if strings.HasPrefix(v.Analyzer.Name, "SA") {
			analyzers = append(analyzers, v.Analyzer)
			continue
		}
		// staticchec.io 3 other analyzers
		if count < 3 {
			analyzers = append(analyzers, v.Analyzer)
			count++
		}
	}*/

	multichecker.Main(analyzers...)
}
