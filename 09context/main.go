package main

import (
	contextbasics "go-context/context-basics"
)

/**
Sure, I'd be happy to help you understand how to use contexts in Go (Golang)! The `context` package in Go is essential for managing deadlines, cancellations, and request-scoped values. Let's start with some basics and build up to more complex use cases.

### Basics of Context in Go

The `context` package provides four main functions to create and interact with contexts:
1. `context.Background()`
2. `context.TODO()`
3. `context.WithCancel(parent Context)`
4. `context.WithTimeout(parent Context, timeout)`
5. `context.WithDeadline(parent Context, deadline)`
6. `context.WithValue(parent Context, key, val)`

#### 1. `context.Background()`
`context.Background()` returns an empty context. It is typically used at the top level of a call chain.

#### 2. `context.TODO()`
`context.TODO()` is used when you're not sure which context to use or it hasn't been decided yet.

#### 3. `context.WithCancel(parent Context)`
`context.WithCancel(parent)` returns a copy of the parent context and a `cancel` function. Calling the `cancel` function cancels the context, which can be used to stop goroutines.

#### 4. `context.WithTimeout(parent Context, timeout)`
`context.WithTimeout(parent, timeout)` returns a copy of the parent context with a deadline set to the current time plus the timeout duration. The returned context is automatically canceled when the timeout expires.

#### 5. `context.WithDeadline(parent Context, deadline)`
`context.WithDeadline(parent, deadline)` returns a copy of the parent context with the specified deadline. The returned context is automatically canceled when the deadline is reached.

#### 6. `context.WithValue(parent Context, key, val)`
`context.WithValue(parent, key, val)` returns a copy of the parent context with an associated key-value pair. Use this sparingly as it is not meant for passing optional parameters to functions.
`

### Example: Using Context with Goroutines

In this example:
- We create a cancellable context with `context.WithCancel`.
- We start a goroutine that checks for the cancellation signal using `ctx.Done()`.
- After 2 seconds, we call the `cancel` function to cancel the context.
- The goroutine stops executing upon receiving the cancellation signal.


**/

func main() {
	// contextbasics.CancelMain()
	// contextbasics.DeadlineMain()
	// contextbasics.TimeoutMain()
	// contextbasics.ValueMain()
	contextbasics.ContextSample()
}
