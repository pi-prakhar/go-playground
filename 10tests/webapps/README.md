# Go-Test-WebApps

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/yourproject)](https://goreportcard.com/report/github.com/yourusername/yourproject)
[![Coverage Status](https://coveralls.io/repos/github/yourusername/yourproject/badge.svg?branch=main)](https://coveralls.io/github/yourusername/yourproject?branch=main)

## Description

This repository contains a simple web application built using Go and Gohtml. The purpose of this project is to demonstrate effective test writing practices in Go. The application showcases various features, including:

- **Routing:** Efficient handling of web routes.
- **Templates:** Use of Gohtml for server-side rendering of HTML templates.
- **Handlers:** Clean and modular handler functions.
- **Testing:** Comprehensive test coverage to ensure code quality and reliability.

By following this example, developers can learn how to structure their Go applications, write maintainable code, and implement thorough testing strategies.
   
## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install)
- Windows CMD terminal

### Running Tests

To ensure the code quality and functionality, you can run various Go test commands. The following commands should be executed in the project's base directory using the Windows CMD terminal.

#### Run All Tests Verbosely

```sh
go test -v ./...
```
This command runs all the tests in the project with verbose output, which includes detailed information about each test that is run.

#### Generate Test Coverage Report

```sh
go test -coverprofile=coverage.out ./...
```
This command runs all the tests and generates a coverage profile named `coverage.out`. This profile can be used to check which parts of the code are covered by tests.

#### Generate Coverage Report for Current Package

```sh
go test -coverprofile=coverage.out
```
This command runs tests and generates a coverage profile for the current package.

#### View Coverage Report in Function-Level Detail

```sh
go tool cover -func=coverage.out
```
This command processes the coverage profile generated by the previous commands and displays a detailed function-level coverage report.

#### Generate HTML Coverage Report

```sh
go tool cover -html=coverage.out -o coverage.html
start coverage.html
```
These commands generate an HTML file from the coverage profile and open it in the default web browser to visually inspect the coverage.