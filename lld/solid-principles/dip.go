package main

import "fmt"

/*
Dependency Inversion Principle (DIP)
High-level modules should not depend on low-level modules. Both should depend on abstractions.

In Go:

1) Use dependency injection and rely on interfaces to decouple high-level and low-level modules.
2) Pass dependencies as parameters to functions or constructors.
*/

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println("Log:", message)
}

type Application struct {
	logger Logger
}

func NewApplication(logger Logger) *Application {
	return &Application{logger: logger}
}

func (app *Application) Run() {
	app.logger.Log("Application is running")
}
