package main

import "fmt"

/*
Single Responsibility Principle (SRP)
A class or module should have only one reason to change, meaning it should have one responsibility.

In Go:

1) Design your packages and structs to focus on one responsibility.
2) Separate concerns by creating smaller, modular functions and files.
3) Avoid putting multiple, unrelated functionalities in a single struct or function.

*/

type ReportGenerator struct {
	// Fields for report generation
}

func (rg *ReportGenerator) GenerateReport() string {
	// Generate report logic
	return "report"
}

type ReportPrinter struct{}

func (rp *ReportPrinter) Print(report string) {
	fmt.Println(report)
}
